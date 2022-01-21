#!/bin/sh

export AWS_PROFILE=isvanilla

CLUSTER_NAME=$(aws cloudformation describe-stacks --stack-name EKSStack --query "Stacks[0].Outputs[?OutputKey=='ClusterName'].OutputValue" --output text)
echo "CLUSTER_NAME: $CLUSTER_NAME"

ARGOCD_SECRETS_ROLE_ARN=$(aws cloudformation describe-stacks --stack-name SecretsStack --query "Stacks[0].Outputs[?OutputKey=='ArgoCDSecretsPolicyARN'].OutputValue" --output text)
echo "ARGOCD_SECRETS_ROLE_ARN: $ARGOCD_SECRETS_ROLE_ARN"

TELEMETRY_SECRETS_ROLE_ARN=$(aws cloudformation describe-stacks --stack-name SecretsStack --query "Stacks[0].Outputs[?OutputKey=='TelemetrySecretsPolicyARN'].OutputValue" --output text)
echo "TELEMETRY_SECRETS_ROLE_ARN: $TELEMETRY_SECRETS_ROLE_ARN"
JAEGER_SECRETS_ROLE_ARN=$(aws cloudformation describe-stacks --stack-name SecretsStack --query "Stacks[0].Outputs[?OutputKey=='JaegerSecretsPolicyARN'].OutputValue" --output text)
echo "JAEGER_SECRETS_ROLE_ARN: $JAEGER_SECRETS_ROLE_ARN"
APP_SECRETS_ROLE_ARN=$(aws cloudformation describe-stacks --stack-name SecretsStack --query "Stacks[0].Outputs[?OutputKey=='AppSecretsPolicyARN'].OutputValue" --output text)
echo "APP_SECRETS_ROLE_ARN: $APP_SECRETS_ROLE_ARN"
CROSSPLANE_SECRETS_ROLE_ARN=$(aws cloudformation describe-stacks --stack-name SecretsStack --query "Stacks[0].Outputs[?OutputKey=='CrossplaneSecretsPolicyARN'].OutputValue" --output text)
echo "CROSSPLANE_SECRETS_ROLE_ARN: $CROSSPLANE_SECRETS_ROLE_ARN"
DNS_ROLE_ARN=$(aws cloudformation describe-stacks --stack-name DNSStack --query "Stacks[0].Outputs[?OutputKey=='ClusterDNSPolicyARN'].OutputValue" --output text)
echo "DNS_ROLE_ARN: $DNS_ROLE_ARN"

USER_ARN=$(aws iam get-user --query 'User.Arn' --output text)
echo "USER_ARN: $USER_ARN"

FQDN=$(aws cloudformation describe-stacks --stack-name DNSStack --query "Stacks[0].Outputs[?OutputKey=='FQDN'].OutputValue" --output text)
echo "FQDN: $FQDN"
ZONE_ID=$(aws cloudformation describe-stacks --stack-name DNSStack --query "Stacks[0].Outputs[?OutputKey=='ZoneID'].OutputValue" --output text)
echo "ZONE_ID: $ZONE_ID"

ES_DOMAIN=$(aws cloudformation describe-stacks --stack-name OpenSearchStack --query "Stacks[0].Outputs[?OutputKey=='OpenSearchDomain'].OutputValue" --output text)
echo "ES_DOMAIN: $ES_DOMAIN"
ES_SECRET=$(aws cloudformation describe-stacks --stack-name OpenSearchStack --query "Stacks[0].Outputs[?OutputKey=='MasterUserSecretName'].OutputValue" --output text | sed 's/.*:secret:\([^:]*\):.*/\1/' | sed 's/-[^-]*$//')
echo "ES_SECRET: $ES_SECRET"

RDS_SECRET=$(aws cloudformation describe-stacks --stack-name RDSStack --query "Stacks[0].Outputs[?OutputKey=='RDSSecretName'].OutputValue" --output text)
echo "RDS_SECRET: $RDS_SECRET"

RDS_HOST=$(aws cloudformation describe-stacks --stack-name RDSStack --query "Stacks[0].Outputs[?OutputKey=='RDSHost'].OutputValue" --output text)
echo "RDS_HOST: $RDS_HOST"

KUBECTL_CONFIG=$(AWS_PROFILE=isvanilla aws cloudformation describe-stacks --stack-name EKSStack --query "Stacks[0].Outputs[?starts_with(OutputKey, 'ClusterConfigCommand')].OutputValue" --output text)
echo "KUBECTL_CONFIG: ${KUBECTL_CONFIG}"
/bin/sh -c "${KUBECTL_CONFIG}"

EKSCTL_VERSION=$(eksctl version)
if [[ $? != 0 ]]; then
  echo "ERROR: eksctl not installed"
  exit 1
fi

# Create Namespaces

kubectl apply -f manifests/jk8s-namespace.yaml
kubectl apply -f manifests/external-dns-namespace.yaml
kubectl apply -f manifests/external-secrets-namespace.yaml
kubectl apply -f manifests/cert-manager-namespace.yaml
kubectl apply -f manifests/app-namespace.yaml
kubectl apply -f manifests/jaeger-namespace.yaml
kubectl apply -f manifests/crossplane-system-namespace.yaml

# Grant AWS User master access to cluster

kubectl patch -n kube-system configmap aws-auth \
  --patch "{\"data\":{\"mapUsers\":\"[{\\\"userarn\\\":\\\"${USER_ARN}\\\",\\\"username\\\":\\\"${USER_ARN}\\\",\\\"groups\\\":[\\\"system:masters\\\"]}]\"}}"

# Create Service Accounts

eksctl utils associate-iam-oidc-provider --cluster=$CLUSTER_NAME --approve

eksctl create iamserviceaccount --cluster=$CLUSTER_NAME \
  --name=external-dns \
  --namespace=external-dns \
  --attach-policy-arn=$DNS_ROLE_ARN \
  --override-existing-serviceaccounts \
  --approve

eksctl create iamserviceaccount --cluster=$CLUSTER_NAME \
  --name=cert-manager \
  --namespace=cert-manager \
  --attach-policy-arn=$DNS_ROLE_ARN \
  --override-existing-serviceaccounts \
  --approve  

# Create Service Accounts for External Secrets 
# Note: These are created in the namespace where the secrets will be used
#       A separate role is created for each namespace
#       This provides granular control over what secrets can be accessed

# argocd external secrets
eksctl create iamserviceaccount --cluster=$CLUSTER_NAME \
  --name=external-secrets \
  --namespace=jk8s \
  --attach-policy-arn=$ARGOCD_SECRETS_ROLE_ARN \
  --override-existing-serviceaccounts \
  --approve

# jaeger external secrets
eksctl create iamserviceaccount --cluster=$CLUSTER_NAME \
  --name=external-secrets \
  --namespace=jaeger \
  --attach-policy-arn=$JAEGER_SECRETS_ROLE_ARN \
  --override-existing-serviceaccounts \
  --approve

# app external secrets
eksctl create iamserviceaccount --cluster=$CLUSTER_NAME \
  --name=external-secrets \
  --namespace=app \
  --attach-policy-arn=$APP_SECRETS_ROLE_ARN \
  --override-existing-serviceaccounts \
  --approve

# crossplane external secrets
eksctl create iamserviceaccount --cluster=$CLUSTER_NAME \
  --name=external-secrets \
  --namespace=crossplane-system \
  --attach-policy-arn=$CROSSPLANE_SECRETS_ROLE_ARN \
  --override-existing-serviceaccounts \
  --approve

# Add Helm Chart Repo
helm repo add jk8s https://jdotw.github.io/jk8s
helm repo update

# Install jk8s bootstrap Helm Chart
FQDN=$FQDN \
  ZONE_ID=$ZONE_ID \
  ES_DOMAIN=$ES_DOMAIN \
  ES_SECRET=$ES_SECRET \
  RDS_SECRET=$RDS_SECRET \
  RDS_HOST=$RDS_HOST \
  envsubst < values.yaml | helm upgrade jk8s jk8s/bootstrap --install -n jk8s --create-namespace -f -

# It always fails the first time so... run it agsin 
FQDN=$FQDN \
  ZONE_ID=$ZONE_ID \
  ES_DOMAIN=$ES_DOMAIN \
  ES_SECRET=$ES_SECRET \
  RDS_SECRET=$RDS_SECRET \
  RDS_HOST=$RDS_HOST \
  envsubst < values.yaml | helm upgrade jk8s jk8s/bootstrap --install -n jk8s --create-namespace -f -

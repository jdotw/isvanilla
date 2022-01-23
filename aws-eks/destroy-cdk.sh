#!/bin/zsh

export AWS_PROFILE=isvanilla
export CDK_DEFAULT_ACCOUNT=132602212048
export CDK_DEFAULT_REGION=ap-southeast-2

cd cdk

EKSCTL_STACKS=$(aws cloudformation list-stacks --query "StackSummaries[?starts_with(StackName, 'eksctl')].StackName" --stack-status-filter CREATE_COMPLETE)
echo $EKSCTL_STACKS | jq --args -c '.[]' - | while read i; do
  NAME=$(echo $i | sed s/\"//g)
  echo "Deleting eksctl Stack $NAME"
  aws cloudformation delete-stack --stack-name $NAME
done

ELB_NAMES=$(aws elb describe-load-balancers --query "LoadBalancerDescriptions[].LoadBalancerName")
echo $ELB_NAMES | jq --args -c '.[]' - | while read i; do
  NAME=$(echo $i | sed s/\"//g)
  echo "Deleting ELB $NAME" 
  aws elb delete-load-balancer --load-balancer-name $NAME
done

IAM_POLICY_ARNS=$(aws iam list-policies --scope Local --query "Policies[].Arn")
echo $IAM_POLICY_ARNS | jq --args -c '.[]' - | while read i; do
  ARN=$(echo $i | sed s/\"//g)
  echo "Deleting IAM Policy $ARN" 
  aws iam delete-policy --policy-arn $ARN
done

ES_DOMAIN_NAMES=$(aws opensearch list-domain-names --query "DomainNames[].DomainName")
echo $ES_DOMAIN_NAMES | jq --args -c '.[]' - | while read i; do
  NAME=$(echo $i | sed s/\"//g)
  echo "Deleting ES/OS Domain $NAME" 
  aws opensearch delete-domain --domain-name $NAME > /dev/null
done


exit 0

cdk destroy EKSStack --force 


#!/bin/zsh


cd cdk
aws --profile isvanilla iam create-service-linked-role --aws-service-name es.amazonaws.com
AWS_PROFILE=isvanilla \
  CDK_DEFAULT_ACCOUNT=132602212048 \
  CDK_DEFAULT_REGION=ap-southeast-2 \
  cdk deploy --all


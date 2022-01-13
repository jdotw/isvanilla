#!/bin/zsh

AWS_PROFILE=test-cluster
CDK_DEFAULT_ACCOUNT=132602212048
CDK_DEFAULT_REGION=ap-southeast-2
STACK_NAME=isvanilla

cd cdk
STACK_NAME=$STACK_NAME cdk deploy --all


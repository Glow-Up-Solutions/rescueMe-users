#!/bin/bash
set -eo pipefail
ENV="${EnvStageName:=dev}"
REGION=${Region:=us-west-2}
STACK_NAME="rescue-me-users-apis"
ARTIFACT_BUCKET="aws-sam-cli-managed-default-samclisourcebucket-lyqa21gm95pw"

# sam build

# sam deploy \
#     --force-upload\
#     --stack-name $STACK_NAME\
#     --template-file template.yaml\
#     --s3-bucket $ARTIFACT_BUCKET 
#     --region $REGION\
#     --capabilities CAPABILITY_IAM\
#     --confirm-changeset\
#     --debug

# Other test - from oficial documentation...
# cd register
# GOOS=linux go build main.go
# cd ../
# sam package \
#     --template-file template.yaml \
#     --s3-bucket $ARTIFACT_BUCKET \
#     --region $REGION
    # --output-template-file out.yaml
# aws cloudformation deploy \
#     --template-file out.yaml \
#     --stack-name $STACK_NAME \
#     --capabilities CAPABILITY_NAMED_IAM
sam build

sam deploy \
    --force-upload\
    --stack-name $STACK_NAME\
    --template-file template.yaml\
    --resolve-s3\
    --region $REGION\
    --capabilities CAPABILITY_IAM\
    --confirm-changeset\
    --debug



# test2
    # sam deploy \
    # --force-upload\
    # --stack-name $STACK_NAME\
    # --template-file template.yaml\
    # --s3-bucket $ARTIFACT_BUCKET 
    # --region $REGION\
    # --capabilities CAPABILITY_IAM\
    # --confirm-changeset\
    # --debug
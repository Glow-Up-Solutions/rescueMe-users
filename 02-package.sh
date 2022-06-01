#!/bin/bash
set -eo pipefail

REGION=${Region:=us-west-2}


sam package \
    --template-file template.yaml \
    --resolve-s3 \
    --region $REGION



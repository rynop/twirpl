#!/bin/bash

MY_PATH="`dirname \"$0\"`"
cd $MY_PATH

cd cmd/twirpl-webservices/
env GOOS=linux go build -o bin/main
rc=$?; if [[ $rc != 0 ]]; then exit $rc; fi

cd $MY_PATH/bin
zip deployment.zip main
mv deployment.zip /tmp/
aws lambda update-function-code --function-name $LAMBDA_NAME --zip-file fileb:///tmp/deployment.zip

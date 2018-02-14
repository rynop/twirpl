# Twirpl

Boilerplate/guide for creating 100% serverless [Twirp](https://github.com/twitchtv/twirp) APIs via AWS Lambda and API Gateway.  Supports both JSON and protobuf.

Please read my [blog post](https://rynop.com/2018/01/23/twirpl-twirp-go-framework-running-completely-serverless/) on why Twirp+APIG+Lambda are a powerful match for creating web APIs.

This project layout is based on [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

## Setup

* [Learn](https://blog.twitch.tv/twirp-a-sweet-new-rpc-framework-for-go-5f2febbf35f#a99f) about Twirp, then [Install Twirp](https://github.com/twitchtv/twirp/wiki) WITH retool.  Make sure to add `$GOPATH/bin` to your $PATH
* `retool add github.com/golang/dep/cmd/dep origin/master`
* `retool do dep init`
* Auto-generate the code:
```
retool do protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. ./rpc/publicservices/service.proto 
retool do protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. ./rpc/adminservices/service.proto 
```
* For this project, the interface implementations have been hand created in `pkg/`. Take a look.
* Build/package/create/deploy your lambda using a [lambda execution role](https://docs.aws.amazon.com/lambda/latest/dg/intro-permission-model.html#lambda-intro-execution-role):
```
#Run these from cmd/twirpl-webservices. I use fish shell, modify for your shell
env GOOS=linux go build -o bin/main
cd bin; zip deployment.zip main
aws lambda create-function \
--region us-east-1 \
--function-name TwirplTest \
--zip-file fileb://./deployment.zip \
--runtime go1.x \
--tracing-config Mode=Active \
--role arn:aws:iam::<account-id>:role/<lambda execution role> \
--handler main
rm deployment.zip main
```
* Use AWS Lambda console to assign a API Gateway trigger.  Make sure to choose `security` of `open`.
* After APIG created, login into APIG console and delete all the resources.  Add a `{proxy+}` with an `ANY` under it.  Hookup a lambda proxy integration to your lambda.  Should look something like this:
![APIG](https://rynop.files.wordpress.com/2018/01/screen-shot-2018-01-23-at-9-49-28-am.png?w=1566)
* Setup APIG to handle `application/protobuf` as a binary:
![APIG bin](https://rynop.files.wordpress.com/2018/01/screen-shot-2018-01-22-at-3-20-18-pm.png?w=1848)
* Deploy your APIG stage (copy down the APIG invocation URL)
* Modify [.gitignore](.gitignore) because you should check in `_tools` and `vendor` to real projects

## Test your endpoints

1. Test locally using JSON:
```
docker build -f build/Dockerfile -t twirpl .
docker run -p 8080:8080 twirpl
#in another terminal tab:
curl -H 'Content-Type:application/json' -d '{"term":"wahooo"}' http://localhost:8080/twirp/com.rynop.twirpl.publicservices.Image/CreateGiphy
```
2. Test APIG endpoint using JSON:
```
#kill running docker container
#comment out http.ListenAndServe(":8080", mux) and un-comment log.Fatal(gateway.ListenAndServe("", mux)) in twirpl.go. Save file.
env LAMBDA_NAME="TwirplTest" ./deploy.sh
curl -H 'Content-Type:application/json' -d '{"term":"wahooo"}' https://<yourAPIG>.execute-api.us-east-1.amazonaws.com/prod/twirp/com.rynop.twirpl.publicservices.Image/CreateGiphy
```

### Testing protobuf

> NOTE: binary support in APIG requires existence of `Accept: application/protobuf` request header

See test case at [twirpl_test.go](./twirpl_test.go) for how to use the go client as well as sending protobuf requests.  

## CI / Automation

[This blog](https://aws.amazon.com/blogs/compute/announcing-go-support-for-aws-lambda/) has a nice writeup for doing most of the above steps via CodePipeline and CodeBuild.  I have included [buildspec.yml](buildspec.yml) and [template.yml](./template.yml) to get you started.

## Quick iteration testing on AWS

Check out [deploy.sh](./deploy.sh) for a quick way to test changes to your lambda code in AWS (lambda must already exist)

```
env LAMBDA_NAME="TwirplTest" ./deploy.sh
```

## Javascript client

A quick proof of concept can be seen at [twirpl_test.js](./twirpl_test.js)

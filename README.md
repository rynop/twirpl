# Twirpl

Boilerplate/guide for creating 100% serverless [Twirp](https://blog.twitch.tv/twirp-a-sweet-new-rpc-framework-for-go-5f2febbf35f#a99f) via AWS Lambda and API Gateway.  Supports both JSON and protobuf.

Please read my [blog post]() on why Twirp+APIG+Lambda are a powerful match for creating web APIs.

## Setup

1. Install [Twirp](https://github.com/twitchtv/twirp/wiki)
1. Two sample high level services are inside `rpc` dir.
1. Auto-generate your code:
```
retool do protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. ./rpc/blog/service.proto 
retool do protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. ./rpc/image/service.proto 
```
4. Interface implementations for the services have already been created in `internal/`. Check them out.
5. Build/package/create/deploy your lambda using a [lambda execution role](https://docs.aws.amazon.com/lambda/latest/dg/intro-permission-model.html#lambda-intro-execution-role):
```
#from proj root. I use fish shell, modify for your shell
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
cd ..
```
6. Use AWS Lambda console to assign a API Gateway trigger.  Make sure to choose `security` of `open`.
7. After APIG created, login into APIG console and delete all the resources.  Add a `{proxy+}` with an `ANY` under it.  Hookup a lambda proxy integration to your lambda.  Should look something like this:
![APIG](https://rynop.files.wordpress.com/2018/01/screen-shot-2018-01-23-at-9-49-28-am.png?w=1566)
8. Setup APIG to handle `application/protobuf` as a binary:
![APIG bin](https://rynop.files.wordpress.com/2018/01/screen-shot-2018-01-22-at-3-20-18-pm.png?w=1848)
9. Deploy your APIG stage (copy down the APIG invocation URL)

## Test your endpoints

1. Test locally using JSON:
```
go build -o bin/main
./bin/main
#in another terminal tab:
curl -H 'Content-Type:application/json' -d '{"term":"wahooo"}' http://localhost:8080/twirp/com.rynop.twirpl.image.Image/CreateGiphy
```
2. Hit APIG endpoint using JSON:
```
#kill ./bin/main that is running for local test
#comment out http.ListenAndServe(":8080", mux) and un-comment log.Fatal(gateway.ListenAndServe("", mux)) in twirpl.go. Save file.
env LAMBDA_NAME="TwirplTest" ./deploy.sh
curl -H 'Content-Type:application/json' -d '{"term":"wahooo"}' https://<your APIG>.execute-api.us-east-1.amazonaws.com/prod/twirp/com.rynop.twirpl.image.Image/CreateGiphy
```

### Testing protobuf

> NOTE: currently waiting on resolution to https://github.com/twitchtv/twirp/pull/46 before this will work.  You can patch your generated `service.twirp.go` files until this issue is resolved.  You can see the addtion on line 412 of each `service.twirp.go` files in this repo.

See test case at [twirpl_test.go](./twirpl_test.go) for how to use the go client as well as sending protobuf requests.  

## CI / Automation

[This blog](https://aws.amazon.com/blogs/compute/announcing-go-support-for-aws-lambda/) has a nice writeup for doing most of the above steps via CodePipeline and CodeBuild.  I have included [buildspec.yml](buildspec.yml) and [template.yml](./template.yml) to get you started.

## Quick iteration testing on AWS

Check out [deploy.sh](./deploy.sh) for a quick way to test changes to your lambda code in AWS (lambda must already exist)

```
env LAMBDA_NAME="TwirplTest" ./deploy.sh
```

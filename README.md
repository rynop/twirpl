# Twirpl

100% serverless [Twirp](https://blog.twitch.tv/twirp-a-sweet-new-rpc-framework-for-go-5f2febbf35f#a99f) via AWS Lambda and API Gateway.

Please read my [blog post]() for why I think Twirp+APIG+Lambda are a powerful match for creating web APIs.

## Install & Configure

1. Install [Twirp](https://github.com/twitchtv/twirp/wiki) (look at bottom of page).
1. Download or fork & clone this repo.
1. Make a folder for each of your high level services inside `rpc`.  Use the ones there as examples, delete when done.
1. Auto-generate your code:
```
retool do protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. \
./rpc/blog/service.proto \
./rpc/image/service.proto 
```
5. Build/package/create/deploy your lambda using a [lambda execution role](https://docs.aws.amazon.com/lambda/latest/dg/intro-permission-model.html#lambda-intro-execution-role):

```
GOOS=linux go build -o main
zip deployment.zip main
aws lambda create-function \
--region us-east-1 \
--function-name TwirplTest \
--zip-file fileb://./deployment.zip \
--runtime go1.x \
--tracing-config Mode=Active
--role arn:aws:iam::<account-id>:role/<role> \
--handler main
rm deployment.zip main
```
6. Use AWS Lambda console to assign a API Gateway trigger.  Make sure to choose `security` of `open`.
7. After APIG created, go into APIG console and delete all the resources.  Add an `ANY` at root, and a `{proxy+}` with an `ANY` under it.  This will route all requests to your lambda.  Should look something like this:
![APIG](https://rynop.files.wordpress.com/2018/01/screen-shot-2018-01-22-at-4-44-47-pm.png?w=2720)
8. Setup APIG to handle `application/protobuf` as a binary:
![APIG bin](https://rynop.files.wordpress.com/2018/01/screen-shot-2018-01-22-at-3-20-18-pm.png?w=1848)
9. Deploy your APIG stage (copy down the APIG invocation URL).
10. Run `go test` to test protobuf.  See [twirpl_test.go](./twirpl_test.go)

## CI / Automation

[This blog](https://aws.amazon.com/blogs/compute/announcing-go-support-for-aws-lambda/) has a nice writeup for how to most of the above steps via CodePipeline and CodeBuild.  I have included [buildspec.yml](buildspec.yml) and [template.yml](./template.yml) to get you started.

## Quick iteration testing on AWS

Check out [deploy.sh](./deploy.sh) for a quick way to test changes to your lambda code in AWS (lambda must already exist)

```
env LAMBDA_NAME="YOUR-staging-Function-CHYO0FUJW8CU" ./deploy.sh
```

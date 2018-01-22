# Twirpl

100% serverless [Twirp](https://blog.twitch.tv/twirp-a-sweet-new-rpc-framework-for-go-5f2febbf35f#a99f) via AWS Lambda and API Gateway.

Please read my [blog post]() for why I think Twirp+APIG+Lambda are a powerful match for creating web APIs.

## Install & Configure

1. Install [Twirp](https://github.com/twitchtv/twirp/wiki) (look at bottom of page).
1. Download or fork & clone this repo.
1. 

## Building

### To auto-generate code:

```
retool do protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. \
./rpc/blog/service.proto \
./rpc/image/service.proto 
```

### Quick iteration testing on AWS

Check out [deploy.sh](./deploy.sh) for a quick way to test changes to your lambda code in AWS

```
env LAMBDA_NAME="YOUR-staging-Function-CHYO0FUJW8CU" ./deploy.sh
```
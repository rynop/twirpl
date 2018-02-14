FROM golang:1.9.3

ARG GITHUB_ORG=rynop
ARG REPO=twirpl
ARG CODE_PATH=cmd/twirpl-webservices

ADD . /go/src/github.com/${GITHUB_ORG}/${REPO}
WORKDIR /go/src/github.com/${GITHUB_ORG}/${REPO}/${CODE_PATH}
RUN go build -o main .

CMD ["./main"]
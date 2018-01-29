FROM golang:1.9.3

WORKDIR /go/src/github.com/rynop/twirpl
ADD . /go/src/github.com/rynop/twirpl
RUN go get -v
RUN go build -o bin/main
CMD ["/go/src/github.com/rynop/twirpl/bin/main"]
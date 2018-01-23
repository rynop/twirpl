package main_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/rynop/twirpl/rpc/image"
	"github.com/twitchtv/twirp"
)

func TestHandler(t *testing.T) {
	//URL for APIG
	urlPrefix := "https://co5gwn0iah.execute-api.us-east-1.amazonaws.com/prod"

	//URL for local testing
	// urlPrefix := "http://localhost:8080"

	// Un-comment for testing locally
	client := image.NewImageProtobufClient(urlPrefix, &http.Client{})

	resp, err := client.CreateGiphy(context.Background(), &image.Search{Term: "wahooo"})
	if err == nil {
		fmt.Println(resp)
	} else {
		if twerr, ok := err.(twirp.Error); ok {
			switch twerr.Code() {
			case twirp.InvalidArgument:
				fmt.Println("Oh no " + twerr.Error())
			default:
				fmt.Println(twerr.Error())
			}
		}
		return
	}
}

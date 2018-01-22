package main_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/rynop/twirpl/rpc/haberdasher"
	"github.com/twitchtv/twirp"
)

func TestHandler(t *testing.T) {
	//Client for APIG
	client := haberdasher.NewHaberdasherProtobufClient("https://XXX.execute-api.us-east-1.amazonaws.com/Stage", &http.Client{})

	// Un-comment for testing locally
	// client := haberdasher.NewHaberdasherProtobufClient("http://localhost:8080", &http.Client{})

	resp, err := client.MakeHat(context.Background(), &haberdasher.Size{Inches: 6})
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

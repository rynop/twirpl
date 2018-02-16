package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rynop/twirpl/pkg/authn"
	"github.com/rynop/twirpl/pkg/blogserver"
	"github.com/rynop/twirpl/pkg/imageserver"
	"github.com/rynop/twirpl/rpc/publicservices"
)

func addHeadersToContext(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, "Authorization", r.Header.Get("Authorization"))
		ctx = context.WithValue(ctx, "X-From-CDN", r.Header.Get("X-From-CDN"))
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}

func main() {
	//Dump all env vars
	for _, pair := range os.Environ() {
		fmt.Println(pair)
	}

	// You can use any mux you like
	mux := http.NewServeMux()

	authnHooks := authn.NewAuthNServerHooks()

	blogServerHandler := publicservices.NewBlogServer(&blogserver.Server{}, authnHooks) //authnHooks reads someKey from ctx
	wrappedBlogHandler := addHeadersToContext(blogServerHandler)
	mux.Handle(publicservices.BlogPathPrefix, wrappedBlogHandler)

	imageHandler := publicservices.NewImageServer(&imageserver.Server{}, authnHooks)
	wrappedImageHandler := addHeadersToContext(imageHandler)
	mux.Handle(publicservices.ImagePathPrefix, wrappedImageHandler)

	appStage, _ := os.LookupEnv("APP_STAGE")
	mux.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong. stage:"+appStage)
		fmt.Fprintln(w, r.Header)
	})

	// Un-comment below to test locally
	listenPort, exists := os.LookupEnv("LISTEN_PORT")
	if !exists {
		listenPort = "8080"
	}

	log.Print("Listening on " + listenPort + " in stage " + appStage + " docker image: --CodeImage--")
	//Comment out below if serving over lambda
	http.ListenAndServe(":"+listenPort, mux)

	// Un-comment below before deploying to Lambda
	// log.Fatal(gateway.ListenAndServe("", mux))
}

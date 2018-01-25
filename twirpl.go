package main

import (
	"fmt"
	"net/http"

	"github.com/rynop/twirpl/internal/blogserver"
	"github.com/rynop/twirpl/internal/imageserver"
	"github.com/rynop/twirpl/rpc/publicservices"
)

func main() {
	// You can use any mux you like
	mux := http.NewServeMux()

	//&blogserver.Server{} implements Blog interface
	blogHandler := publicservices.NewBlogServer(&blogserver.Server{}, nil)
	// The generated code includes a const, <ServiceName>PathPrefix, which
	// can be used to mount your service on a mux.
	mux.Handle(publicservices.BlogPathPrefix, blogHandler)

	imageHandler := publicservices.NewImageServer(&imageserver.Server{}, nil)
	mux.Handle(publicservices.ImagePathPrefix, imageHandler)

	mux.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})

	// Un-comment below to test locally
	http.ListenAndServe(":8080", mux)

	// Un-comment below before deploying to Lambda
	// log.Fatal(gateway.ListenAndServe("", mux))
}

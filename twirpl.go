package main

import (
	"net/http"

	"github.com/rynop/twirpl/internal/blogserver"
	"github.com/rynop/twirpl/internal/imageserver"
	"github.com/rynop/twirpl/rpc/blog"
	"github.com/rynop/twirpl/rpc/image"
)

func main() {
	// You can use any mux you like
	mux := http.NewServeMux()

	//&blogserver.Server{} implements Blog interface
	blogHandler := blog.NewBlogServer(&blogserver.Server{}, nil)
	// The generated code includes a const, <ServiceName>PathPrefix, which
	// can be used to mount your service on a mux.
	mux.Handle(blog.BlogPathPrefix, blogHandler)

	imageHandler := image.NewImageServer(&imageserver.Server{}, nil)
	mux.Handle(image.ImagePathPrefix, imageHandler)

	// Un-comment below to test locally
	http.ListenAndServe(":8080", mux)

	// Un-comment below before deploying to Lambda
	// log.Fatal(gateway.ListenAndServe("", mux))
}

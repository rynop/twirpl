package main

import (
	"log"
	"net/http"

	"github.com/apex/gateway"
	"github.com/rynop/twirpl/internal/haberdasherserver"
	"github.com/rynop/twirpl/rpc/haberdasher"
)

func main() {
	// You can use any mux you like
	mux := http.NewServeMux()

	//&haberdasherserver.Server{} implements Haberdasher interface
	haberdasherHandler := haberdasher.NewHaberdasherServer(&haberdasherserver.Server{}, nil)
	// The generated code includes a const, <ServiceName>PathPrefix, which
	// can be used to mount your service on a mux.
	mux.Handle(haberdasher.HaberdasherPathPrefix, haberdasherHandler)

	log.Fatal(gateway.ListenAndServe("", mux))
	// http.ListenAndServe(":8080", mux)
}

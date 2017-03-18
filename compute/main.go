package main

import (
	"github.com/stupschwartz/qubit/server/api"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServeTLS(":8443", "server.crt", "server.key", api.Handlers()))
}

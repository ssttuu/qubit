package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServeTLS(":8443", "server.crt", "server.key", nil))
}

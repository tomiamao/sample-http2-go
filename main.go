package main

import (
	"log"
	"net/http"
)

func main() {
	// Create a server on port 5432
	srv := &http.Server{Addr: ":5432", Handler: http.HandlerFunc(handle)}
	log.Printf("Serving on https://0.0.0.0:5432")
	log.Fatal(srv.ListenAndServeTLS("*.simplifiednetworks.co.chained.crt", "*.simplifiednetworks.co.key"))
}

func handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got connection from: %s using HTTP protocol: %s", r.RemoteAddr, r.Proto)
	// Send a message back to the client
	w.Write([]byte("Testing!!!"))
}

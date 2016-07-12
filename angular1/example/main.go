package main

import (
	"net/http"
	"log"
	"goTransport/goTransport"
)

var chttp = http.NewServeMux()

func main() {

	http.Handle("/", http.FileServer(http.Dir("../")))
	transport := goTransport.NewTransport("/ws")
	transport.SetRPCMethod("ping", ping)

	log.Print("goTransport server spawning at port: 8081/example")
	http.Handle("/ws/", transport.HttpHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}


func ping(parameters []interface{}) {
	log.Print("called jaja")
}
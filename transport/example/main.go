package main

import (
	"net/http"
	"log"
	"github.com/iain17/goTransport/transport"
	"errors"
	"github.com/iain17/goTransport/transport/lib/interfaces"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("../../")))
	transporter := transport.New("/ws")
	transporter.Method("ping", ping)

	log.Print("goTransport server spawning at port: 8081")
	log.Print("Angular 1 example available at: localhost:8081/angular1/example/")
	http.Handle("/ws/", transporter.GetHttpHandler())
	log.Fatal(http.ListenAndServe(":8081", nil))
}


func ping(session interfaces.Session, message string) (string, error) {
	log.Print("called", message)
	log.Print(session)
	return "bar", errors.New("test")
}
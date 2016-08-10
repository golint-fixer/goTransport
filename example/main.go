package main

import (
	"errors"
	"github.com/iain17/goTransport"
	"github.com/iain17/goTransport/lib/interfaces"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("../goTransport-client/")))
	transporter := goTransport.New("/ws")
	transporter.Method("ping", ping)

	log.Print("goTransport server spawning at port: 8081")
	log.Print("Angular 1 example available at: localhost:8081/src/example/")
	http.Handle("/ws/", transporter.GetHttpHandler())
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func ping(session interfaces.ICallableSession, message string) (string, error) {
	log.Print("called with parameter: ", message)

	log.Print("Calling example method client side.")
	promise := session.Call("example", []interface{}{
		"A test",
		1337,
	}, 0)
	promise.OnSuccess(func(v interface{}) {
		log.Print("Success: ", v)
	}).OnFailure(func(v interface{}) {
		log.Print("Failure: ", v)
	})

	return "bar", errors.New("test")
}

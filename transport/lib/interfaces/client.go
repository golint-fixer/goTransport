package interfaces

import (
	"net/http"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
)

type Client interface {
	GetHttpHandler() http.Handler
	Listen(socket sockjs.Session)
	Method(name string, method CallableMethod)
	GetMethod(name string) CallableMethod
}

type Session interface {
	Messaged(data string)
	Send(message string)
	GetClient() Client
}
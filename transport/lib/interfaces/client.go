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
	CallableSession
	Messaged(data string) error
	Send(message string)
	GetClient() Client
	GetCurrentId() uint64
	SetCurrentId(id uint64)
	IncrementCurrentId()
}

type CallableSession interface {
	Call(name string, parameters []interface{})
}
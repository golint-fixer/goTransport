package interfaces

import (
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"net/http"
)

type IClient interface {
	GetHttpHandler() http.Handler
	Listen(socket sockjs.Session)
	Method(name string, method CallableMethod)
	GetMethod(name string) CallableMethod
}

type ISession interface {
	ICallableSession
	Messaged(data string) error
	Send(message string)
	GetClient() IClient
	GetCurrentId() uint64
	SetCurrentId(id uint64)
	IncrementCurrentId()
}

type ICallableSession interface {
	Call(name string, parameters []interface{})
}

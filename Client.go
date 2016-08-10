package goTransport

import (
	"github.com/iain17/goTransport/lib"
	"github.com/iain17/goTransport/lib/interfaces"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"net/http"
	"sync"
)

type client struct {
	HttpHandler   http.Handler
	methods_mutex *sync.Mutex
	methods       map[string]interfaces.CallableMethod
	connectFunc   func(interfaces.ICallableSession)
}

func New(prefix string, connectFunc func(interfaces.ICallableSession)) interfaces.IClient {
	client := &client{
		methods_mutex: new(sync.Mutex),
		methods:       make(map[string]interfaces.CallableMethod),
		connectFunc:   connectFunc,
	}
	client.HttpHandler = sockjs.NewHandler(prefix, sockjs.DefaultOptions, client.Listen)
	return client
}

func (client *client) Listen(socket sockjs.Session) {
	session := lib.NewSession(socket, client)
	client.connectFunc(session)
	for {
		if msg, err := socket.Recv(); err == nil {
			go session.Messaged(msg)
			continue
		}
		break
	}
}

func (client *client) GetHttpHandler() http.Handler {
	return client.HttpHandler
}

func (client *client) Method(name string, method interfaces.CallableMethod) {
	client.methods_mutex.Lock()
	client.methods[name] = method
	client.methods_mutex.Unlock()
}

func (client *client) GetMethod(name string) interfaces.CallableMethod {
	if a, ok := client.methods[name]; ok {
		return a
	}

	return nil
}

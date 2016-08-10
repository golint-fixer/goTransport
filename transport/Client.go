package transport

import (
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"net/http"
	"github.com/iain17/goTransport/transport/lib/interfaces"
	"sync"
	"github.com/iain17/goTransport/transport/lib"
)

type client struct {
	HttpHandler http.Handler
	methods_mutex *sync.Mutex
	methods map[string]interfaces.CallableMethod
}

func New(prefix string) interfaces.Client {
	client := &client{
		methods_mutex: new(sync.Mutex),
		methods: make(map[string]interfaces.CallableMethod),
	}
	client.HttpHandler = sockjs.NewHandler(prefix, sockjs.DefaultOptions, client.Listen)
	return client
}

func (client *client) Listen(socket sockjs.Session) {
	session := lib.NewSession(socket, client)
	for {
		if msg, err := socket.Recv(); err == nil {
			go session.Messaged(msg)
			continue
		}
		break
	}
}

func (client *client) GetHttpHandler() http.Handler {
	return client.HttpHandler;
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
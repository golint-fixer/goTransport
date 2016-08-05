package transport

import (
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"net/http"
	"github.com/iain17/goTransport/transport/lib/interfaces"
	"github.com/iain17/goTransport/transport/lib/messageType"
	"sync"
	"github.com/iain17/goTransport/transport/lib/Session"
)

func init() {
	messageType.Init()
}

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
	manager := Session.NewSession(socket, client)
	for {
		if msg, err := socket.Recv(); err == nil {
			manager.Messaged(msg)
		}
		break
	}
}

func (client *client) GetHttpHandler() http.Handler {
	return client.HttpHandler;
}
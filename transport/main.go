package transport

import (
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"net/http"
	"github.com/iain17/goTransport/transport/lib"
	"github.com/iain17/goTransport/transport/lib/messageType"
)

type Client interface {
	Method(name string, method lib.CallableMethod)
	GetHttpHandler() http.Handler
}

type client struct {
	HttpHandler http.Handler
	messageManager lib.MessageManager
}

func init() {
	messageType.Init()
}

func New(prefix string) Client {
	transport := &client{
		messageManager: lib.NewMessageManager(),
	}
	transport.HttpHandler = sockjs.NewHandler(prefix, sockjs.DefaultOptions, transport.messageManager.Listen)
	return transport
}

func (transport *client) GetHttpHandler() http.Handler {
	return transport.HttpHandler;
}
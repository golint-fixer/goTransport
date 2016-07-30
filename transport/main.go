package transport

import (
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"net/http"
	"github.com/iain17/goTransport/transport/lib/MessageManager"
	"github.com/iain17/goTransport/transport/lib/interfaces"
	"github.com/iain17/goTransport/transport/lib/messageType"
)

func init() {
	messageType.Init()
}

type Client interface {
	Method(name string, method interfaces.CallableMethod)
	GetHttpHandler() http.Handler
}

type client struct {
	HttpHandler http.Handler
	messageManager interfaces.MessageManager
}

func New(prefix string) Client {
	transport := &client{
		messageManager: MessageManager.NewMessageManager(),
	}
	transport.HttpHandler = sockjs.NewHandler(prefix, sockjs.DefaultOptions, transport.messageManager.Listen)
	return transport
}

func (transport *client) GetHttpHandler() http.Handler {
	return transport.HttpHandler;
}
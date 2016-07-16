package transport

import (
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"net/http"
	"github.com/iain17/goTransport/transport/Message"
)

type Transport interface {
	Method(name string, method RPCMethod)
	GetHttpHandler() http.Handler
}

type transport struct {
	HttpHandler http.Handler
	methods map[string]RPCMethod
	messageManager Message.MessageManager
}

func New(prefix string) Transport {
	transport := &transport{
		methods: make(map[string]RPCMethod),
		messageManager: Message.NewMessageManager(),
	}
	transport.HttpHandler = sockjs.NewHandler(prefix, sockjs.DefaultOptions, transport.messageManager.Listen)
	return transport
}

func (transport *transport) Method(name string, method RPCMethod) {

}

func (transport *transport) GetHttpHandler() http.Handler {
	return transport.HttpHandler;
}
package transport

import (
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"net/http"
	"encoding/json"
	"log"
)

type Transport struct {
	HttpHandler http.Handler
	MessageManager *MessageManager
}

func NewTransport(prefix string) *Transport {
	transport := &Transport{
		MessageManager: NewMessageManager()
		//methods: make(map[string]RPCMethod),
	}
	transport.HttpHandler = sockjs.NewHandler(prefix, sockjs.DefaultOptions, transport.listen)
	return transport
}
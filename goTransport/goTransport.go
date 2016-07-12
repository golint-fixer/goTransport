package goTransport

import (
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"net/http"
	"encoding/json"
	"log"
)

func init() {
	initStorage()
}

type Transport struct {
	HttpHandler http.Handler
	methods map[string]RPCMethod
}

func NewTransport(prefix string) *Transport {
	transport := &Transport{
		methods: make(map[string]RPCMethod),
	}
	transport.HttpHandler = sockjs.NewHandler(prefix, sockjs.DefaultOptions, transport.listen)
	return transport
}

func (transport *Transport) listen(session sockjs.Session) {
	for {
		if msg, err := session.Recv(); err == nil {
			var message Message
			err = json.Unmarshal([]byte(msg), &message)
			if err != nil {
				log.Print(err)
			}
			message.Transport = transport
			message.Session = session
			message.Json = []byte(msg)
			message.Call()
			continue
		}
		break
	}
}
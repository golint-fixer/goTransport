package transport

import "github.com/iain17/goTransport/transport/lib"

func (transport *client) Method(name string, method lib.CallableMethod) {
	transport.messageManager.SetMethod(name, method)
}
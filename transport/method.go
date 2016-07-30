package transport

import (
	"github.com/iain17/goTransport/transport/lib/interfaces"
)

func (transport *client) Method(name string, method interfaces.CallableMethod) {
	transport.messageManager.SetMethod(name, method)
}
package transport

import (
	"github.com/iain17/goTransport/transport/lib/interfaces"
)

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
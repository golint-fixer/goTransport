package goTransport

import (
	"sync"
	"log"
)

var methods_mutex = new(sync.Mutex)

type RPCMethod func(parameters []interface{})

func (transport *Transport) SetRPCMethod(name string, method RPCMethod) {
	methods_mutex.Lock()
	transport.methods[name] = method
	methods_mutex.Unlock()
	log.Print(transport.methods)
}

func (transport *Transport) getRPCMethod(name string) RPCMethod {
	if a, ok := transport.methods[name]; ok {
		return a
	}

	return nil
}
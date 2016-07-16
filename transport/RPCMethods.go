package transport

import (
	"sync"
)

var methods_mutex = new(sync.Mutex)

type RPCMethod interface{}

//func (transport *transport) Method(name string, method RPCMethod) {
//	methods_mutex.Lock()
//	transport.methods[name] = method
//	methods_mutex.Unlock()
//}

//func (transport *transport) getMethod(name string) RPCMethod {
//	if a, ok := transport.methods[name]; ok {
//		return a
//	}
//
//	return nil
//}
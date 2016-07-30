package MessageManager

import "github.com/iain17/goTransport/transport/lib/interfaces"

func (manager *messageManager) SetMethod(name string, method interfaces.CallableMethod) {
	manager.methods_mutex.Lock()
	manager.methods[name] = method
	manager.methods_mutex.Unlock()
}

func (manager *messageManager) GetMethod(name string) interfaces.CallableMethod {
	if a, ok := manager.methods[name]; ok {
		return a
	}

	return nil
}
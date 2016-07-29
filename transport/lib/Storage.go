package lib

type CallableMethod interface{}

func (manager *messageManager) SetMethod(name string, method CallableMethod) {
	manager.methods_mutex.Lock()
	manager.methods[name] = method
	manager.methods_mutex.Unlock()
}

func (manager *messageManager) GetMethod(name string) CallableMethod {
	if a, ok := manager.methods[name]; ok {
		return a
	}

	return nil
}
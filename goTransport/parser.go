package goTransport

import "sync"

//Parser
type Handler interface {
	Call()
}

type Parser struct {
	//Validate the message. and returns a handler
	Parse func(*Message, MessageType) (Handler, error)
	ReturnMessageType MessageType
}

var parsers map[MessageType]*Parser
var parsers_mutex = new(sync.Mutex)

func initStorage() {
	parsers = make(map[MessageType]*Parser)
}

func SetParser(messageType MessageType, parser *Parser) {
	parsers_mutex.Lock()
	parsers[messageType] = parser
	parsers_mutex.Unlock()
}

func GetParser(messageType MessageType) *Parser {
	if a, ok := parsers[messageType]; ok {
		return a
	}

	return nil
}
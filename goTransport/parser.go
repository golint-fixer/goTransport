package goTransport

import "sync"

//Parser
type IHandler interface {
	Test123()
}

type Handler struct {
	Message Message
}

type Parser struct {
	//Validate the message. and returns a handler
	Get func() IHandler
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

func GetHandler(messageType MessageType) *Parser {
	if a, ok := parsers[messageType]; ok {
		return a
	}

	return nil
}
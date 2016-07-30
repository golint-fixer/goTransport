package interfaces

import "gopkg.in/igm/sockjs-go.v2/sockjs"

type CallableMethod interface{}

type MessageType int

type IMessage interface {
	SetId(id uint64)
	GetId() uint64
	GetType() MessageType
	//setReply()
	Validate(manager MessageManager, session sockjs.Session) error
	Run(manager MessageManager, session sockjs.Session) error
	Reply(replyMessage IMessage, manager MessageManager, session sockjs.Session)
	//serialize() string
}

type MessageManager interface {
	Listen(session sockjs.Session)
	SetMethod(name string, method CallableMethod)
	GetMethod(name string) CallableMethod
	Send(message string, session sockjs.Session)
}
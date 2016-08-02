package interfaces

import "gopkg.in/igm/sockjs-go.v2/sockjs"

type CallableMethod interface{}

type MessageType int

type IMessage interface {
	Initialize(manager MessageManager, session *sockjs.Session)
	GetManager() MessageManager
	GetSession() *sockjs.Session

	SetId(id uint64)
	GetId() uint64
	GetType() MessageType
	//setReply()
	Validate() error
	Run() error
	Reply(replyMessage IMessage)
	Send()
	//serialize() string
}

type MessageManager interface {
	Listen(session sockjs.Session)
	SetMethod(name string, method CallableMethod)
	GetMethod(name string) CallableMethod
	Send(message string, session *sockjs.Session)
}
package interfaces

type CallableMethod interface{}

type MessageType int

type IMessage interface {
	Initialize(manager ISession)
	GetSession() ISession

	SetId(id uint64)
	GetId() uint64
	GetType() MessageType
	//setReply()
	Sending() error
	Received(previousMessage IMessage) error
	Reply(replyMessage IMessage)
	//Send()
	//serialize() string
}

package interfaces

type CallableMethod interface{}

type MessageType int

type IMessage interface {
	Initialize(manager Session)
	GetSession() Session

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
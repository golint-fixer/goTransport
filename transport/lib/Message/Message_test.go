package Message

import (
	"testing"
	"github.com/iain17/goTransport/transport/lib/interfaces"
)

type messageTest struct {
	Message
	Foo string `json:"foo"`
}

func NewMessageTest(foo string) *messageTest {
	return &messageTest{
		Message: NewMessage(MessageTypeTest),
		Foo: foo,
	}
}

func TestMessage_SetGetId(t *testing.T) {
	message := NewMessageTest("bar")
	expectedResult := uint64(1338)

	message.SetId(expectedResult)
	result := message.GetId()

	if result != expectedResult {
		t.Fatalf("Expected %d. Received: %d", expectedResult, result)
	}
}

func TestUnSerialize(t *testing.T) {
	messageType := interfaces.MessageType(12312)
	exampleMessage := NewMessage(messageType)
	Set(&exampleMessage)

	data := `12312{"type":0,"id":131,"name":"ping","parameters":["hai"]}`
	message := UnSerialize(data)

	if message == nil {
		t.Fatal("UnSerialize failed. Returned nil")
	}

	if message.GetId() != 131 {
		t.Fatalf("UnSerialize failed. Id expected to be 131 but received: %d", message.GetId())
	}
}



//func TestMessage_Send(t *testing.T) {
//	message := NewMessageTest("test")
//	message.Initialize()
//	message.Send()
//	data := serialize(message)
//	log.Print(data)
//}

//var validate_called = false
//func (message *messageTest) Validate() error {
//validate_called = true
//return nil
//}
//
//var run_called = false
//func (message *messageTest) Run() error {
//run_called = true
//return 	nil
//}
//
////func TestMessage_Start(t *testing.T) {
////	message := NewMessageTest("bar")
////	error := Start(message)
////
////	if error != nil {
////		t.Fatal(error)
////	}
////
////	if !validate_called {
////		t.Fatal("message.validate not called")
////	}
////
////	if !run_called {
////		t.Fatal("message.run not called")
////	}
////}
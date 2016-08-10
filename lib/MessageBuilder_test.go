package lib

import (
	"github.com/iain17/goTransport/lib/interfaces"
	"reflect"
	"testing"
)

type invalidMessage struct {
}

func TestMessageBuilder_build(t *testing.T) {
	exampleMessage := NewMessage(interfaces.MessageType(1337))
	result := build(reflect.TypeOf(&exampleMessage), `{"id":1,"type":1,"name":"ping","parameters":["hai"]}`)
	if result == nil {
		t.Fatal("It should have build this message")
	}

	result = build(reflect.TypeOf(&exampleMessage), `invalidJson`)
	if result != nil {
		t.Fatal("It should not have build this message. Definition isn't of Message interface type.")
	}

	result = build(reflect.TypeOf(invalidMessage{}), `{"id":1,"type":1,"name":"ping","parameters":["hai"]}`)
	if result != nil {
		t.Fatal("It should not have build this message. Definition isn't of Message interface type.")
	}
}

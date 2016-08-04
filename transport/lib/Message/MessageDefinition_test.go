package Message

import (
	"testing"
	"reflect"
	"github.com/iain17/goTransport/transport/lib/interfaces"
	"log"
)

func TestGet(t *testing.T) {
	messageType := interfaces.MessageType(1337)
	exampleMessage := NewMessage(messageType)
	Set(&exampleMessage)

	resultDefinition := Get(messageType)

	if resultDefinition == nil {
		t.Fatal("resultDefinition turned out to be an unexpected nil")
	}

	log.Print(resultDefinition.Elem().Name())

	if resultDefinition.Elem().Name() != reflect.TypeOf(exampleMessage).Name() {
		t.Fatalf("Expected resultDefinition to be %s, but in return got %s ", reflect.TypeOf(exampleMessage).Name(), resultDefinition.Elem().Name())
	}

	if Get(interfaces.MessageType(1338)) != nil {
		t.Fatalf("Expected to receive a nil value when requesting a non existing messageType.")
	}
}
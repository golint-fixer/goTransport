package lib

import (
	"github.com/iain17/goTransport/lib/interfaces"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"net/http"
	"testing"
)

//FakeClient
type fakeClient struct {
}

func (client *fakeClient) Listen(socket sockjs.Session) {

}

func (client *fakeClient) GetHttpHandler() http.Handler {
	return nil
}

func (client *fakeClient) Method(name string, method interfaces.CallableMethod) {

}

func ping(session interfaces.CallableSession, message string) {

}

func (client *fakeClient) GetMethod(name string) interfaces.CallableMethod {
	return ping
}

func TestSession_GetClient(t *testing.T) {
	client := &fakeClient{}
	session := NewSession(nil, client)
	if session.GetClient() != client {
		t.Fatal("GetClient doesn't return the same client")
	}
}

func TestSession_CurrentId(t *testing.T) {
	session := NewSession(nil, nil)
	if session.GetCurrentId() != 0 {
		t.Fatal("Current Id should start at 0")
	}
	session.IncrementCurrentId()
	if session.GetCurrentId() != 1 {
		t.Fatal("IncrementCurrentId should increment by +1")
	}
	session.SetCurrentId(123)
	if session.GetCurrentId() != 123 {
		t.Fatal("SetCurrentId should have set the currentId to 123")
	}
}

//FakeSocket
var SendCalled bool

type fakeSocket struct {
}

func (client *fakeSocket) ID() string {
	return "test"
}

func (client *fakeSocket) Recv() (string, error) {
	return "", nil
}

func (client *fakeSocket) Send(string) error {
	SendCalled = true
	return nil
}

func (client *fakeSocket) Close(status uint32, reason string) error {
	return nil
}

func TestSession_Messaged(t *testing.T) {
	client := &fakeClient{}
	socket := &fakeSocket{}
	session := NewSession(socket, client)

	SendCalled = false
	err := session.Messaged(`1` + headerDelimiter + `{"id":1,"type":1,"name":"ping","parameters":["hai"]}`)
	if err != nil {
		t.Fatal("Messaged should have accepted this message")
	}
	if SendCalled == false {
		t.Fatal("Messaged should've replied.")
	}

	SendCalled = false
	err = session.Messaged(`1` + headerDelimiter + `{"id":1,"type":9090909090}`)
	if err == nil {
		t.Fatal("Messaged should have stopped this message")
	}
	if SendCalled == false {
		t.Fatal("Messaged should've replied.")
	}

	SendCalled = false
	err = session.Messaged(`corrupt message`)
	if err == nil {
		t.Fatal("Messaged should have stopped this message")
	}
	if SendCalled == true {
		t.Fatal("Messaged shouldn't have replied.")
	}
}

func TestSession_Call(t *testing.T) {
	client := &fakeClient{}
	socket := &fakeSocket{}
	session := NewSession(socket, client)

	SendCalled = false
	session.Call("example", nil)
	if SendCalled != true {
		t.Fatal("Call should've sent a message requesting the method call.")
	}
}

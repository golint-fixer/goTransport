package goTransport

import (
	"reflect"
	"testing"
)

func testMethod() {

}

func TestClient_Method(t *testing.T) {
	client := New("", nil)
	client.Method("A test", testMethod)
	method := client.GetMethod("A test")
	if reflect.TypeOf(method) != reflect.TypeOf(testMethod) {
		t.Fatal("Method setting and getting failed.")
	}

	method = client.GetMethod("unknown")
	if method != nil {
		t.Fatal("Nil value should be returned when asking for a non existant method.")
	}
}

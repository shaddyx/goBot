package gobot

import (
	"fmt"
	"testing"
)

func TestRouter_AddHandler(t *testing.T) {
	called := false
	f := func(event string, msg IncomingChatMessage) *OutgoingChatMessage {
		fmt.Println("incoming:" + event)
		if event != "test" {
			t.Error("Error, message must be only test")
		}
		called = true
		return &OutgoingChatMessage{}
	}
	router := NewRouter()
	router.AddHandler("test", f)
	router.CallHandler("test", IncomingChatMessage{})
	router.CallHandler("test1", IncomingChatMessage{})
	if !called {
		t.Errorf("Not called")
	}
}

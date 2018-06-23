package gobot

import (
	"fmt"
	"testing"
)

func TestRouter_AddHandler(t *testing.T) {
	called := false
	f := func(msg IncomingChatMessage) *OutgoingChatMessage {
		fmt.Println("incoming:" + msg.Text)
		if msg.Text != "test" {
			t.Error("Error, message must be only test")
		}
		called = true
		return &OutgoingChatMessage{}
	}
	router := NewRouter()
	router.AddHandler("test", f)
	router.CallHandler(IncomingChatMessage{
		Text: "test",
	})
	router.CallHandler(IncomingChatMessage{
		Text: "test1",
	})
	if !called {
		t.Errorf("Not called")
	}
}

func TestRouter_AddRegexHandler(t *testing.T) {
	called := false
	f := func(msg IncomingChatMessage) *OutgoingChatMessage {
		fmt.Println("incoming:" + msg.Text)
		if msg.Text != "test" && msg.Text != "test1" {
			t.Error("Error, message must be only test")
		}

		called = true
		return &OutgoingChatMessage{}
	}
	router := NewRouter()
	router.AddRegexHandler("test", f)
	router.CallHandler(IncomingChatMessage{
		Text: "test",
	})
	router.CallHandler(IncomingChatMessage{
		Text: "test1",
	})
	router.CallHandler(IncomingChatMessage{
		Text: "zzzzzzzzzzzz",
	})
	if !called {
		t.Errorf("Not called")
	}
}

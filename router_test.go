package gobot

import (
	"fmt"
	"strconv"
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

func TestRouter_SetDefaultHandler(t *testing.T) {
	called := 0
	f := func(msg IncomingChatMessage) *OutgoingChatMessage {
		fmt.Println("incoming:" + msg.Text)
		if msg.Text != "test" && msg.Text != "test1" {
			t.Error("Error, message must be only test")
		}

		called++
		return &OutgoingChatMessage{}
	}
	router := NewRouter()
	router.SetDefaultHandler(f)
	router.AddHandler("zzzzzzzzzzzz", func(msg IncomingChatMessage) *OutgoingChatMessage {
		return &OutgoingChatMessage{}
	})
	router.CallHandler(IncomingChatMessage{
		Text: "test",
	})
	router.CallHandler(IncomingChatMessage{
		Text: "test1",
	})
	router.CallHandler(IncomingChatMessage{
		Text: "zzzzzzzzzzzz",
	})
	if called != 2 {
		t.Errorf("Called only " + strconv.Itoa(called) + " times")
	}
}

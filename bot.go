package gobot

import "reflect"

type AbstractBot struct {
	Token string
	BotInterface
	updates chan IncomingChatMessage
}

type BotInterface interface {
	Connect() error
	Send(message OutgoingChatMessage) error
	SendAnswer(message OutgoingChatMessage, incomingMessage *IncomingChatMessage) error
	Disconnect() error
	GetUpdates() chan IncomingChatMessage
}

func (b AbstractBot) String() string {
	return reflect.TypeOf(b).Name() + "[token=" + b.Token + "]"
}

func (b AbstractBot) Disconnect() error {
	return nil
}

func (b AbstractBot) GetUpdates() chan IncomingChatMessage {
	return b.updates
}

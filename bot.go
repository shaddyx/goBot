package gobot

import "reflect"

type AbstractBot struct {
	Token string
	BotInterface
	updates *UpdatesIterator
}

type BotInterface interface {
	Connect() error
	Send(message OutgoingChatMessage) error
	SendAnswer(message OutgoingChatMessage, incomingMessage *IncomingChatMessage) error
	GetUpdates() *UpdatesIterator
	Disconnect() error
}

func (b AbstractBot) GetUpdates() *UpdatesIterator {
	if b.updates == nil {
		b.updates = NewUpdatesUterator()
	}
	return b.updates
}

func (b AbstractBot) String() string {
	return reflect.TypeOf(b).Name() + "[token=" + b.Token + "]"
}

func (b AbstractBot) Disconnect() error {
	return nil
}

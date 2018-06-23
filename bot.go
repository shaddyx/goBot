package gobot

import "reflect"

type AbstractBot struct {
	Token   string
	updates *UpdatesIterator
}

type BotInterface interface {
	Connect() error
	Send(msg OutgoingChatMessage) error
	SendAnswer(msg OutgoingChatMessage, message IncomingChatMessage) error
	GetUpdates() *UpdatesIterator
	Disconnect() error
}

func (b *AbstractBot) GetUpdates() *UpdatesIterator {
	if b.updates == nil {
		b.updates = NewUpdatesUterator()
	}
	return b.updates
}

func (b *AbstractBot) String() string {
	return reflect.TypeOf(b).Name() + "[token=" + b.Token + "]"
}

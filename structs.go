package gobot

type KeyButton struct {
	Text string
}

type KeyboardRow struct {
	Buttons []KeyButton
}

func CreateRows(rows ...KeyboardRow) *[]KeyboardRow {
	return &rows
}

func CreateRow(rows ...KeyButton) KeyboardRow {
	return KeyboardRow{
		Buttons: rows,
	}
}

type OutgoingChatMessage struct {
	MessageId        string
	ReplyToMessageID string
	To               string
	Text             string
	Buttons          *[]KeyboardRow
}

type IncomingChatMessage struct {
	MessageId    string
	Text         string
	From         string
	FromUserName string
}

func NewIncomingChatMessage(messageId string, text string, from string, fromUserName string) IncomingChatMessage {
	return IncomingChatMessage{
		MessageId:    messageId,
		Text:         text,
		From:         from,
		FromUserName: fromUserName,
	}
}

package gobot

import (
	"fmt"
	"log"
	"strconv"

	"github.com/pkg/errors"
	"gopkg.in/telegram-bot-api.v4"
)

type TelegramBot struct {
	AbstractBot
	BotInterface
	nativeBot *tgbotapi.BotAPI
}

func (b *TelegramBot) Connect() error {
	log.Println(b.Token)
	b.updates = NewUpdatesUterator()
	log.Println("Initializing native telegram bot api")
	bot, err := tgbotapi.NewBotAPI(b.Token)
	if err != nil {
		return errors.Wrap(err, "Error initializing native telegram bot")
	}
	b.nativeBot = bot
	log.Println("Api initialized")

	go (func() {
		log.Println("starting processUpdates")
		updatesErr := b.processUpdates()
		if updatesErr != nil {
			log.Fatal(errors.Wrap(err, "Error listening updates"))
		}
	})()

	return nil
}

func (b *TelegramBot) getTelegramId(message OutgoingChatMessage) (int64, error) {
	return strconv.ParseInt(message.To, 10, 0)
}

func (b *TelegramBot) createReplyMarkup(buttons *[]KeyboardRow) *tgbotapi.ReplyKeyboardMarkup {
	rows := [][]tgbotapi.KeyboardButton{}

	for _, sourceButtonRow := range *buttons {
		row := []tgbotapi.KeyboardButton{}
		for _, sourceButton := range sourceButtonRow.Buttons {
			row = append(row, tgbotapi.NewKeyboardButton(sourceButton.Text))
		}
		rows = append(rows, row)
	}
	return &tgbotapi.ReplyKeyboardMarkup{
		Keyboard: rows,
	}
}

func (b *TelegramBot) Send(message OutgoingChatMessage) error {
	if message.To == "" {
		return errors.New("Error, message has no TO field:" + message.To)
	}
	id, err := b.getTelegramId(message)
	if err != nil {
		return errors.Wrap(err, "Error getting telegramId:"+message.To)
	}
	msg := tgbotapi.NewMessage(id, message.Text)
	if message.ReplyToMessageID != "" {
		incomingId, err := strconv.Atoi(message.ReplyToMessageID)
		if err != nil {
			return errors.Wrap(err, "Error parsing incoming message id:"+message.MessageId)
		}
		msg.ReplyToMessageID = incomingId
	}

	//
	//	TODO: optimize this, do not send markup, if already sent
	//
	if message.Buttons == nil || len(*message.Buttons) == 0 {
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	} else {
		msg.ReplyMarkup = b.createReplyMarkup(message.Buttons)
	}

	log.Println("Sending message...")
	_, sendErr := b.nativeBot.Send(msg)
	if sendErr != nil {
		return errors.Wrap(sendErr, "Error sending message:"+message.To)
	}
	log.Println("Message sent")
	return nil
}

func (b *TelegramBot) SendAnswer(message OutgoingChatMessage, incomingMessage *IncomingChatMessage) error {
	if incomingMessage == nil {
		panic("incomingMessage is nil")
	}

	if message.To == "" {
		message.To = incomingMessage.From
	}
	message.ReplyToMessageID = incomingMessage.MessageId
	b.Send(message)
	return nil
}

func (b *TelegramBot) processUpdates() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.nativeBot.GetUpdatesChan(u)
	if err != nil {
		return errors.Wrap(err, "Error getting native updates")
	}
	log.Println("Starting listening updates")
	for update := range updates {
		log.Println("Update incoming:" + fmt.Sprintf("%v", update))
		if update.Message == nil {
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		inMsgId := strconv.Itoa(update.Message.MessageID)
		inUserId := strconv.Itoa(update.Message.From.ID)
		b.updates.AddMessage(NewIncomingChatMessage(inMsgId, update.Message.Text, inUserId, update.Message.From.UserName))
	}
	return nil
}

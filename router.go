package gobot

import (
	"log"
	"regexp"
)

type RouterHandlerFunc func(msg IncomingChatMessage) *OutgoingChatMessage

type handlerWrapper struct {
	handler RouterHandlerFunc
}

type Router struct {
	handlers              map[string]RouterHandlerFunc
	regexHandlers         map[string]RouterHandlerFunc
	defaultHandlerWrapper *handlerWrapper
}

func NewRouter() Router {
	return Router{
		handlers:      make(map[string]RouterHandlerFunc),
		regexHandlers: make(map[string]RouterHandlerFunc),
		defaultHandlerWrapper: &handlerWrapper{
			handler: nil,
		},
	}
}

func (r Router) AddHandler(cmd string, f RouterHandlerFunc) {
	if r.handlers[cmd] != nil {
		panic("Handler already exists:" + cmd)
	}
	r.handlers[cmd] = f
}

func (r Router) AddRegexHandler(cmd string, f RouterHandlerFunc) {
	if r.regexHandlers[cmd] != nil {
		panic("Handler already exists:" + cmd)
	}
	r.regexHandlers[cmd] = f
}

func (r Router) SetDefaultHandler(f RouterHandlerFunc) {
	if r.defaultHandlerWrapper.handler != nil {
		panic("Default handler already exists")
	}
	r.defaultHandlerWrapper.handler = f
}

func (r Router) CallHandler(msg IncomingChatMessage) *OutgoingChatMessage {
	if r.handlers[msg.Text] != nil {
		return r.handlers[msg.Text](msg)
	} else {
		for key, handler := range r.regexHandlers {
			match, _ := regexp.MatchString(key, msg.Text)
			if match {
				return handler(msg)
			}
		}
	}
	if r.defaultHandlerWrapper.handler != nil {
		return r.defaultHandlerWrapper.handler(msg)
	}
	return nil
}

func (r Router) RmHandler(cmd string) {
	delete(r.handlers, cmd)
}

func (r Router) RmRegexHandler(cmd string) {
	delete(r.regexHandlers, cmd)
}

func (r Router) ListenUpdates(bot BotInterface) {
	for msg := range bot.GetUpdates() {
		res := r.CallHandler(msg)
		log.Printf("Received answer from handler: %v", res)
		if res != nil {
			log.Printf("Sending answer from handler: %v", res)
			bot.Send(*res)
		}
	}
}

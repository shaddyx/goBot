package gobot

import "regexp"

type RouterHandlerFunc func(msg IncomingChatMessage) *OutgoingChatMessage

type Router struct {
	handlers       map[string]RouterHandlerFunc
	regexHandlers  map[string]RouterHandlerFunc
	defaultHandler RouterHandlerFunc
}

func NewRouter() Router {
	return Router{
		handlers:      make(map[string]RouterHandlerFunc),
		regexHandlers: make(map[string]RouterHandlerFunc),
	}
}

func (r *Router) AddHandler(cmd string, f RouterHandlerFunc) {
	if r.handlers[cmd] != nil {
		panic("Handler already exists:" + cmd)
	}
	r.handlers[cmd] = f
}

func (r *Router) AddRegexHandler(cmd string, f RouterHandlerFunc) {
	if r.regexHandlers[cmd] != nil {
		panic("Handler already exists:" + cmd)
	}
	r.regexHandlers[cmd] = f
}

func (r *Router) SetDefaultHandler(f RouterHandlerFunc) {
	if r.defaultHandler != nil {
		panic("Default handler already exists")
	}
	r.defaultHandler = f
}

func (r *Router) CallHandler(msg IncomingChatMessage) *OutgoingChatMessage {
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
	if r.defaultHandler != nil {
		return r.defaultHandler(msg)
	}
	return nil
}

func (r *Router) RmHandler(cmd string) {
	delete(r.handlers, cmd)
}

func (r *Router) RmRegexHandler(cmd string) {
	delete(r.regexHandlers, cmd)
}

func (r *Router) ListenUpdates(bot BotInterface) {
	for msg := range bot.GetUpdates().Messages {
		res := r.CallHandler(msg)
		if res != nil {
			bot.SendAnswer(*res, &msg)
		}
	}
}

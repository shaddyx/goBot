package gobot

type RouterHandlerFunc func(event string, msg IncomingChatMessage) *OutgoingChatMessage

type Router struct {
	handlers map[string]RouterHandlerFunc
}

func NewRouter() Router {
	return Router{
		handlers: make(map[string]RouterHandlerFunc),
	}
}

func (r *Router) AddHandler(cmd string, f RouterHandlerFunc) {
	if r.handlers[cmd] != nil {
		panic("Handler already exists:" + cmd)
	}
	r.handlers[cmd] = f
}

func (r *Router) CallHandler(cmd string, msg IncomingChatMessage) *OutgoingChatMessage {
	if r.handlers[cmd] != nil {
		return r.handlers[cmd](cmd, msg)
	}
	return nil
}

func (r *Router) RmHandler(cmd string) {
	delete(r.handlers, cmd)
}

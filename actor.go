package actor

import (
	"github.com/jbrodriguez/pubsub"
)

// MailboxHandler -
type MailboxHandler func(msg *pubsub.Message)

// Actor -
type Actor struct {
	bus      *pubsub.PubSub
	mailbox  chan *pubsub.Mailbox
	registry map[string]MailboxHandler
}

// NewActor creates an instance of an actor
func NewActor(bus *pubsub.PubSub) *Actor {
	a := &Actor{
		bus:      bus,
		mailbox:  bus.CreateMailbox(),
		registry: make(map[string]MailboxHandler),
	}

	return a
}

// Register associates a handler with a message topic
func (a *Actor) Register(topic string, handler MailboxHandler) {
	a.bus.Sub(a.mailbox, topic)
	a.registry[topic] = handler
}

// React starts the message handling loop
func (a *Actor) React() {
	for mbox := range a.mailbox {
		a.dispatch(mbox.Topic, mbox.Content)
	}
}

// dispatch is used internally to invoke a handler
func (a *Actor) dispatch(topic string, msg *pubsub.Message) {
	a.registry[topic](msg)
}

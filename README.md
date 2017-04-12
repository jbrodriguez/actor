## Installation

```go get github.com/jbrodriguez/actor```

## Documentation

Visit [docs](http://godoc.org/github.com/jbrodriguez/actor)

## Rationale

Actor implements an actor-like paradigm in go.

Based on a [pubsub](https://github.com/jbrodriguez/pubsub) library, the use case is as follows:

Create a central bus that will transport messages accross the system (locally only):

```bus := pubsub.New(623)```

Create a struct that will behave like an actor

```
type Server struct {
	bus *pubsub.Pubsub
	actor *actor.Actor
}

func NewServer(bus *pubsub.Pubsub) {
	server := &Server{
		bus: bus,
		actor: actor.NewActor(bus),
	}
	return server
}

func (s *Server) Start() {
	// do stuff
	s.actor.Register("topic:message", s.Handler)

	// ...
	
	s.actor.React()
}

func (s *Server) Handler(msg *pubsub.Msg) {
	// do stuff
}
```

A more complete example can be found [here](https://github.com/jbrodriguez/mediagui)
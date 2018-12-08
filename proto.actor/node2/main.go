package main

import (
	"flag"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/remote"
	messages "github.com/ob-vss-ws18/ob-vss-ws18/proto.actor/echomessages"
)

type MyActor struct{}

func (*MyActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *messages.Echo:
		msg.Sender.Tell(&messages.Response{
			SomeValue: "result",
		})
	}
}

var flagBind = flag.String("bind", "localhost:8091", "Bind to address")

func main() {
	flag.Parse()
	remote.Start(*flagBind)
	props := actor.FromProducer(
		func() actor.Actor { return &MyActor{} })
	actor.SpawnNamed(props, "hello")
	time.Sleep(500 * time.Second)
}

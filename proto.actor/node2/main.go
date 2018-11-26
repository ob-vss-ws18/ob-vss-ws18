package main

import (
	console "github.com/AsynkronIT/goconsole"
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

func main() {
	remote.Start("localhost:8091")
	props := actor.FromProducer(
		func() actor.Actor { return &MyActor{} })
	actor.SpawnNamed(props, "hello")
	console.ReadLine()
}

package main

import (
	"fmt"

	console "github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/remote"
	messages "github.com/ob-vss-ws18/ob-vss-ws18/proto.actor/echomessages"
)

type MyActor struct {
	count int
}

func (state *MyActor) Receive(context actor.Context) {
	switch context.Message().(type) {
	case *messages.Response:
		state.count++
		fmt.Println(state.count)
	}
}

func main() {
	remote.Start("localhost:8090")
	props := actor.FromProducer(func() actor.Actor {
		return &MyActor{}
	})
	pid := actor.Spawn(props)
	message := &messages.Echo{Message: "hej", Sender: pid}

	//this is the remote actor we want to communicate with
	remote := actor.NewPID("localhost:8091", "hello")
	for i := 0; i < 10; i++ {
		remote.Tell(message)
	}

	console.ReadLine()
}

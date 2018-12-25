package main

import (
	"flag"
	"sync"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/remote"
	"github.com/ob-vss-ws18/ob-vss-ws18/proto.actor/echomessages"
)

type MyActor struct {
	wg *sync.WaitGroup
}

func (state *MyActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *echomessages.Echo:
		msg.Sender.Tell(&echomessages.Response{
			SomeValue: "result",
		})
	case *actor.Stopped:
		state.wg.Done()
	}
}

var flagBind = flag.String("bind", "localhost:8091", "Bind to address")

func main() {
	flag.Parse()
	remote.Start(*flagBind)
	var wg sync.WaitGroup
	props := actor.FromProducer(
		func() actor.Actor {
			wg.Add(1)
			return &MyActor{&wg}
		})
	actor.SpawnNamed(props, "hello")
	wg.Wait()
}

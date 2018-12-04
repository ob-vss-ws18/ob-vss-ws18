package main

import (
	"fmt"

	console "github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
)

type hello struct{ Who string }
type goodbye struct{ until string }
type helloActor struct{}

func (state *helloActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *hello:
		fmt.Printf("Hello %v\n", msg.Who)
	case *goodbye:
		fmt.Printf("ok cu %v\n", msg.until)
	}
}

func main() {
	props := actor.FromProducer(func() actor.Actor {
		return &helloActor{}
	})
	pid := actor.Spawn(props)
	pid.Tell(&hello{Who: "Roger"})
	pid.Tell(&goodbye{until: "Tomorrow"})
	console.ReadLine()
}

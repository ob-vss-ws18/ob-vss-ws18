package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/AsynkronIT/goconsole"

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

var (
	flagBind  = flag.String("bind", "localhost:8090", "Bind to address")
	flagNode2 = flag.String("node2", "localhost:8091", "node2 host:port")
)

func main() {

	flag.Parse()

	remote.Start(*flagBind)
	props := actor.FromProducer(func() actor.Actor {
		return &MyActor{}
	})
	pid := actor.Spawn(props)
	message := &messages.Echo{Message: "hej", Sender: pid}

	fmt.Println("Sleeping 5 seconds...")
	time.Sleep(5 * time.Second)
	fmt.Println("Awake...")

	//this is the remote actor we want to communicate with
	fmt.Printf("Trying to connect to %s\n", *flagNode2)
	remote := actor.NewPID(*flagNode2, "hello")
	for i := 0; i < 10; i++ {
		remote.Tell(message)
	}

	console.ReadLine()
}

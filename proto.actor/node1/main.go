package main

import (
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/remote"
	"github.com/ob-vss-ws18/ob-vss-ws18/proto.actor/echomessages"
)

type MyActor struct {
	count int
	wg    *sync.WaitGroup
}

func (state *MyActor) Receive(context actor.Context) {
	switch context.Message().(type) {
	case *echomessages.Response:
		state.count++
		fmt.Println(state.count)
	case *actor.Stopped:
		state.wg.Done()
	}
}

var (
	flagBind  = flag.String("bind", "localhost:8090", "Bind to address")
	flagNode2 = flag.String("node2", "localhost:8091", "node2 host:port")
)

func main() {

	flag.Parse()

	remote.Start(*flagBind)
	var wg sync.WaitGroup
	props := actor.FromProducer(func() actor.Actor {
		wg.Add(1)
		return &MyActor{0, &wg}
	})
	pid := actor.Spawn(props)
	message := &echomessages.Echo{Message: "hej", Sender: pid}

	fmt.Println("Sleeping 5 seconds...")
	time.Sleep(5 * time.Second)
	fmt.Println("Awake...")

	//this is the remote actor we want to communicate with
	fmt.Printf("Trying to connect to %s\n", *flagNode2)
	remote := actor.NewPID(*flagNode2, "hello")
	for i := 0; i < 10; i++ {
		remote.Tell(message)
	}

	wg.Wait()
}

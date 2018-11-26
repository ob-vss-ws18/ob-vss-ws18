package main

import (
	"fmt"

	"github.com/vesose/proto.actor/messages"
)

func main() {
	p := messages.Person{
		Id:    1234,
		Name:  "Oliver Braun",
		Email: "ob@cs.hm.edu",
		Phones: []*messages.Person_PhoneNumber{
			{Number: "+49 89 1265-3790", Type: messages.WORK},
		},
	}
	fmt.Println(p.String())

}

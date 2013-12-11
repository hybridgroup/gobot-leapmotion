package main

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot-leapmotion"
)

func main() {
	leapAdaptor := new(gobotLeap.LeapAdaptor)
	leapAdaptor.Name = "leap"
	leapAdaptor.Port = "127.0.0.1:6437"

	leap := gobotLeap.NewLeap(leapAdaptor)
	leap.Name = "leap"

	work := func() {
		go func() {
			for {
				fmt.Println(gobot.On(leap.Events["Message"]))
			}
		}()
	}

	robot := gobot.Robot{
		Connections: []interface{}{ leapAdaptor },
		Devices:     []interface{}{ leap }, 
		Work:        work,
	}

	robot.Start()
}

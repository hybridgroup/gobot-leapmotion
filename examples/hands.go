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
				printHands(gobot.On(leap.Events["Message"]).(gobotLeap.LeapFrame))
			}
		}()
	}

	robot := gobot.Robot{
		Connections: []interface{}{leapAdaptor},
		Devices:     []interface{}{leap},
		Work:        work,
	}

	robot.Start()
}

func printHands(frame gobotLeap.LeapFrame) {
	for key, hand := range frame.Hands {
		fmt.Println("Hand", key, hand)
	}
}

package gobotLeap

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"github.com/hybridgroup/gobot"
)

type LeapAdaptor struct {
	gobot.Adaptor
	Leap *websocket.Conn
}

func (me *LeapAdaptor) Connect() {
	origin := fmt.Sprintf("http://%v", me.Port)
	url := fmt.Sprintf("ws://%v/v3.json", me.Port)
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		panic(err)
	}
	me.Leap = ws
}

func (me *LeapAdaptor) Disconnect() {
}

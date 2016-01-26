package ibot

import (
    "glog"
    "ui"
	// "fmt"
    "time"
)

func (self *Slingo) Start(command_chan chan ui.Command) {
    self.command_chan = command_chan

    for {
        result := self.run(0, 0)
        glog.Infof("running: %v ", result)
        time.Sleep(time.Duration(1) * time.Second)
    }
}

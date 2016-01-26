package main

import (
	"fmt"

	"core"
	"ui"
	"ai/ibot"
)



func main() {

	catas []Catapult

	cata := new_catapult()
	go cata.interprete()
    cata.Bot = ibot.bot

	//ibot.bot.Stop()
	catas.append(cata)

    for cata := range catas {
        cata.Bot.Start(cata.command_chan)
    }

	for {
		// one cycle
		for cata := range catas {
			select {
			case insturction := <- cata.instruction_chan:
				cata.execute(insturction)

			case <- time.After(time.Second):
				// todo optimize
				glog.Infoln("Timeout trigger break.")
			}
		}

        refresh_ground()
	}

}

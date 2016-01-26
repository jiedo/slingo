package main

import (
	"fmt"
    "time"
    "glog"
	"core"
	"ui"
	"ai/ibot"
)

func init () {
    flag.Parse()
    flag.Lookup("logtostderr").Value.Set("true")
    //flag.Lookup("logDir").Value.Set("./")
}

var (
    g_battle_ground = core.Ground {"ground", make([]core.Catapult), make([]core.Ball) }
)

func main() {

	cata := g_battle_ground.New_catapult()
    cata.Bot = ibot.bot

	//ibot.bot.Stop()
    g_battle_ground.Start_all_catapults()

	for {
		// one cycle
        g_battle_ground.Step_all_catapults()

        g_battle_ground.Refresh_ground()

        time.Sleep(time.Duration(1) * time.Second)
	}

}

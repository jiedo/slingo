package main

import (
    "time"
    "flag"
    "glog"
	"core"
	"ai/ibot"
)

func init () {
    flag.Parse()
    flag.Lookup("logtostderr").Value.Set("true")
}

func main() {
    glog.Infof("This is a Catapults AI War platform.")

    ////////////////////////////////////////////////////////////////
    glog.Infof("bots are join in...")

	cata := core.G_battle_ground.New_catapult()
    glog.Infof("bot(%s) in.", ibot.Bot.GetName())
    cata.Bot = ibot.Bot


    glog.Infof("join ok.")

    ////////////////////////////////////////////////////////////////
    glog.Infof("now begin...")
    core.G_battle_ground.Start_all_catapults()
	//ibot.bot.Stop()
    glog.Infof("bots are all running now.")

    ////////////////////////////////////////////////////////////////
    glog.Infof("begin cycle.")
	for {
		// one cycle
        glog.Infof("cycle.")
        core.G_battle_ground.Step_all_catapults()

        core.G_battle_ground.Refresh_ground()

        time.Sleep(time.Duration(1) * time.Second)
	}
    glog.Infof("end.")
}

package core

import (
	"ui"
    "glog"
    "time"
)


const (
    GROUND_CAPACITY int = 100
    CARTRIDGE_CAPACITY int = 100
)


var (
    G_battle_ground = Ground {"ground", make([]*Catapult, 0, GROUND_CAPACITY), make([]*Ball, 0, GROUND_CAPACITY) }
)



func (self *Ground) New_catapult() (cata *Catapult) {
	command_chan := make(chan ui.Command)
	instruction_chan := make(chan Instruction)
    glog.Infof("before new: %v", cata)
    cata = &Catapult{}
    glog.Infof("after new: %v", cata)
	cata.command_chan = command_chan
	cata.instruction_chan = instruction_chan
	go cata.interprete()

    self.catapults = append(self.catapults, cata)
    glog.Infof("catapults: %d", len(self.catapults))
	return cata
}


func (self *Ground) Start_all_catapults() {
    for _, cata := range self.catapults {
        cata.name = cata.Bot.GetName()
        glog.Infof("bot(%s) start.", cata.name)
        go cata.Bot.Start(cata.command_chan)
    }
}


func (self *Ground) Step_all_catapults() {
    for _, cata := range self.catapults {
        select {
        case insturction := <- cata.instruction_chan:
            cata.execute(insturction)
        case <- time.After(time.Second):
            // todo optimize
        }
    }
}

func (self *Ground) Update_ground() {
    glog.Info("Update...")
    // to implement
    for i, cata := range self.catapults {
        glog.Infof("----------------bot%d(%s):----------------", i, cata.name)
        glog.Infof("%v", cata)

    }

    for i, ball := range self.balls {
        glog.Infof("----------------ball%d:----------------", i)
        glog.Infof("%v", ball)

    }

}


func (self *Ground) Refresh_ground() {
    glog.Info("refresh...")
    // to implement
    for i, cata := range self.catapults {
        glog.Infof("----------------bot%d(%s):----------------", i, cata.name)
        glog.Infof("%v", cata)

    }
}

func (self *Ground) scan(position ui.Vector, direction float64, scope float64, distance float64) ([]*Ball, []*Catapult) {
    // to implement
    return self.balls, self.catapults
}

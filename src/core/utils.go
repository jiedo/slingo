package core

import (
	"ui"
    "glog"
)


func (self *Ground) New_catapult() (cata Catapult) {
	command_chan := make(chan ui.Command)
	instruction_chan := make(chan Instruction)

	cata.command_chan = command_chan
	cata.instruction_chan = instruction_chan
	// cata.Init()
	go cata.interprete()

    self.catapults.append(cata)
	return cata
}


func (self *Ground) Start_all_catapults() {
    for cata := range self.catapults {
        go cata.Bot.Start(cata.command_chan)
    }
}


func (self *Ground) Step_all_catapults() {
    for cata := range self.catapults {
        select {
        case insturction := <- cata.instruction_chan:
            cata.execute(insturction)
        case <- time.After(time.Second):
            // todo optimize
        }
    }
}

func (self *Ground) Refresh_ground() {
    glog.Info("refresh...")
}

func (self *Ground) Scan(position Vector, direction float64, scope float64, distance float64) ([]Ball, []Catapult) {
    return self.balls, self.catapults
}

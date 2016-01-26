package main

import (
	"fmt"
	
	"core"
	"ui"
	"ai/ibot"		
)



func new_catapult() chan Command {
	command_chan := make(chan Command)
	instruction_chan := make(chan Instruction)
	
	cata Catapult
	cata.command_chan = command_chan
	cata.instruction_chan = instruction_chan	
	cata.Init()
	
	bot Slingo
	bot.Init()
	bot.Start(command_chan)
	bot.Stop()
	return cata
}


func (self *Catapult) interprete() {
	for {
		command := <- self.command_chan
		swich (command.cmd) {
		case "run":
			self.instruction_chan <- Instruction{"run", "begin"}
			
			self.instruction_chan <- Instruction{"run", "end"}
		case "get":		
		}
	}
}

func (self *Catapult) execute(instruction Instruction) {
	if instruction.stage == "end" {
		insturction.result_chan <- insturction.result
		return
	}
	
	swich (instruction.cmd) {
	case "run":
		insturction.result = "x"

	case "get":
		insturction.result = "y"		
	}

}


func main() {

	catas []Catapult
	
	cata := new_catapult()

	go cata.interprete()
	
	catas.append(cata)

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


	}



	
}

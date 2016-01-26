package core

import (
	"os"
	"glog"
)



type Ball struct {
    material int
	weight int
    size int

    pos_x float64
    pos_y float64
    pos_z float64

    speed_x float64
    speed_y float64
    speed_z	float64

    energy int
}

type Catapult struct {
    life int
    energy int
    weight int

    direction float64
    pos_x float64
    pos_y float64
    pos_z float64

    left_wheel_speed_x float64
    left_wheel_speed_y float64
    left_wheel_speed_z float64

    right_wheel_speed_x float64
    right_wheel_speed_y float64
    right_wheel_speed_z float64

	left_wheel_force float64
    right_wheel_force float64
    aim_direction float64
    aim_elevation float64

    capacity_energy int
    capacity_weight int
    capacity_size int

	balls []Ball
}


type Ground struct {
    name string
    process int
    stdout string
    returncode int
    err error
}



func new_catapult() chan Command {
	command_chan := make(chan Command)
	instruction_chan := make(chan Instruction)

	cata Catapult
	cata.command_chan = command_chan
	cata.instruction_chan = instruction_chan
	cata.Init()
	return cata
}

package core

import "ui"


type Ball struct {
    material ui.Material
	weight int
    size int

    pos ui.Vector
    speed ui.Vector

    is_carried bool
}


type Instruction struct {
    name string
    stage int
    parameter interface{}
    result interface{}
	result_chan chan interface{}
}

type Catapult struct {
    name string
    life int
    energy int
    weight int

    direction float64
    pos ui.Vector
    left_wheel_speed ui.Vector
    right_wheel_speed ui.Vector

	left_wheel_force float64
    right_wheel_force float64
    aim_direction float64
    aim_elevation float64

    capacity_energy int
    capacity_weight int
    capacity_size int

    load_slot *Ball
	balls []*Ball

	command_chan chan ui.Command
	instruction_chan chan Instruction
    Bot ui.AI
}

type Ground struct {
    name string

    catapults []*Catapult
    balls []*Ball
}

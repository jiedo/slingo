package core

import "ui"


type Ball struct {
    is_carried bool

    // update by physics engine
    speed ui.Vector
    pos ui.Vector

    material ui.Material
	weight int
    size int
}


type Instruction struct {
    name string
    stage int
    parameter interface{}
    result interface{}
	result_chan chan interface{}
}

type Catapult struct {
    // update by physics engine
    left_wheel_speed ui.Vector
    right_wheel_speed ui.Vector
    pos ui.Vector
    // change with speed
    direction float64

    // bot can set direct
	left_wheel_force float64
    right_wheel_force float64
    aim_direction float64
    aim_elevation float64

    // communicate with bot
	command_chan chan ui.Command
	instruction_chan chan Instruction

    // control by bot
    load_slot *Ball
	balls []*Ball

    // bot ai
    life int
    energy int
    Bot ui.AI

    // const
    name string
    weight int
    capacity_energy int
    capacity_weight int
    capacity_size int
}

type Ground struct {
    name string

    catapults []*Catapult
    balls []*Ball
}

package core

import "ui"

type Material int

const (
    Stone Material = iota
    Plumbum
)

type Vector struct {
    x float64
    y float64
    z float64
}

type Ball struct {
    material Material
	weight int
    size int

    pos Vector
    speed Vector

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
    life int
    energy int
    weight int

    direction float64
    pos Vector
    left_wheel_speed Vector
    right_wheel_speed Vector

	left_wheel_force float64
    right_wheel_force float64
    aim_direction float64
    aim_elevation float64

    capacity_energy int
    capacity_weight int
    capacity_size int

    load_slot Ball
	balls []Ball

	command_chan chan ui.Command
	instruction_chan chan Instruction
    Bot ui.AI
}

type Ground struct {
    name string

    catapults []Catapult
    balls []Ball
}

package ui

// x_chin := make(chan int)

type Command struct {
    name string
    parameter interface{}
	result_chan chan interface{}
    err error
}

type Instruction struct {
    name string
    stage int
    parameter interface{}
    result interface{}
	result_chan chan interface{}
    err error
}

type Client interface {
	Start(command chan Command)
	Init()
	Stop()
}

type Catapult interface {
	run(left_wheel_speed float64, right_wheel_speed float64)
	aim(direction_angle float64, elevation_angle float64)
	fire()
	reload(material int)
	pick_up_ball(material int, number int)
	throw_away_ball(material int, number int)
	repair(life_point int)
	get_state() (energy int, life int)
	get_position() (pos_x float64, pos_y float64, pos_z float64)
	get_speed() (speed_x float64, speed_y float64, speed_z float64)
	get_direction() (derection_angle float64)
	get_wheel_speed() (left_wheel_speed float64, right_wheel_speed float64)
	get_wheel_force() (left_wheel_force float64, right_wheel_force float64)
	get_carried_balls() []Ball
	is_loaded() (yes_or_no boolean)
	get_aim() (direction_angle float64, elevation_angle float64)
	scan(direction, scope, distance) ([]Ball, []Catapult)
}

type AI interface {
	Client
	Catapult
}

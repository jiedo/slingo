package ui


// x_chin := make(chan int)

type Material int

const (
    Stone Material = iota
    Plumbum
)

type Vector struct {
    X float64
    Y float64
    Z float64
}

type ParaRun struct {
    Left_wheel_force float64
    Right_wheel_force float64
}

type ParaAim struct {
    Direction float64
    Elevation float64
}

type ParaThrowPick struct {
    Material Material
    Number int
}

type ParaScan struct {
    Direction float64
    Scope float64
    Distance float64
}


type Ball struct {
    Material Material
	Weight int
    Size int

    Pos Vector
    Speed Vector

    Is_carried bool
}

type Catapult struct {
    Life int
    Energy int
    Weight int

    Direction float64
    Pos Vector
    Left_wheel_speed Vector
    Right_wheel_speed Vector

	Left_wheel_force float64
    Right_wheel_force float64
    Aim_direction float64
    Aim_elevation float64

    Capacity_energy int
    Capacity_weight int
    Capacity_size int
}


type ResultScan struct {
    Balls []Ball
    Catapults []Catapult
}

type Command struct {
    Name string
	Result_chan chan interface{}
    Parameter interface{}
}

type AI interface {
	Start(command chan Command)
	Init()
	Stop()
}

// type Catapult interface {
// 	run(left_wheel_force float64, right_wheel_force float64)
// 	aim(direction_angle float64, elevation_angle float64)
// 	fire(speed float64)
// 	reload(material Material)
// 	pick_up_ball(material Material, number int)
// 	throw_away_ball(material Material, number int)
// 	repair(life_point int)
// 	get_state() (energy int, life int)
// 	get_position() (pos Vector)
// 	get_speed() (speed Vector)
// 	get_direction() (derection_angle float64)
// 	get_wheel_speed() (left_wheel_speed float64, right_wheel_speed float64)
// 	get_wheel_force() (left_wheel_force float64, right_wheel_force float64)
// 	get_carried_balls() (n_plumbum int, n_stone int)
// 	is_loaded() (yes_or_no bool)
// 	get_aim() (direction_angle float64, elevation_angle float64)
// 	scan(direction, scope, distance) ([]Ball, []Catapult)
// }

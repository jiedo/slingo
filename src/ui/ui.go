package ui


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


type Pill struct {
    Material Material
	Weight int
    Size int

    Pos Vector
    Speed Vector

    Is_carried bool
}

type Sling struct {
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
    Pills []*Pill
    Slings []*Sling
}

type Command struct {
    Name string
	Result_chan chan interface{}
    Parameter interface{}
}

type AI interface {
	Start(command chan Command)
    GetName() string
	Init()
	Stop()
}

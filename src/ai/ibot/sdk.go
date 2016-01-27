package ibot


import (
    "glog"
    "ui"
)


type Slingo struct {
    command_chan chan ui.Command // must keep
    name string                  // should keep
}


func is_channel_ok(command_chan chan ui.Command, name string, function_name string) bool{
    if command_chan == nil {
        glog.Errorf("AI[%s] want %s, but command channel is nil.",
            name, function_name)
        return false
    }
    return true
}


// force > 0: go ahead
// force = 0: stop
// force < 0: go back
// 不同的调整幅度需要不同的耗时和耗能
func (self *Slingo) run(left_wheel_force float64, right_wheel_force float64) bool{
    if !is_channel_ok(self.command_chan, self.name, "run") {
        return false
    }
	result_chan := make(chan interface{})
	self.command_chan <- ui.Command{"run", result_chan, ui.ParaRun{
        Left_wheel_force: left_wheel_force,
        Right_wheel_force: right_wheel_force,
    }}
	result := <- result_chan
    if result_value, ok := result.(bool); ok {
        return result_value
    } else {
        return false
    }
}


// 发射水平方向
// 0 <= direction < 360
// 发射仰角
// 0 <= elevation < 90
// 不同的调整幅度需要不同的耗时和耗能
func (self *Slingo) aim(direction_angle float64, elevation_angle float64) bool{
    if !is_channel_ok(self.command_chan, self.name, "aim") {
        return false
    }
	result_chan := make(chan interface{})
	self.command_chan <- ui.Command{"aim", result_chan, ui.ParaAim{
        Direction: direction_angle,
        Elevation: elevation_angle,
    }}
	result := <- result_chan
    if result_value, ok := result.(bool); ok {
        return result_value
    } else {
        return false
    }
}


// 不同的weight/speed需要不同的耗时和耗能
func (self *Slingo) fire(speed float64) bool{
    if !is_channel_ok(self.command_chan, self.name, "fire") {
        return false
    }
	result_chan := make(chan interface{})
	self.command_chan <- ui.Command{"fire", result_chan, speed}
	result := <- result_chan
    if result_value, ok := result.(bool); ok {
        return result_value
    } else {
        return false
    }
}


// 不同的weight需要不同的耗时和耗能
func (self *Slingo) reload(material ui.Material) bool{
    if !is_channel_ok(self.command_chan, self.name, "reload") {
        return false
    }
	result_chan := make(chan interface{})
	self.command_chan <- ui.Command{"reload", result_chan, material}
	result := <- result_chan
    if result_value, ok := result.(bool); ok {
        return result_value
    } else {
        return false
    }
}

// 不同的weight需要不同的耗时和耗能
func (self *Slingo) pick_up(material ui.Material, number int) int{
    if !is_channel_ok(self.command_chan, self.name, "pick_up") {
        return 0
    }
	result_chan := make(chan interface{})
	self.command_chan <- ui.Command{"pick_up", result_chan, ui.ParaThrowPick{
        Material: material,
        Number: number,
    }}
	result := <- result_chan
    if result_value, ok := result.(int); ok {
        return result_value
    } else {
        return 0
    }
}

// 耗时1
func (self *Slingo) throw_away(material ui.Material, number int) int{
    if !is_channel_ok(self.command_chan, self.name, "throw_away") {
        return 0
    }
	result_chan := make(chan interface{})
	self.command_chan <- ui.Command{"throw_away", result_chan, ui.ParaThrowPick{
        Material: material,
        Number: number,
    }}
	result := <- result_chan
    if result_value, ok := result.(int); ok {
        return result_value
    } else {
        return 0
    }
}


// 不同的life point需要不同的耗时和耗能
func (self *Slingo) repair(life_point int) bool{
    if !is_channel_ok(self.command_chan, self.name, "repair") {
        return false
    }
	result_chan := make(chan interface{})
	self.command_chan <- ui.Command{"repair", result_chan, life_point}
	result := <- result_chan
    if result_value, ok := result.(bool); ok {
        return result_value
    } else {
        return false
    }
}


// 耗时1
func (self *Slingo) get_state() (energy int, life int) {
    if !is_channel_ok(self.command_chan, self.name, "get_state") {
        return
    }
	result_chan := make(chan interface{})
	self.command_chan <- ui.Command{"get_state", result_chan, nil}
	result := <- result_chan
    if result_value, ok := result.([2]int); ok {
        energy = result_value[0]
        life = result_value[1]
        return
    } else {
        return
    }
}

// 耗时1
// 战车中心坐标
func (self *Slingo) get_position() (pos ui.Vector) {
    if !is_channel_ok(self.command_chan, self.name, "get_position") {
        return
    }
	result_chan := make(chan interface{})
	self.command_chan <- ui.Command{"get_position", result_chan, nil}
	result := <- result_chan
    if result_value, ok := result.(ui.Vector); ok {
        return result_value
    } else {
        return
    }
}


// 耗时1
// 向量speed
// 等价于标量direction和标量wheel_speed
func (self *Slingo) get_speed() (speed ui.Vector) {
    if !is_channel_ok(self.command_chan, self.name, "get_speed") {
        return
    }
	result_chan := make(chan interface{})
	self.command_chan <- ui.Command{"get_speed", result_chan, nil}
	result := <- result_chan
    if result_value, ok := result.(ui.Vector); ok {
        return result_value
    } else {
        return
    }
}


// 耗时1
// 标量
func (self *Slingo) get_direction() (direction_angle float64) {
    if !is_channel_ok(self.command_chan, self.name, "get_direction") {
        return
    }
	result_chan := make(chan interface{})
	self.command_chan <- ui.Command{"get_direction", result_chan, nil}
	result := <- result_chan
    if result_value, ok := result.(float64); ok {
        return result_value
    } else {
        return
    }
}

// 耗时1
// 标量
func (self *Slingo) get_wheel_speed() (left_wheel_speed float64, right_wheel_speed float64) {
    if !is_channel_ok(self.command_chan, self.name, "get_wheel_speed") {
        return
    }
	result_chan := make(chan interface{})
	self.command_chan <- ui.Command{"get_wheel_speed", result_chan, nil}
	result := <- result_chan
    if result_value, ok := result.([2]float64); ok {
        left_wheel_speed = result_value[0]
        right_wheel_speed = result_value[1]
    }
    return
}

// 耗时1
// 标量
func (self *Slingo) get_wheel_force() (left_wheel_force float64, right_wheel_force float64) {
    if !is_channel_ok(self.command_chan, self.name, "get_wheel_force") {
        return
    }
	result_chan := make(chan interface{})
	self.command_chan <- ui.Command{"get_wheel_force", result_chan, nil}
	result := <- result_chan
    if result_value, ok := result.([2]float64); ok {
        left_wheel_force = result_value[0]
        right_wheel_force = result_value[1]
    }
    return
}


// 耗时1
func (self *Slingo) get_cartridge() (n_plumbum int, n_stone int) {
    if !is_channel_ok(self.command_chan, self.name, "get_cartridge") {
        return
    }
	result_chan := make(chan interface{})
	self.command_chan <- ui.Command{"get_cartridge", result_chan, nil}
	result := <- result_chan
    if result_value, ok := result.([2]int); ok {
        n_plumbum = result_value[0]
        n_stone = result_value[1]
    }
    return
}

// 耗时1，也许可以废弃
func (self *Slingo) is_loaded() (yes_or_no bool) {
    if !is_channel_ok(self.command_chan, self.name, "is_loaded") {
        return
    }
	result_chan := make(chan interface{})
	self.command_chan <- ui.Command{"is_loaded", result_chan, nil}
	result := <- result_chan
    if result_value, ok := result.(bool); ok {
        return result_value
    } else {
        return false
    }
}

// 耗时1
func (self *Slingo) get_aim() (direction_angle float64, elevation_angle float64) {
    if !is_channel_ok(self.command_chan, self.name, "get_aim") {
        return
    }
	result_chan := make(chan interface{})
	self.command_chan <- ui.Command{"get_aim", result_chan, nil}
	result := <- result_chan
    if result_value, ok := result.([2]float64); ok {
        direction_angle = result_value[0]
        elevation_angle = result_value[1]
    }
    return
 }

// 耗时随scan面积不同而不同
// direction是方位角度，scope是左右范围角度和，distance是扫描距离深度
func (self *Slingo) scan(direction float64, scope float64, distance float64) (pills []*ui.Pill, slings[]*ui.Sling) {
    if !is_channel_ok(self.command_chan, self.name, "scan") {
        return
    }
	result_chan := make(chan interface{})
	self.command_chan <- ui.Command{"scan", result_chan, ui.ParaScan{
        Direction: direction,
        Scope: scope,
        Distance: distance,
    }}
	result := <- result_chan
    if result_value, ok := result.(ui.ResultScan); ok {
        pills = result_value.Pills
        slings = result_value.Slings
    }
    return
}


func (self *Slingo) Stop() {
    self.command_chan = nil
}


func (self *Slingo) GetName() string{
    return self.name
}

var Bot *Slingo

func init() {
	Bot = &Slingo{}
	Bot.Init()
}

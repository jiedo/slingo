package core

// speed > 0: go ahead
// speed = 0: stop
// speed < 0: go back
// 不同的调整幅度需要不同的耗时和耗能        
func (self *Catapult) run(left_wheel_speed float64, right_wheel_speed float64) {
	
}


// 发射水平方向
// 0 <= direction < 360
// 发射仰角
// 0 <= elevation < 90    
// 不同的调整幅度需要不同的耗时和耗能    
func (self *Catapult) aim(direction_angle float64, elevation_angle float64) {
}


// 不同的weight需要不同的耗时和耗能
func (self *Catapult) fire() {
}


// 不同的weight需要不同的耗时和耗能
func (self *Catapult) reload(material int) {
}

// 不同的weight需要不同的耗时和耗能
func (self *Catapult) pick_up_ball(material int, number int) {
}

// 耗时1
func (self *Catapult) throw_away_ball(material int, number int) {
}


// 不同的life point需要不同的耗时和耗能
func (self *Catapult) repair(life_point int) {
}



#// 获取信息

// 耗时1
func (self *Catapult) get_state() (energy int, life int) {
}

// 耗时1
// 战车中心坐标
func (self *Catapult) get_position() (pos_x float64, pos_y float64, pos_z float64) {
}


// 耗时1    
// 向量speed
// 等价于标量direction和标量wheel_speed
func (self *Catapult) get_speed() (speed_x float64, speed_y float64, speed_z float64) {
}   


// 耗时1   
// 标量
func (self *Catapult) get_direction() derection_angle float64 {
}

// 耗时1    
// 标量
func (self *Catapult) get_wheel_speed() left_wheel_speed float64, right_wheel_speed float64 {
}   

// 耗时1   
// 标量
func (self *Catapult) get_wheel_force() left_wheel_force float64, right_wheel_force float64 {
 }   


// 耗时1    
func (self *Catapult) get_carried_balls() []Ball {
}

// 耗时1，也许可以废弃
func (self *Catapult) is_loaded() yes_or_no boolean {
}

// 耗时1    
func (self *Catapult) get_aim() (direction_angle float64, elevation_angle float64) {
 }

// 耗时随scan面积不同而不同
// direction是方位角度，scope是左右范围角度和，distance是扫描距离深度
func (self *Catapult) scan(direction, scope, distance) ([]Ball, []Catapult) {
}




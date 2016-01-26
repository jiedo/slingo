package core

import (
    "ui"
)

// 调整命令到指令, 每个bot启动一个, 并持续运行
func (self *Catapult) interprete() {
	for {
		cmd := <- self.command_chan
		switch (cmd.Name) {
		case "run":
			self.instruction_chan <- Instruction {
                name: cmd.Name,
                result_chan: cmd.Result_chan,
                stage: 2,       // start
                parameter: ui.ParaRun{
                    Left_wheel_force: cmd.Parameter.Left_wheel_force,
                    Right_wheel_force: cmd.Parameter.Right_wheel_force,
                },
            }
			self.instruction_chan <- Instruction {
                name: cmd.Name,
                result_chan: cmd.Result_chan,
                stage: 1,       // doing
                parameter: ui.ParaRun{
                    Left_wheel_force: cmd.Parameter.Left_wheel_force,
                    Right_wheel_force: cmd.Parameter.Right_wheel_force,
                },
            }
			self.instruction_chan <- Instruction {
                name: cmd.Name,
                result_chan: cmd.Result_chan,
                stage: 0,       // end
                parameter: ui.ParaRun{
                    Left_wheel_force: cmd.Parameter.Left_wheel_force,
                    Right_wheel_force: cmd.Parameter.Right_wheel_force,
                },
            }
		case "aim":
			self.instruction_chan <- Instruction {
                name: cmd.Name,
                result_chan: cmd.Result_chan,
                stage: 2,       // start
                parameter: ui.ParaAim{
                    Direction: cmd.Parameter.Direction,
                    Elevation: cmd.Parameter.Elevation,
                },
            }
			self.instruction_chan <- Instruction {
                name: cmd.Name,
                result_chan: cmd.Result_chan,
                stage: 1,       // doing
                parameter: ui.ParaAim{
                    Direction: cmd.Parameter.Direction,
                    Elevation: cmd.Parameter.Elevation,
                },
            }
			self.instruction_chan <- Instruction {
                name: cmd.Name,
                result_chan: cmd.Result_chan,
                stage: 0,       // end
                parameter: ui.ParaAim{
                    Direction: cmd.Parameter.Direction,
                    Elevation: cmd.Parameter.Elevation,
                },
            }
		case "fire":
			self.instruction_chan <- Instruction {
                name: cmd.Name,
                result_chan: cmd.Result_chan,
                stage: 0,       // start
                parameter: cmd.Parameter,
            }
		case "reload":
			self.instruction_chan <- Instruction {
                name: cmd.Name,
                result_chan: cmd.Result_chan,
                stage: 0,       // start
                parameter: cmd.Parameter,
            }
		case "pick_up_ball":
			self.instruction_chan <- Instruction {
                name: cmd.Name,
                result_chan: cmd.Result_chan,
                stage: 0,       // start
                parameter: ui.ParaThrowPick{
                    Material: cmd.Parameter.Material,
                    Number: cmd.Parameter.Number,
                },
            }
		case "throw_away_ball":
			self.instruction_chan <- Instruction {
                name: cmd.Name,
                result_chan: cmd.Result_chan,
                stage: 0,       // start
                parameter: ui.ParaThrowPick{
                    Material: cmd.Parameter.Material,
                    Number: cmd.Parameter.Number,
                },
            }
		case "repair":
			self.instruction_chan <- Instruction {
                name: cmd.Name,
                result_chan: cmd.Result_chan,
                stage: 0,       // start
                parameter: cmd.Parameter,
            }
		case "get_state":
			self.instruction_chan <- Instruction {
                name: cmd.Name,
                result_chan: cmd.Result_chan,
                stage: 0,       // start
            }
		case "get_position":
			self.instruction_chan <- Instruction {
                name: cmd.Name,
                result_chan: cmd.Result_chan,
                stage: 0,       // start
            }
		case "get_speed":
			self.instruction_chan <- Instruction {
                name: cmd.Name,
                result_chan: cmd.Result_chan,
                stage: 0,       // start
            }
		case "get_direction":
			self.instruction_chan <- Instruction {
                name: cmd.Name,
                result_chan: cmd.Result_chan,
                stage: 0,       // start
            }
		case "get_wheel_speed":
			self.instruction_chan <- Instruction {
                name: cmd.Name,
                result_chan: cmd.Result_chan,
                stage: 0,       // start
            }
		case "get_wheel_force":
			self.instruction_chan <- Instruction {
                name: cmd.Name,
                result_chan: cmd.Result_chan,
                stage: 0,       // start
            }
		case "get_carried_balls":
			self.instruction_chan <- Instruction {
                name: cmd.Name,
                result_chan: cmd.Result_chan,
                stage: 0,       // start
            }
		case "is_loaded":
			self.instruction_chan <- Instruction {
                name: cmd.Name,
                result_chan: cmd.Result_chan,
                stage: 0,       // start
            }
		case "get_aim":
			self.instruction_chan <- Instruction {
                name: cmd.Name,
                result_chan: cmd.Result_chan,
                stage: 0,       // start
            }
		case "scan":
			self.instruction_chan <- Instruction {
                name: cmd.Name,
                result_chan: cmd.Result_chan,
                stage: 0,       // start
                parameter: ui.ParaScan{
                    Direction: cmd.Parameter.Direction,
                    Scope: cmd.Parameter.Scope,
                    Distance: cmd.Parameter.Distance,
                },
            }
		}                       // end of switch
	}
}


// 执行单条指令, 每个bot每周期运行一次
func (self *Catapult) execute(instruction Instruction) {
    insturction.result = nil
	switch (instruction.name) {
	case "run":
        insturction.result = self.run(
            insturction.parameter.left_wheel_force,
            insturction.parameter.right_wheel_force)
	case "aim":
        insturction.result = self.aim(
            insturction.parameter.derection,
            insturction.parameter.elevation)
    case "fire":
        insturction.result = self.fire(insturction.parameter)
    case "reload":
        insturction.result = self.reload(insturction.parameter)
    case "pick_up_ball":
        insturction.result = self.pick_up_ball(
            insturction.parameter.material,
            insturction.parameter.number)
    case "throw_away_ball":
        insturction.result = self.throw_away_ball(
            insturction.parameter.material,
            insturction.parameter.number)
    case "repair":
        insturction.result = self.repair(insturction.parameter)
    case "get_state":
        energy, life := self.get_state()
        insturction.result = [2]int{energy, life}
    case "get_position":
        pos := self.get_position()
        insturction.result = pos
    case "get_speed":
        speed := self.get_speed()
        insturction.result = speed
    case "get_direction":
        derection_angle := self.get_direction()
        insturction.result = derection_angle
    case "get_wheel_speed":
        left_wheel_speed, right_wheel_speed := self.get_wheel_speed()
        insturction.result = [2]float64{left_wheel_speed, right_wheel_speed}
    case "get_wheel_force":
        left_wheel_force, right_wheel_force := self.get_wheel_force()
        insturction.result = [2]float64{left_wheel_force, right_wheel_force}
    case "get_carried_balls":
        n_plumbum, n_stone := self.get_carried_balls()
        insturction.result = [2]int{n_plumbum, n_stone}
    case "is_loaded":
        yes_or_no := self.is_loaded()
        insturction.result = yes_or_no
    case "get_aim":
        derection_angle, elevation_angle := self.get_aim()
        insturction.result = [2]float64{derection_angle, elevation_angle}
    case "scan":
        balls, catapults := self.scan(
            insturction.parameter.direction,
            insturction.parameter.scope,
            insturction.parameter.distance)

        ui_balls = make([]ui.Ball)
        for ball := range balls {
            append(ui_balls, ui.Ball{
                Material:   ball.material,
                Weight:     ball.weight,
                Size:       ball.size,
                Pos:        ball.pos,
                Speed:      ball.speed,
                Is_carried: ball.is_carried,
            })
        }

        ui_catapults = make([]ui.Catapult)
        for catapult := range catapults {
            append(ui_catapults, ui.Catapult{
                Life:               catapult.life,
                Energy:             catapult.energy,
                Weight:             catapult.weight,
                Direction:          catapult.direction,
                Pos:                catapult.pos,
                Left_wheel_speed:   catapult.left_wheel_speed,
                Right_wheel_speed:  catapult.right_wheel_speed,
                Left_wheel_force :  catapult.left_wheel_force,
                Right_wheel_force : catapult.right_wheel_force,
                Aim_direction :     catapult.aim_direction,
                Aim_elevation :     catapult.aim_elevation,
                Capacity_energy:    catapult.capacity_energy,
                Capacity_weight :   catapult.capacity_weight,
                Capacity_size :     catapult.capacity_size,
            })
        }
        insturction.result = ui.ResultScan{
            balls: ui_balls,
            catapults: ui_catapults,
        }
	}
	if instruction.stage == 0 {  // finish
		insturction.result_chan <- insturction.result
		return
	}
}

// speed > 0: go ahead
// speed = 0: stop
// speed < 0: go back
// 不同的调整幅度需要不同的耗时和耗能
func (self *Catapult) run_speed(left_wheel_speed float64, right_wheel_speed float64) {

}


// force > 0: go ahead
// force = 0: free
// force < 0: go back
// 不同的调整幅度需要不同的耗时和耗能
// 返回值可以是坐标和速度, 便于精细控制
func (self *Catapult) run(left_wheel_force float64, right_wheel_force float64) {
    self.left_wheel_force = left_wheel_force
    self.right_wheel_force = right_wheel_force
    return true
}


// 发射水平方向
// 0 <= direction < 360
// 发射仰角
// 0 <= elevation < 90
// 不同的调整幅度需要不同的耗时和耗能
func (self *Catapult) aim(direction_angle float64, elevation_angle float64) {
    if direction_angle >= 360 || direction_angle < 0 {
        return false
    }
    if elevation_angle >= 90 || elevation_angle < 0 {
        return false
    }
    self.aim_direction = direction_angle
    self.aim_elevation = elevation_angle
    return true
}


// 不同的weight/speed需要不同的耗时和耗能
func (self *Catapult) fire(speed float64) {
    if speed <= 0 {
        return false
    }
    if !self.is_loaded() {
        return false
    }
    ball = self.load_slot
    ball.pos.x = self.pos.x
    ball.pos.y = self.pos.y
    ball.pos.z = self.pos.z

    cosd = math.cos(self.aim_direction)
    sind = math.sin(self.aim_direction)
    cose = math.cos(self.aim_elevation)
    sine = math.sin(self.aim_elevation)

    ball.speed.x = speed * cose * cosd
    ball.speed.y = speed * cose * sind
    ball.speed.z = speed * sine

    ball.is_carried = false
    self.load_slot = nil
    return true
}


// 不同的weight需要不同的耗时和耗能
func (self *Catapult) reload(material Material) {
    if self.is_loaded() {
        return true
    }
    j := 0
    for i:=0; i!=len(self.balls); i++ {
        ball := self.balls[i]
        if material == ball.material && ball.is_carried {
            self.load_slot = ball
            self.balls[i] = self.balls[j]
            j++
            break
        }
    }
    if j > 0 {
        self.balls = self.balls[j:]
        return true
    }
    return false
}

// 不同的weight需要不同的耗时和耗能
// return number of picked up
func (self *Catapult) pick_up_ball(material Material, number int) {
    if number <= 0 {
        return 0
    }
    // todo: 读取ground.balls
    j := 0
    for i:=0; i!=len(ground.balls) && j!=number; i++ {
        ball := ground.balls[i]
        if material == ball.material && !ball.is_carried {
            ball.is_carried = true
            append(self.balls, ball)
            j++
        }
    }
    return j
}

// 耗时1
// number 可能多于已有的同类material ball数目
// return: number of throwed
func (self *Catapult) throw_away_ball(material Material, number int) {
    if number <= 0 {
        return 0
    }
    // todo: 将balls归还地面
    j := 0
    for i:=0; i!=len(self.balls) && j!=number; i++ {
        ball := self.balls[i]
        if material == ball.material && ball.is_carried {
            ball.is_carried = false
            self.balls[i] = self.balls[j]
            j++
        }
    }
    if j > 0 {
        self.balls = self.balls[j:]
    }
    return j
}

// 不同的life point需要不同的耗时和耗能
func (self *Catapult) repair(life_point int) {
    self.life += life_point
    return true
}

// 耗时1
func (self *Catapult) get_state() (energy int, life int) {
    return self.energy, self.life
}

// 耗时1
// 战车中心坐标
func (self *Catapult) get_position() (pos Vector) {
    return self.pos
}

// 耗时1
// 向量speed
// 等价于标量direction和标量wheel_speed
func (self *Catapult) get_speed() (speed Vector) {
    return self.speed
}

// 耗时1
// 标量
func (self *Catapult) get_direction() (derection_angle float64) {
    return self.direction
}

// 耗时1
// 标量
func (self *Catapult) get_wheel_speed() (left_wheel_speed float64, right_wheel_speed float64) {
    left_wheel_speed = math.sqrt(self.left_wheel_speed.x * self.left_wheel_speed.x +
        self.left_wheel_speed.y * self.left_wheel_speed.y +
        self.left_wheel_speed.z * self.left_wheel_speed.z)
    right_wheel_speed = math.sqrt(self.right_wheel_speed.x * self.right_wheel_speed.x +
        self.right_wheel_speed.y * self.right_wheel_speed.y +
        self.right_wheel_speed.z * self.right_wheel_speed.z)
    return
}

// 耗时1
// 标量
func (self *Catapult) get_wheel_force() (left_wheel_force float64, right_wheel_force float64) {
    return self.left_wheel_force, self.right_wheel_force
}

// 耗时1
func (self *Catapult) get_carried_balls() (n_plumbum int, n_stone int) {
    n_plumbum = 0
    n_stone = 0
    for ball := range self.balls {
        if ball.material == Plumbum {
            n_plumbum += 1
        } else if ball.material == Stone {
            n_stone += 1
        }
    }
    return
}

// 耗时1，也许可以废弃
func (self *Catapult) is_loaded() (yes_or_no bool) {
    return (self.load_slot != nil)
}

// 耗时1
func (self *Catapult) get_aim() (direction_angle float64, elevation_angle float64) {
    return self.aim_direction, self.aim_elevation
}

// 耗时随scan面积不同而不同
// direction是方位角度，scope是左右范围角度和，distance是扫描距离深度
func (self *Catapult) scan(direction float64, scope float64, distance float64) ([]Ball, []Catapult) {
    return g_battle_ground.scan(self.pos, direction, scope, distance)
}

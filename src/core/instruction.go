package core

import (
    "ui"
    "math"
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
                parameter: cmd.Parameter,
            }
			self.instruction_chan <- Instruction {
                name: cmd.Name,
                result_chan: cmd.Result_chan,
                stage: 1,       // doing
                parameter: cmd.Parameter,
            }
			self.instruction_chan <- Instruction {
                name: cmd.Name,
                result_chan: cmd.Result_chan,
                stage: 0,       // end
                parameter: cmd.Parameter,
            }
		case "aim":
			self.instruction_chan <- Instruction {
                name: cmd.Name,
                result_chan: cmd.Result_chan,
                stage: 2,       // start
                parameter: cmd.Parameter,
            }
			self.instruction_chan <- Instruction {
                name: cmd.Name,
                result_chan: cmd.Result_chan,
                stage: 1,       // doing
                parameter: cmd.Parameter,
            }
			self.instruction_chan <- Instruction {
                name: cmd.Name,
                result_chan: cmd.Result_chan,
                stage: 0,       // end
                parameter: cmd.Parameter,
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
		case "pick_up":
			self.instruction_chan <- Instruction {
                name: cmd.Name,
                result_chan: cmd.Result_chan,
                stage: 0,       // start
                parameter: cmd.Parameter,
            }
		case "throw_away":
			self.instruction_chan <- Instruction {
                name: cmd.Name,
                result_chan: cmd.Result_chan,
                stage: 0,       // start
                parameter: cmd.Parameter,
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
		case "get_cartridge":
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
                parameter: cmd.Parameter,
            }
		}                       // end of switch
	}
}


// 执行单条指令, 每个bot每周期运行一次
func (self *Catapult) execute(instruction Instruction) {
    instruction.result = nil
	switch (instruction.name) {
	case "run":
        if para, ok := instruction.parameter.(ui.ParaRun); ok {
            instruction.result = self.run(
                para.Left_wheel_force,
                para.Right_wheel_force)
        }
	case "aim":
        if para, ok := instruction.parameter.(ui.ParaAim); ok {
            instruction.result = self.aim(
                para.Direction,
                para.Elevation)
        }
    case "fire":
        if para, ok := instruction.parameter.(float64); ok {
            instruction.result = self.fire(para)
        }
    case "reload":
        if para, ok := instruction.parameter.(ui.Material); ok {
            instruction.result = self.reload(para)
        }
    case "pick_up":
        if para, ok := instruction.parameter.(ui.ParaThrowPick); ok {
            instruction.result = self.pick_up(
                para.Material,
                para.Number)
        }
    case "throw_away":
        if para, ok := instruction.parameter.(ui.ParaThrowPick); ok {
            instruction.result = self.throw_away(
                para.Material,
                para.Number)
        }
    case "repair":
        if para, ok := instruction.parameter.(int); ok {
            instruction.result = self.repair(para)
        }
    case "get_state":
        energy, life := self.get_state()
        instruction.result = [2]int{energy, life}
    case "get_position":
        pos := self.get_position()
        instruction.result = pos
    case "get_speed":
        speed := self.get_speed()
        instruction.result = speed
    case "get_direction":
        direction_angle := self.get_direction()
        instruction.result = direction_angle
    case "get_wheel_speed":
        left_wheel_speed, right_wheel_speed := self.get_wheel_speed()
        instruction.result = [2]float64{left_wheel_speed, right_wheel_speed}
    case "get_wheel_force":
        left_wheel_force, right_wheel_force := self.get_wheel_force()
        instruction.result = [2]float64{left_wheel_force, right_wheel_force}
    case "get_cartridge":
        n_plumbum, n_stone := self.get_cartridge()
        instruction.result = [2]int{n_plumbum, n_stone}
    case "is_loaded":
        yes_or_no := self.is_loaded()
        instruction.result = yes_or_no
    case "get_aim":
        direction_angle, elevation_angle := self.get_aim()
        instruction.result = [2]float64{direction_angle, elevation_angle}
    case "scan":
        if para, ok := instruction.parameter.(ui.ParaScan); ok {
            balls, catapults := self.scan(
                para.Direction,
                para.Scope,
                para.Distance)

            ui_pills := make([]*ui.Pill, CARTRIDGE_CAPACITY)
            for _, ball := range balls {
                ui_pills = append(ui_pills, &ui.Pill{
                    Material:   ball.material,
                    Weight:     ball.weight,
                    Size:       ball.size,
                    Pos:        ball.pos,
                    Speed:      ball.speed,
                    Is_carried: ball.is_carried,
                })
            }

            ui_slings := make([]*ui.Sling, GROUND_CAPACITY)
            for _, catapult := range catapults {
                ui_slings = append(ui_slings, &ui.Sling{
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
            instruction.result = ui.ResultScan{
                Pills: ui_pills,
                Slings: ui_slings,
            }
        }
	}
	if instruction.stage == 0 {  // finish
		instruction.result_chan <- instruction.result
		return
	}
}

// speed > 0: go ahead
// speed = 0: stop
// speed < 0: go back
// 不同的调整幅度需要不同的耗时和耗能
func (self *Catapult) run_speed(left_wheel_speed float64, right_wheel_speed float64) bool{
    return false
}


// force > 0: go ahead
// force = 0: free
// force < 0: go back
// 不同的调整幅度需要不同的耗时和耗能
// 返回值可以是坐标和速度, 便于精细控制
func (self *Catapult) run(left_wheel_force float64, right_wheel_force float64) bool{
    self.left_wheel_force = left_wheel_force
    self.right_wheel_force = right_wheel_force
    return true
}


// 发射水平方向
// 0 <= direction < 360
// 发射仰角
// 0 <= elevation < 90
// 不同的调整幅度需要不同的耗时和耗能
func (self *Catapult) aim(direction_angle float64, elevation_angle float64) bool{
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
func (self *Catapult) fire(speed float64) bool{
    if speed <= 0 {
        return false
    }
    if !self.is_loaded() {
        return false
    }
    ball := self.load_slot
    ball.pos.X = self.pos.X
    ball.pos.Y = self.pos.Y
    ball.pos.Z = self.pos.Z

    cosd := math.Cos(self.aim_direction)
    sind := math.Sin(self.aim_direction)
    cose := math.Cos(self.aim_elevation)
    sine := math.Sin(self.aim_elevation)

    ball.speed.X = speed * cose * cosd
    ball.speed.Y = speed * cose * sind
    ball.speed.Z = speed * sine

    ball.is_carried = false
    self.load_slot = nil
    return true
}


// 不同的weight需要不同的耗时和耗能
func (self *Catapult) reload(material ui.Material) bool{
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
func (self *Catapult) pick_up(material ui.Material, number int) int{
    if number <= 0 {
        return 0
    }
    // todo: 读取ground.balls
    j := 0
    for i:=0; i!=len(G_battle_ground.balls) && j!=number; i++ {
        ball := G_battle_ground.balls[i]
        if material == ball.material && !ball.is_carried {
            ball.is_carried = true
            self.balls = append(self.balls, ball)
            j++
        }
    }
    return j
}

// 耗时1
// number 可能多于已有的同类material ball数目
// return: number of throwed
func (self *Catapult) throw_away(material ui.Material, number int) int{
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
func (self *Catapult) repair(life_point int) bool{
    self.life += life_point
    return true
}

// 耗时1
func (self *Catapult) get_state() (energy int, life int) {
    return self.energy, self.life
}

// 耗时1
// 战车中心坐标
func (self *Catapult) get_position() (pos ui.Vector) {
    return self.pos
}

// 耗时1
// 向量speed
// 等价于标量direction和标量wheel_speed
func (self *Catapult) get_speed() (speed ui.Vector) {
    return ui.Vector{
        X: (self.left_wheel_speed.X + self.right_wheel_speed.X)/2,
        Y: (self.left_wheel_speed.Y + self.right_wheel_speed.Y)/2,
        Z: (self.left_wheel_speed.Z + self.right_wheel_speed.Z)/2,
    }
}

// 耗时1
// 标量
func (self *Catapult) get_direction() (direction_angle float64) {
    return self.direction
}

// 耗时1
// 标量
func (self *Catapult) get_wheel_speed() (left_wheel_speed float64, right_wheel_speed float64) {
    left_wheel_speed = math.Sqrt(self.left_wheel_speed.X * self.left_wheel_speed.X +
        self.left_wheel_speed.Y * self.left_wheel_speed.Y +
        self.left_wheel_speed.Z * self.left_wheel_speed.Z)
    right_wheel_speed = math.Sqrt(self.right_wheel_speed.X * self.right_wheel_speed.X +
        self.right_wheel_speed.Y * self.right_wheel_speed.Y +
        self.right_wheel_speed.Z * self.right_wheel_speed.Z)
    return
}

// 耗时1
// 标量
func (self *Catapult) get_wheel_force() (left_wheel_force float64, right_wheel_force float64) {
    return self.left_wheel_force, self.right_wheel_force
}

// 耗时1
func (self *Catapult) get_cartridge() (n_plumbum int, n_stone int) {
    n_plumbum = 0
    n_stone = 0
    for _, ball := range self.balls {
        if ball.material == ui.Plumbum {
            n_plumbum += 1
        } else if ball.material == ui.Stone {
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
func (self *Catapult) scan(direction float64, scope float64, distance float64) ([]*Ball, []*Catapult) {
    return G_battle_ground.scan(self.pos, direction, scope, distance)
}

package core


type paraRun struct {
    left_wheel_speed float64
    right_wheel_speed float64
}

type paraAim struct {
    direction float64
    elevation float64
}

type paraThrowPick struct {
    material int
    number int
}

type paraScan struct {
    direction float64
    scope float64
    distance float64
}

type resultScan struct {
    balls []Ball
    catapults []Catapult
}


// 调整命令到指令, 每个bot启动一个, 并持续运行
func (self *Catapult) interprete() {
	for {
		cmd := <- self.command_chan
		swich (cmd.name) {
		case "run":
			self.instruction_chan <- Instruction {
                name: cmd.name,
                stage: 0,       // start
                parameter: paraRun{
                    left_wheel_speed: cmd.parameter.left_wheel_speed,
                    right_wheel_speed: cmd.parameter.right_wheel_speed,
                },
            }
			self.instruction_chan <- Instruction {
                name: cmd.name,
                stage: 1,       // doing
                parameter: paraRun{
                    left_wheel_speed: cmd.parameter.left_wheel_speed,
                    right_wheel_speed: cmd.parameter.right_wheel_speed,
                },
            }
			self.instruction_chan <- Instruction {
                name: cmd.name,
                stage: -1,       // end
                parameter: paraRun{
                    left_wheel_speed: cmd.parameter.left_wheel_speed,
                    right_wheel_speed: cmd.parameter.right_wheel_speed,
                },
            }
		case "aim":

			self.instruction_chan <- Instruction {
                name: cmd.name,
                stage: 0,       // start
                parameter: paraAim{
                    direction: cmd.parameter.direction,
                    elevation: cmd.parameter.elevation,
                },
            }
			self.instruction_chan <- Instruction {
                name: cmd.name,
                stage: 1,       // doing
                parameter: paraAim{
                    direction: cmd.parameter.direction,
                    elevation: cmd.parameter.elevation,
                },
            }
			self.instruction_chan <- Instruction {
                name: cmd.name,
                stage: -1,       // end
                parameter: paraAim{
                    direction: cmd.parameter.direction,
                    elevation: cmd.parameter.elevation,
                },
            }
		case "fire":
			self.instruction_chan <- Instruction {
                name: cmd.name,
                stage: 0,       // start
            }

		case "reload":
			self.instruction_chan <- Instruction {
                name: cmd.name,
                stage: 0,       // start
                parameter: cmd.parameter
            }



		}
	}
}


// 执行单条指令, 每个bot每周期运行一次
func (self *Catapult) execute(instruction Instruction) {
	if instruction.stage < 0 {  // end
		insturction.result_chan <- insturction.result
		return
	}
    insturction.result = nil
	swich (instruction.name) {
	case "run":
        self.run(insturction.parameter.left_wheel_speed, insturction.parameter.right_wheel_speed)
	case "aim":
        self.aim(insturction.parameter.derection, insturction.parameter.elevation)
    case "fire":
        self.fire()
    case "reload":
        self.reload(insturction.parameter)
    case "pick_up_ball":
        self.pick_up_ball(insturction.parameter.material, insturction.parameter.number)
    case "throw_away_ball":
        self.throw_away_ball(insturction.parameter.material, insturction.parameter.number)
    case "repair":
        self.repair(insturction.parameter)
    case "get_state":
        energy, life := self.get_state()
        insturction.result = [2]int { energy, life, }
    case "get_position":
        pos_x, pos_y, pos_z := self.get_position()
        insturction.result = [3]float64 { pos_x, pos_y, pos_z, }
    case "get_speed":
        speed_x, speed_y, speed_z := self.get_speed()
        insturction.result = [3]float64 { speed_x, speed_y, speed_z, }
    case "get_direction":
        derection_angle := self.get_direction()
        insturction.result = derection_angle
    case "get_wheel_speed":
        left_wheel_speed, right_wheel_speed := self.get_wheel_speed()
        insturction.result = [2]float64 { left_wheel_speed, right_wheel_speed }
    case "get_wheel_force":
        left_wheel_force, right_wheel_force := self.get_wheel_force()
        insturction.result = [2]float64 { left_wheel_force, right_wheel_force }
    case "get_carried_balls":
        // []Ball
        balls := self.get_carried_balls()
        insturction.result = balls
    case "is_loaded":
        yes_or_no := self.is_loaded()
        insturction.result = yes_or_no
    case "get_aim":
        derection_angle, elevation_angle := self.get_aim()
        insturction.result = [2]float64 { derection_angle, elevation_angle }
    case "scan":
        balls, catapults := self.scan(
            insturction.parameter.direction,
            insturction.parameter.scope,
            insturction.parameter.distance)
        insturction.result = resultScan {
            balls: balls,
            catapults: catapults,
        }
	}
}

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
    self.aim_direction = direction_angle
    self.aim_elevation = elevation_angle
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
func (self *Catapult) get_wheel_speed() (left_wheel_speed float64, right_wheel_speed float64) {
}

// 耗时1
// 标量
func (self *Catapult) get_wheel_force() (left_wheel_force float64, right_wheel_force float64) {
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
func (self *Catapult) scan(direction float64, scope float64, distance float64) ([]Ball, []Catapult) {
}

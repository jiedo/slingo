# 基本数据

catapult:

    server:

    life
    energy
    weight
    capacity_energy
    capacity_weight
    capacity_size

    direction
    pos_x
    pos_y
    pos_z
    
    left_wheel_speed_x
    left_wheel_speed_y
    left_wheel_speed_z
                     
    right_wheel_speed_x
    right_wheel_speed_y
    right_wheel_speed_z

    balls[]
    
    
    user:
    
    left_wheel_force
    right_wheel_force
    aim_direction
    aim_elevation


ball:

    type
    weight
    size
    
    pos_x
    pos_y
    pos_z
    
    speed_x
    speed_y        
    speed_z
    

# user interface

## 控制

运动, 调用后保持，直到重新调用。
    
    # speed > 0: go ahead
    # speed = 0: stop
    # speed < 0: go back
    # 不同的调整幅度需要不同的耗时和耗能        
    run(left_wheel_speed, right_wheel_speed)

调整攻击角度，设置后保持，不随车方向变动。

    # 发射水平方向
    # 0 <= direction < 360
    # 发射仰角
    # 0 <= elevation < 90    
    # 不同的调整幅度需要不同的耗时和耗能    
    aim(direction_angle, elevation_angle)
    
发射

    # 不同的weight需要不同的耗时和耗能
    fire()

弹药

    # 不同的weight需要不同的耗时和耗能
    reload(type)

    # 不同的weight需要不同的耗时和耗能
    pick_up_ball(type, number)
    
    # 耗时1
    throw_away_ball(type, number)

修复

    # 不同的life point需要不同的耗时和耗能
    repair(life_point)



## 获取信息

获取自身的生命值和能量。当生命值为0时即死亡，能量过低将不能够执行某动作。
出生时拥有一定生命值和能量，生命值和能量值都有上限。能量会自动恢复，生命值可以通过修复(repair)来恢复。
修复需要耗费能量。

    # 耗时1
    energy, life = get_state()
    
    # 耗时1
    # 战车中心坐标
    position_x, position_y, position_z = get_position()


    # 耗时1    
    # 向量speed
    # 等价于标量direction和标量wheel_speed
    speed_x, speed_y, speed_z = get_speed()    


    # 耗时1   
    # 标量
    derection_angle = get_direction()
    # 耗时1    
    # 标量
    left_wheel_speed, right_wheel_speed = get_wheel_speed()    

    # 耗时1   
    # 标量
    left_wheel_force, right_wheel_force = get_wheel_force()    

    
    # 耗时1    
    ball[] = get_carried_balls()

    # 耗时1，也许可以废弃
    yes_or_no = is_loaded()

    # 耗时1    
    direction_angle, elevation_angle = get_aim()

    # 耗时随scan面积不同而不同
    # direction是方位角度，range是左右范围角度和，distance是扫描距离深度
    ball[], catapult[] = scan(direction, range, distance)


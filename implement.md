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
    pos.X
    pos.Y
    pos.Z

    left_wheel_speed.X
    left_wheel_speed.Y
    left_wheel_speed.Z

    right_wheel_speed.X
    right_wheel_speed.Y
    right_wheel_speed.Z

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

    pos.X
    pos.Y
    pos.Z

    speed.X
    speed.Y
    speed.Z


# user interface

## 控制

运动, 调用后保持，直到重新调用。

    # force > 0: go ahead
    # force = 0: free
    # force < 0: go back
    # 不同的调整幅度需要不同的耗时和耗能
    run(left_wheel_force, right_wheel_force)

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
    reload(material)

    # 不同的weight需要不同的耗时和耗能
    pick_up_ball(material, number)

    # 耗时1
    throw_away_ball(material, number)

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
    # 战车中心坐标 point3d
    position = get_position()


    # 耗时1
    # 向量speed vector3d
    # 等价于标量direction和标量wheel_speed
    speed = get_speed()


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

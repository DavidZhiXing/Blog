wifi充电时序

```sequence
Title: wifi充电
Note over AGV: 行驶至指定范围：400mm
Note over AGV: 行驶至终点未收到位信号则报异常
下位机->AGV: 检测到激光信号

AGV->下位机: 充电使能
下位机-->充电机: 建立连接
充电机->下位机: 到位
Note over 充电机: 5秒未检测到AGV
充电机-->充电机: 异常
Note over 充电机: 25秒未检测到AGV
充电机-->充电机: 异常

充电机->下位机: 伸缩到位
下位机-->AGV: 伸缩到位
AGV->下位机: 结束充电
下位机-->AGV: 伸缩杆收回到位
AGV->AGV: 离开

```

### Agv实现：

伺服定位：

1. 任务特性（ChargingRange)
2. 减速缓慢移动
3. 收到下位机信号停止
4. 发送充电使能

增加BreadownTester检测充电低速异常

修改离开逻辑

### 添加一个Tester

1. 选择继承的基类 `BreakDownTester/ SwitchBaseTester/`
2. 实现`Test`逻辑
3. (可能)`SwitchSensor`添加相应的属性
4. (可能)`SwitchSensorValue`添加相应的开关量
5. 配置`Breakdown.config`

### AGV与下位机通讯

1. `ChargingSlow`  //充电低速
2. `InChargingArea` //充电到位信号
3. `ChargingPile` //充电桩编号

### 新增配置

1. Agv.config `IsWifiCharging` 默认不为`true`，设为`False`则停止充电达到15秒后离开
2. Breakdowns.config `WifiChargingTester`

### DeviceBase 有价值的地方

告诉你如何添加一个设备

告诉你如何管理一个设备

告诉你如何删除一个设备

告诉你如何监控一个设备

###  一个Device有价值的地方

实现具体的业务

锂电池

铅酸电池

灯光

传感器

### 抽象类ConfigureBase有价值的地方

告诉你如何添加一个日志

告诉你如何删除一个日志

告诉你如何定制化输出

#### 具体配置类有价值的地方

配置相应的功能，业务
### Can转Wifi充电流程AGV配置

------

`Agv.config` 需要以下节点

``` xml
<setting Type="IsWifiCharging" Config="true" />
```



伺服定位需要设定起点离充电桩`400mm`，终点离充电桩`200mm`的`TaskMove`任务,需要添加的特性: 

```json
"Charging": true,
"ChargingPile": 1 // [充电桩编号]根据实际情况修改
```

`ErrorCode.config`需要更新: 

``` xml
<ErrorCodeItem Name="ChargingSlow" ErrorCode="0x240002D" Information="充电低速"  />
<ErrorCodeItem Name="WifiChargingError" ErrorCode="0x2400034" Information="伺服定位失败"  />
```

`Breakdowns.config`[可选]:

``` xml
<!-- 充电低速 -->
<WarningMonitor ErrorCode="0x220002D" Enable="true" Seconds="600" Threshold="3"/>
```




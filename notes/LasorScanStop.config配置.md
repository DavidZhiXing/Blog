###  `LasorScanStop.config` 配置

后激光通道

- 0：近距

  ``` xml
  <后激光通道>0</后激光通道>
  ```

  

- 1：中距

  ``` xml
  <后激光通道>1</后激光通道>
  ```

  

- 2：远距

  ``` xml
  <后激光通道>2</后激光通道>
  ```

  

- 3：安全激光

  ``` xml
  <后激光通道>3</后激光通道>
  ```

  

- -1：不开启

  ``` xml
  <后激光通道>-1</后激光通道>
  ```

  

**如何配置条件？**

前进

``` xml
<向前行驶>是</向前行驶>
```

前进屏蔽

``` xml
  <Config>
    <向前行驶>是</向前行驶>
    <后激光通道>-1</后激光通道>
  </Config>
```

角度

``` xml
<舵角范围>-0.1396,0.1396</舵角范围> // [-8°,8°]
```

直线后退

``` xml
  <Config>
    <向前行驶>否</向前行驶>
    <舵角范围>-0.1396,0.1396</舵角范围>
  </Config> 
```

直线后退启用近距

``` xml
  <Config>
    <向前行驶>否</向前行驶>
    <舵角范围>-0.1396,0.1396</舵角范围>
    <后激光通道>0</后激光通道>
  </Config> 
```

直线后退到取放货位置前（1.5m）启用近距

``` xml
  <Config>
    <向前行驶>否</向前行驶>
    <舵角范围>-0.1396,0.1396</舵角范围>
    <剩余距离>1.5,200</剩余距离>
    <后激光通道>0</后激光通道>
  </Config> 
```



后退左转弯

``` xml
  <Config>
  	<向前行驶>否</向前行驶>
    <舵角范围>-1.57,-0.1396</舵角范围>
  	<后激光通道>1</后激光通道>
  </Config>
```

后退右转弯

``` xml
 <Config>
   	<向前行驶>否</向前行驶>
   	<舵角范围>0.1396,1.57</舵角范围>
   	<后激光通道>2</后激光通道>
  </Config>
```

货叉高度 0到0.02米屏蔽

``` xml
  <Config>
    <货叉高度范围>0.0,0.02</货叉高度范围>
    <后激光通道>-1</后激光通道>
  </Config>
```


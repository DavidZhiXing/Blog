**驻车问题**

Parking位

下发没有使用Parking Enable output

下位机使用PLC

显示是在Function 构造

刷新呢？

- Action是一个委托，构造方法在Action里面
- 使用Action的多播
- 多播一组方法列表，刷新device==>sensor(bit)==>actuator(bit)
- 使用定时器，定时调用委托，刷新同步界面

**总结：**非常原始的方法，首先委托不好理解，递归非常深，每次我都需要找很久；

然后，继承混乱，一个UI为什么要继承Sensor接口？

还有一个问题，UI设计不合理，使用Grid但是，Grid却不是自适应的宽度的；

由此可见，开发者思维能力非常强；但是缺乏对WPF的真正理解，也缺乏泛型设计的深刻理解，在用WinForm的方式开发Wpf，其实也就是对MVVM模式理解不深

**解决办法**：使用WPF的绑定；

我的理想办法：

有一个实时的Sensors表，这个表某个线程定期更新，每次更新Invoke界面线程的`ViewModel`；

- 那么我们有没有这个sensors表呢，OK，我们去找找吧



**不可自启区域问题**

AutoRestartConfig==>配置文件的配置位

SafeHeloper==>

**Localization解决方法**

1. 通过编译项目以设置 x:Uid 并使用 LocBaml 工具实现；

   ​	实现步骤略微复杂

2. 通过 DynamicResource 实现；

   - 在 XAML 中，引用 DynamicResource 的属性必须为依赖属性，否则会出错；
   - 在 C# 代码中引用稍微有点麻烦，需要从 Resource Dictionary 中获取并转化为字符串

3.  通过 Resx 文件实现

   这是一种比较传统的、且普遍的方式，说它传统，是因为在 WinForm 中就可以这么做；说它普遍，是因为在 ASP.net MVC 中也可以这么做；并且，在 UWP 中的实现方式也与此有点类似

   [Resx实现Localization](./Resx实现Localization.md)

**TTL需求分析**

放货前后自适应

- 什么是线库？

- 前后的间距，可配置

- 检测位置，如何设定

- 如何检测？

- 如何平移前后？

- 结果包含什么？

- 库位后方如何定义

- 怎么算远

- AjustMask

- StorageAdaptiveOffset

  | 分析                      | 实际                                                         |
  | ------------------------- | ------------------------------------------------------------ |
  | 参考大前移的AjustMask流程 | AjustMask流程需要一些参数，是否需要添加，这些参数是什么意义，对流程有什么样的影响，这些都会花费大量的时间和精力 |
  |                           | Func<double> 这个函数是起什么作用，如何设定呢                |
  |                           | 如何平移货位                                                 |
  |                           | 如何进行货位检测                                             |
  |                           | 如何发送货位给控制模块                                       |
  |                           |                                                              |
  |                           |                                                              |
  |                           |                                                              |
  |                           |                                                              |

  通过查看代码，如何确定一段距离，

  要对大前移的流程改造，我发现有些东西我就不太清楚了：

  如何停止在指定的位置；

  如何检测货位

  如何对货位进行移动

  AjustMask是什么？

  目标平移怎么操作，X前后，Y左右，Z上下，P俯仰，C夹抱



高度自适应取货

- 如何绘制测试地图

- PerceptionWhileLifting-->平衡重

- Add VelocityLimitOfZ和AdjustMask

  | 分析                                   | 实际                                                         |
  | -------------------------------------- | ------------------------------------------------------------ |
  | 参考大前移的PerceptionWhileLifting流程 | 大前移的PerceptionWhileLifting流程非常复杂，不好移植，首先不确定哪些字段需要，哪些字段需要，这非常花费时间 |
  |                                        | 还要分析GetPerceptionStatus，复杂度+1                        |
  |                                        | 还要分析AddTaskStepAtTheLastPart， 复杂度+1                  |

  如何得到流程：

  已知：大前移代码，基本的框架，需求描述

  工具：逻辑思维，设计模式，经验，读代码，查找字符串，查找所有引用

  未知：大前移相关流程、代码量、代码思路、代码业务

  未知：如何移植流程到CounterBlanced

  行动：如何转换到对方的思维上去进行思考？

  

  ``` mermaid
  graph TD;
   A[start] --> B[GetPerceptionStatus]
   B --> B1[GetPathForPerception]
   subgraph  
   B1 --> B10{PerceptionHeight<0}
   B10 -- yes --> B11[计算不加俯仰的高度]
   B10 -- no --> B12[perception_z = PerceptionHeight]
   B11 --> B20{perception_z > MaxLiftZ}
   B20 -- yes --> B21[判断是否能完成]
   B21 --> B30{perception_z > MaxLiftZ}
   B30 -- yes --> B31[修改俯仰高度]
   end
   B31 -->C
   C[TaskStepMoveFork] --> D{ParkWhileMoveFork ?}
   D -- Yes --> E[TaskStepMoveForkWithMove]
   D -- No --> F[TaskStepMoveForkBeforeStop]
   E --> G[DetectPelletPositionWhileLifting]
   F --> G
   G --> H{AdjustMask}
   H -- Y --> H1[TaskStepMoveFork2 -- OrderY]
   H -- Z --> H2[TaskStepMoveFork2 -- OrderZ]
   H -- other --> H3[TaskStepMoveFork]
   H2 --> I[AddTaskStepAtTheLastPart]
   H1 --> I
   H3 --> I
  
   I --> I0{LastLengthForTraveling > 0}
    subgraph  
   I0 -- yes --> I10[CutLastLineOfBackwardPath]
   I0 -- no --> I2[TaskStepMoveForkWithMove]
   I10 --> I1[TaskStepMoveForkWithMove]
   I1 --> I11[TaskStepMove]
   end
   I11 --> End
   I2 --> End
   
  ```

  得到流程图之后，还需要判断哪些能移植，哪些不能移植，这些复杂度都跟流程图呈线性比例



高度自适应放货：

- 第一层放货
- 堆叠
- AdaptivePilling
- 限制货叉速度
- 没有货物的情况是什么意思，经常是什么意思？

删除TravelLengthForGet / TravelLengthForRack

问题：

1. DoCorrection()相应的逻辑现在是否需要修改

2. 大前移的流程比较复杂，我现在是全盘移植到CounterBalanced

   平衡重没有前后移和俯仰，那么我需要找出来那些是控制这些的流程

3. ForkMoveOder都是什么意思

### 更新Location

如何混合博世定位和相机定位？

光是这样分离接口是否可行？


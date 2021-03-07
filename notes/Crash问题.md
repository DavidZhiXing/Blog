### Crash问题

代码：

``` C# 
fixed (byte* dData = &(downData.ToArray()[0]), uData = &(upData.ToArray()[0]))
```

看看抛出的问题是什么：

``` c#
Terminating
引发类型为“System.OutOfMemoryException”的异常。
   在 System.Linq.Buffer`1..ctor(IEnumerable`1 source)
   在 System.Linq.Enumerable.ToArray[TSource](IEnumerable`1 source)
   在 A6ejffWZoqyuonkZFg.AQR9VcGujosmT3h4Dl.RfPKv6rMk(OqeHZ5g7JvAFyqMsgG  , Byte[]  , Byte[]  , Int32[]  , Int32[]  , Int64  , aMNDUc91MNcc8lF5gB&  ) 位置 D:\repositories\AGV\AGV.Function\PostureAnalyser\Correction.cs:行号 251
```

原理是这样：这段代码会分配一段内存，但是其他地方一直在消耗内存，当这边分配的时候找不到足够空间就报异常了，换句话说，其他只要需要分配内存的地方都会报。

吃内存的是这段配置引起的，暂时改为False

``` c#
FisheyeU.config Strategy.calibration=true
```



不过这个地方还是挺危险的....特别是内存不足的时候，容易触发垃圾回收，而这里因为固定内存了，不会移动，无法回收，性能堪忧。
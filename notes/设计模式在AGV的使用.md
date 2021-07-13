### 装饰者模式

扩展货叉功能

基本货叉：LiftZ;

左右：LiftZY;

前后：LiftZX;

俯仰：LiftZP

虚拟：LiftMock;

Decorator.P().Y().X().Z();

使用工厂模式创建：Decorator

``` C#
var decorator = DecoratorProvider.Get();
```

扩展：是否Move，Get, Put, Stop, Charge，Path，Scan, Localization都可以使用这种模式？但是这些是行为，第一直觉是使用策略模式。


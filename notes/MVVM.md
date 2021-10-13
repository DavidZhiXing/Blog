### MVVM

实现ViewModel的几种方式

1. `ViewModel`直接继承`INotifyPropertyChanged`

   `INotifyPropertyChanged`只有一个事件`PropertyChanged`

什么是Visua3d

Provides services and properties that are common to visual 3-D objects, including hit-testing, coordinate transformation, and bounding box calculations.

Unlike the [Model3D](https://docs.microsoft.com/en-us/dotnet/api/system.windows.media.media3d.model3d?view=net-5.0) class, [Visual3D](https://docs.microsoft.com/en-us/dotnet/api/system.windows.media.media3d.visual3d?view=net-5.0) objects cannot be shared or reused.

Access [Visual3D](https://docs.microsoft.com/en-us/dotnet/api/system.windows.media.media3d.visual3d?view=net-5.0) services by using static methods on the [VisualTreeHelper](https://docs.microsoft.com/en-us/dotnet/api/system.windows.media.visualtreehelper?view=net-5.0) class.

[Visual3D](https://docs.microsoft.com/en-us/dotnet/api/system.windows.media.media3d.visual3d?view=net-5.0) objects are optimized to be scene nodes. For example, they cache bounds. Whenever you can, use [Visual3D](https://docs.microsoft.com/en-us/dotnet/api/system.windows.media.media3d.visual3d?view=net-5.0) objects for unique instances of objects within your scene. This usage contrasts with that of [Model3D](https://docs.microsoft.com/en-us/dotnet/api/system.windows.media.media3d.model3d?view=net-5.0) objects, which are lightweight objects that are optimized to be shared and reused. For example, use a [Model3D](https://docs.microsoft.com/en-us/dotnet/api/system.windows.media.media3d.model3d?view=net-5.0) object to build a model of a car; and use ten [ModelVisual3D](https://docs.microsoft.com/en-us/dotnet/api/system.windows.media.media3d.modelvisual3d?view=net-5.0) objects to place ten cars in your scene.

设计原则是，一个场景一个Visual3D, 一个物体一个Model3d, 我的理解是，其他小孩都是Visual3D，然后Model3d添加到Visual3D上，提问: 那么一个Visual3D能否直接添加2个以上的Model3d?

那什么是ModelUIElement3D呢？与 ModelVisual3D有什么区别

为什么clone ModelVisual3D时需要Freeze呢？请看代码

``` c# 

            if (v is ModelVisual3D)
            {
                var m = v as ModelVisual3D;
                var clone = new ModelVisual3D();
                clone.Transform = m.Transform;
                if (m.Content != null && m.Content.CanFreeze)
                {
                    m.Content.Freeze();
                    var clonedModel = m.Content.Clone();
                    clone.Content = clonedModel;
                }

                if (m.Children.Count > 0)
                {
```



设置两个球体相交

``` xml
            <helix:SphereVisual3D Radius="5"/>
            <helix:SphereVisual3D Center="5,0,0" Radius="1" Fill="White"/>
```

**如何理解Center**

**如何切割**

剪掉的几何体

剩下的几何体

切割的时候需要Lock,需要遍历每个Chidren, 每个Chidre需要遍历每个三角形和位置，性能略低

**什么是Cutting Plane Group**

包含一组Plane

包含切割操作：内切，外切

**什么是Plane3D**

如何创建一个Plane? 

1. 两个Point
2. 一个Piont+一个Vector(法线)

**什么是Point3D**

表示Position， 一个point就是一个点，两个point就是一个直线，3个piont就是一个面，两个point跟3个point有交点point就是intersection



**什么是Vector3D**

表示法线Nomal



**为什么导入的3d中心点是（0,0,0）但是却不在世界坐标的（0，0，0）**

**Wiimote.lib是任天堂遥控器驱动程序**

**如何让Manipulator跟visual3d的center重合？**

**Wrong: [double].ToString()不会打double的完整位数**

为什么TestMethod没有提示，下次使用直接搜索MsTest
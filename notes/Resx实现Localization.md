### Resx实现Localization

它的操作步骤大概如下：

1. 创建一个 WPF 项目；

2. 在主窗口中添加几个需要本地化其内容的控件，如 TextBlock、Button等；

3. 展示此项目的 Properties 文件夹，这里已经有一个默认的 Resources.resx，在其中添加本地化字符串；

4. 将其复制，并粘贴到当前位置，将新文件改名为 Resources.en-US.resx，并修改其中的本地化字符串为对应的语言值；

5. 将上述两个 resx 文件的 Access Modifier 修改为 Public，默认值是 Internal；如下图：

   ![](./pics/676860-20170323150501299-972189791.png)

6. 在 Xaml 中添加命名空间引用：

```xaml
xmlns:loc="clr-namespace:WpfLocalizationTest.Properties"
```



7. 使用 {x:Static} 标记引用资源，如下：

```xaml
<TextBlock Text="{x:Static loc:Resources.Main_Menu_Home}" />
```



8. 如果需要在代码中引用，也非常简单：

```c#
var homePage = Properties.Resources.Main_Menu_Home;
```

此时，有两种方式可以测试本地化效果，第一种方式是在控制面板中改变操作系统的语言，第二种方式是通过代码改变 CurrentUICulture；这里我们使用后者，在 App.xaml.cs 中添加如下代码：

```c#
    protected override void OnStartup(StartupEventArgs e)
    {
        base.OnStartup(e);
        Thread.CurrentThread.CurrentUICulture = new System.Globalization.CultureInfo("en-US");
    }
```
**在大型项目中使用**

上面是在 Demo 中展示如何本地化应用，那么在大型项目中是如何来应用这种方式呢？以下简单介绍一下，可以供参考：

1. 在解决方案中创建一个类库项目（或在已有项目中进行后续操作），此项目的用途主要是作资源用；

2. 在其中创建 Localization 文件夹，在这个文件夹下可以再创建针对各个模块的子文件夹，然后在此创建 resx 文件、添加本地化字符串并复制。文件结构如下：

![](./pics/2.png)

3. 要在主程序中使用，方式和我们在 Demo 中提到的是一样的。首先在主程序项目中添加对新创建项目的引用，然后在 XAML 中添加命名空间和相关代码：

```xml
xmlns:loc2="clr-namespace:App.Resources.Localization.MainModule;assembly=App.Resources"
```

```xaml
<Button Margin="5" Padding="5" Content="{x:Static loc2:Resources.Main_New}" />
<Button Margin="5" Padding="5" Content="{x:Static loc2:Resources.Main_Open}" />
```

如果软件中有大量的文本需要本地化，那么在 resx 文件中的资源项将会非常多，这时，要在两个甚至更多资源文件中添加、删除、对比、检查项目时，将会非常困难。有没有更方便的方法呢？这里推荐一个 VS Extension: ResX Manager

![](./pics/3.png)
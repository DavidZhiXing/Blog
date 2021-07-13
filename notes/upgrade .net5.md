普通情况按照以下方法即可：（几个小时即可完成）

- [参考](https://docs.microsoft.com/en-us/dotnet/core/porting/upgrade-assistant-wpf-framework)

沙盒下的坑：

- `nuget package`问题

- 串口问题

- 试图加载格式不正确dll的问题

- 找不到config和dll问题

- wcf问题

  - .net5不支持WCF了
  - 是使用CoreWcf？
  - 还是gRpc
  - 或者.net core mvc?

- 切换回其他fx的分支问题

  > ##### Your project does not reference “.NETFramework,Version=v4.6.2” framework. Add a reference to “.NETFramework,Version=v4.6.2” in the “TargetFrameworks”

  Solution:

  ```cs
  git clean -xdf
  ```

  That should do the trick. It worked for us also in Jenkins. (we simply replayed the failed build with a modified script that ran git clean first).

  For some reason MSBuild / Visual Studio get confused when switching between branches that target different versions of .NET Framework, so I had to do git cleans regularly while switching between branches.

  Or:

  Please make the next steps

  1. Clean solution
  2. Clean folder "packages"
  3. Delete folder "bin"
  4. Delete folder "obj"
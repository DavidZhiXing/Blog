最近想要自动更新版本号：
- 目前看到的有两种方式，通过终端更新Assembly
- 通过Msbuild Task，（这个方式我有点不太熟悉）
## 终端方式
``` ps1
# Get the current date in the desired format (mmddyy)
$date = Get-Date -Format "yyMMdd"

# Get the current time in the desired format (HHmmss)
$time = Get-Date -Format "HHmmss"


# Construct the version string
$mainversion = "1.4.0."
$version = $mainversion + "*"
$fileversion = $mainversion + $date + "_alpha"

# Read the assembly.cs file
$assemblyFile = ".\Properties\AssemblyInfo.cs"
Write-Host "Reading file: $assemblyFile"
$content = Get-Content $assemblyFile

# Update the version information in the assembly.cs file
$content = $content -replace '(?<=AssemblyVersion\(").*?(?="\))', $version
$content = $content -replace '(?<=AssemblyFileVersion\(").*?(?="\))', $fileversion

# Powershell write original @ as ？， so need to fixed it
$content = $content -replace '(?<=AssemblyCopyright\(").*?(?="\))', "Copyright `@  2021"

# Write the updated content back to the assembly.cs file
Write-Host "Writing updated content to file: $assemblyFile"
$content | Set-Content $assemblyFile

Write-Host "Version updated successfully to: $version"
Write-Host "File version updated successfully to: $fileversion"

```
其中有了三个细节：
- 第一是Version不支持自定义格式，参考资料https://stackoverflow.com/questions/37941195/the-specified-version-string-does-not-conform-to-the-required-format-major-mi
- 第二个是 有些特殊字符串 需要进行处理，比如 @ 需要替换成 `@
- 第三个是 我修改的版本号 会产生警告信息需要使用 `#pragma warning disable CS7035` 屏蔽
以下是Bing的答案


[1]: https://adamtheautomator.com/powershell-replace/ "How to Use PowerShell Replace to Replace Text [Examples] - ATA Learning"
[2]: https://lazyadmin.nl/powershell/powershell-replace/ "How To use PowerShell Replace to replace a String or Character - LazyAdmin"
[3]: https://www.partitionwizard.com/partitionmanager/powershell-replace.html "How to Use PowerShell Replace to Replace String and Text"
[4]: https://shellgeek.com/powershell-replace-text-in-a-string/ "PowerShell – Replace text in a String [Examples] - ShellGeek"

I think you are trying to replace the © symbol in a PowerShell script. According to some web sources[^1^][1] [^2^][2], you can use the **replace () method** or the **-replace operator** to replace characters or strings in PowerShell. However, the © symbol may be a special character that needs to be escaped with a backtick (`) before using it in the replace function. For example:

```powershell
# Using the replace () method
$content = $content -replace '(?<=AssemblyCopyright(").*?(?="))', "`© 2021"

# Using the -replace operator
$content = $content.replace ("©", "`©")
```

You can also try using the **Unicode code point** of the © symbol, which is **U+00A9**, and prefix it with **\u**. For example:

```powershell
# Using the replace () method
$content = $content -replace '(?<=AssemblyCopyright(").*?(?="))', "\u00A9 2021"

# Using the -replace operator
$content = $content.replace ("©", "\u00A9")
```

I hope this helps you fix your script. You can find more information and examples on how to use PowerShell replace in the links below.

[^1^][1]: [How to Use PowerShell Replace to Replace Text [Examples] - ATA Learning](https://adamtheautomator.com/powershell-replace/)
[^2^][2]: [How To use PowerShell Replace to replace a String or Character - LazyAdmin](https://lazyadmin.nl/powershell/powershell-replace/)

## 使用MSBuild Task
参考资料 https://stackoverflow.com/questions/8739493/msbuild-for-updating-the-assemblyinfo-file

代码解读：

这段代码是一个 MSBuild 脚本，用于生成版本信息文件。它包含了一个名为 "SolutionInfo" 的目标（Target），在执行时会输出一条消息，并使用 AssemblyInfo 任务来生成版本信息文件。

具体解释如下：

- `<Target Name="SolutionInfo">`：定义了一个名为 "SolutionInfo" 的目标。
- `<Message Text="Creating Version File: $(Major).$(Minor).$(Build).$(Revision)"/>`：输出一条消息，内容为创建版本文件的信息。其中 `$(Major)`、`$(Minor)`、`$(Build)` 和 `$(Revision)` 是变量，表示主要版本号、次要版本号、构建号和修订号。
- `<AssemblyInfo>`：这是一个 MSBuild 任务，用于生成程序集的元数据（assembly metadata）。
    - `CodeLanguage="CS"`：指定编程语言为 C#。
    - `OutputFile="$(BuildInputDir)\SolutionInfo.cs"`：指定生成的文件路径和名称。
    - `AssemblyTitle="$(Company) $(Product)$(ProductAppendix)"`：设置程序集标题。
    - `AssemblyDescription="$(Company) $(Product)$(ProductAppendix)"`：设置程序集描述。
    - `AssemblyCompany="$(Company)"`：设置公司名称。
    - `AssemblyProduct="$(Product)"`：设置产品名称。
    - `AssemblyCopyright="Copyright © $(Company)"` ：设置版权信息，其中 "$(Company)" 是变量表示公司名称。
    - `ComVisible="false"` ：将 ComVisible 属性设为 false ，表示该程序集不可见性对 COM 组件无效，默认值是 true 代表可见性有效
    - `CLSCompliant="false"` ：将 CLSCompliant 属性设为 false ，表示该程序集不符合公共语言规范（Common Language Specification），默认值是 true 代表符合规范
    - `Guid="9E77382C-5FE3-4313-B099-7A9F24A4C328"`：设置程序集的 GUID。
    - `AssemblyVersion="$(Major).$(Minor).$(Build).$(Revision)"`：设置程序集版本号，使用之前定义的变量。
    - `AssemblyFileVersion="$(Major).$(Minor).$(Build).$(Revision)"`：设置文件版本号，使用之前定义的变量。



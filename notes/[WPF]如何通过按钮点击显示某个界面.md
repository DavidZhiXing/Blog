---

### 	传统方式

1.  ViewModel需要一个Bool变量ShowPanel
2. 需要一个Converter将ShowPanel的值进行转换，BoolToVissibilityConverter
3. 点击按钮绑定Command事件btnShowPannelCommand
4. Panel的Vissibility绑定ShowPanel
5. 需要的前置知识，
   - 如何定义一个ViewModel属性
   - 如何定义一个Command
   - 如何定义一个Converter
   - 如何绑定一个ViewModel的属性

---

### 	变化方式

 	1. 定义一个CheckBox
 	2. 把CheckBox的样式修改为Button
 	3. Panel的Visibility绑定CheckBox的Value
 	4. 前置知识：
     - 如何绑定其它Element的值
     - 如何自定义样式
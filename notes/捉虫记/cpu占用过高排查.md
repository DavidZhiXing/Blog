- 现场拿到dump
- 现场拿到pdb文件
- 导入windbg(x64)
- 在windbg中打开dump文件
- 导入pdb符合
- !runaway , 可以看到占用过高的线程

![](/notes/捉虫记/image/runaway.png)


- 排在前面的线程是占用过高的线程
- ~49 s 切换到高cpu现场

![](/notes/捉虫记/image/thread.png)


- alt + 6 查看调用堆栈

![](/notes/捉虫记/image/stack.png)

### 总结
---
发生了一个小错误，导致看不到堆栈信息；
- 没用加载调试符号pdb
- 使用64位模式分析32位的程序（由于沟通不当产生的误解）

windbg常用调试命令
``` c
!analyze -v  　　　　　　//找出出错的堆
.exrc    　　　　　　　　 //找到程序崩溃的位置
!heap    　　　　　　　　 //打印出错函数的局部位置
!for_each_frame dv /t  //显示call stack内容
~*kbn 　　　　　　　　　　//显示所有线程信息
~线程号 s      　　　　　//切换线程
kbn                　　//显示当前线程信息
.reload            　　//加载符号信息
!runaway 19    　　　　//查看19号线程所用时间
.load wow64exts
!sw                   //切换到64位
```





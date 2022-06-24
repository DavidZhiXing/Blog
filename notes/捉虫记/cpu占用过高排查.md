- 现场拿到dump
- 现场拿到pdb文件
- 导入windbg(x64)
- 在windbg中打开dump文件
- 导入pdb符合
- !runaway , 可以看到占用过高的线程
![](Blog/notes/捉虫记/image/runaway.png)
- 排在前面的线程是占用过高的线程
- ~49 s 切换到高cpu现场
![](Blog/notes/捉虫记/image/thread.png)
- alt + 6 查看调用堆栈
![](Blog/notes/捉虫记/image/stack.png)

### 关于使用dynamic的一些考虑

- 凡是用到Dynamic或DLR（动态语言运行时）的代码，都无法进行深度优化。因为性能优化往往需要把逻辑抽象层层剥离，而DLR就像套了一个巨大的抽象层。dynamic看起来简单，实则不然

  ``` c#
  static void Main(String args) {
      int a = 123;
      int b = 345;
      
      int c = a + b;
      Console.WriteLine(c);
  }
  ```

  这段代码的IL同样也很简单

  ``` c#
          static void Main(string[] args)
          {
  00007FFD4CFA0FB0  push        rbp  
  00007FFD4CFA0FB1  push        rdi  
  00007FFD4CFA0FB2  push        rsi  
  00007FFD4CFA0FB3  sub         rsp,40h  
  00007FFD4CFA0FB7  mov         rbp,rsp  
  00007FFD4CFA0FBA  xor         eax,eax  
  00007FFD4CFA0FBC  mov         dword ptr [rbp+34h],eax  
  00007FFD4CFA0FBF  mov         dword ptr [rbp+30h],eax  
  00007FFD4CFA0FC2  mov         dword ptr [rbp+2Ch],eax  
  00007FFD4CFA0FC5  mov         qword ptr [rbp+38h],rax  
  00007FFD4CFA0FC9  mov         qword ptr [rbp+60h],rcx  
  00007FFD4CFA0FCD  cmp         dword ptr [7FFD4D03FA98h],0  
  00007FFD4CFA0FD4  je          ConsoleApp1.Program.Main(System.String[])+02Bh (07FFD4CFA0FDBh)  
  00007FFD4CFA0FD6  call        00007FFDACBED820  
  00007FFD4CFA0FDB  nop  
              int a = 123;
  00007FFD4CFA0FDC  mov         dword ptr [rbp+34h],7Bh  
              int b = 345;
  00007FFD4CFA0FE3  mov         dword ptr [rbp+30h],159h  
  
              int c = a + b;
  00007FFD4CFA0FEA  mov         ecx,dword ptr [rbp+34h]  
  00007FFD4CFA0FED  add         ecx,dword ptr [rbp+30h]  
  00007FFD4CFA0FF0  mov         dword ptr [rbp+2Ch],ecx  
              Console.WriteLine(c);
  00007FFD4CFA0FF3  mov         ecx,dword ptr [rbp+2Ch]  
  00007FFD4CFA0FF6  call        方法存根对象: CLRStub[JumpStub]@7ffd4cfa1050 (07FFD4CFA0F70h)  
  00007FFD4CFA0FFB  nop  
          }
  ```

  

  然后替换成Dynamic

  ``` c#
  static void Main(String args) {
      dybamic a = 123;
      dybamic b = 345;
      
      dybamic c = a + b;
      Console.WriteLine(c);
  }
  ```

  查看其IL代码就知道即便简单WriteLine也变得不那么简单了，原来的代码则变成了一堆内存分配，委托，动态方法调用拼成的大杂烩，这些对象称为CallSite.

  ``` c#
     Console.WriteLine(c);
  00007FFD4D481311  mov         rcx,257B3082C50h  
  00007FFD4D48131B  cmp         qword ptr [rcx],0  
  00007FFD4D48131F  je          ConsoleApp2.Program.Main(System.String[])+0227h (07FFD4D481327h)  
  00007FFD4D481321  nop  
  00007FFD4D481322  jmp         ConsoleApp2.Program.Main(System.String[])+0347h (07FFD4D481447h)  
  00007FFD4D481327  mov         rcx,7FFD4D521B78h  
  00007FFD4D481331  call        00007FFDACF1B500  
  00007FFD4D481336  mov         qword ptr [rbp+0B8h],rax  
  00007FFD4D48133D  mov         dword ptr [rbp+0B4h],100h  
  00007FFD4D481347  mov         rcx,257B30830C0h  
  00007FFD4D481351  mov         rcx,qword ptr [rcx]  
  00007FFD4D481354  mov         qword ptr [rbp+0A8h],rcx  
  00007FFD4D48135B  xor         ecx,ecx  
  00007FFD4D48135D  mov         qword ptr [rbp+0A0h],rcx  
  00007FFD4D481364  mov         rcx,qword ptr [rbp+0B8h]  
  00007FFD4D48136B  call        00007FFDACF45760  
  00007FFD4D481370  mov         qword ptr [rbp+98h],rax  
  00007FFD4D481377  mov         rcx,7FFD4D547B48h  
  00007FFD4D481381  mov         edx,2  
  00007FFD4D481386  call        CORINFO_HELP_NEWARR_1_OBJ (07FFDACF878F0h)  
  00007FFD4D48138B  mov         qword ptr [rbp+90h],rax  
  00007FFD4D481392  xor         ecx,ecx  
  00007FFD4D481394  mov         dword ptr [rbp+8Ch],ecx  
  00007FFD4D48139A  mov         ecx,21h  
  00007FFD4D48139F  xor         edx,edx  
  00007FFD4D4813A1  call        方法存根对象: Microsoft.CSharp.RuntimeBinder.CSharpArgumentInfo.Create(Microsoft.CSharp.RuntimeBinder.CSharpArgumentInfoFlags, System.String) (07FFD4D4807D8h)  
  00007FFD4D4813A6  mov         qword ptr [rbp+80h],rax  
  00007FFD4D4813AD  mov         rcx,qword ptr [rbp+90h]  
  00007FFD4D4813B4  mov         edx,dword ptr [rbp+8Ch]  
  00007FFD4D4813BA  mov         r8,qword ptr [rbp+80h]  
  00007FFD4D4813C1  call        00007FFDACF869C0  
  00007FFD4D4813C6  mov         dword ptr [rbp+7Ch],1  
  00007FFD4D4813CD  xor         ecx,ecx  
  00007FFD4D4813CF  xor         edx,edx  
  00007FFD4D4813D1  call        方法存根对象: Microsoft.CSharp.RuntimeBinder.CSharpArgumentInfo.Create(Microsoft.CSharp.RuntimeBinder.CSharpArgumentInfoFlags, System.String) (07FFD4D4807D8h)  
  00007FFD4D4813D6  mov         qword ptr [rbp+70h],rax  
  00007FFD4D4813DA  mov         rcx,qword ptr [rbp+90h]  
  00007FFD4D4813E1  mov         edx,dword ptr [rbp+7Ch]  
  00007FFD4D4813E4  mov         r8,qword ptr [rbp+70h]  
  00007FFD4D4813E8  call        00007FFDACF869C0  
  00007FFD4D4813ED  mov         rcx,qword ptr [rbp+90h]  
  00007FFD4D4813F4  mov         qword ptr [rsp+20h],rcx  
  00007FFD4D4813F9  mov         ecx,dword ptr [rbp+0B4h]  
  00007FFD4D4813FF  mov         rdx,qword ptr [rbp+0A8h]  
  00007FFD4D481406  mov         r8,qword ptr [rbp+0A0h]  
  00007FFD4D48140D  mov         r9,qword ptr [rbp+98h]  
  00007FFD4D481414  call        方法存根对象: Microsoft.CSharp.RuntimeBinder.Binder.InvokeMember(Microsoft.CSharp.RuntimeBinder.CSharpBinderFlags, System.String, System.Collections.Generic.IEnumerable`1<System.Type>, System.Type, System.Collections.Generic.IEnumerable`1<Microsoft.CSharp.RuntimeBinder.CSharpArgumentInfo>) (07FFD4D480848h)  
  00007FFD4D481419  mov         qword ptr [rbp+68h],rax  
  00007FFD4D48141D  mov         rdx,qword ptr [rbp+68h]  
  00007FFD4D481421  mov         rcx,7FFD4D548758h  
  00007FFD4D48142B  call        方法存根对象: System.Runtime.CompilerServices.CallSite`1[[System.__Canon, System.Private.CoreLib]].Create(System.Runtime.CompilerServices.CallSiteBinder) (07FFD4D480768h)  
  00007FFD4D481430  mov         qword ptr [rbp+60h],rax  
  00007FFD4D481434  mov         rcx,257B3082C50h  
  00007FFD4D48143E  mov         rdx,qword ptr [rbp+60h]  
  00007FFD4D481442  call        CORINFO_HELP_CHECKED_ASSIGN_REF (07FFDACF86870h)  
  00007FFD4D481447  mov         rcx,257B3082C50h  
  00007FFD4D481451  mov         rcx,qword ptr [rcx]  
  00007FFD4D481454  mov         rcx,qword ptr [rcx+18h]  
  00007FFD4D481458  mov         qword ptr [rbp+50h],rcx  
  00007FFD4D48145C  mov         rcx,257B3082C50h  
  00007FFD4D481466  mov         rcx,qword ptr [rcx]  
  00007FFD4D481469  mov         qword ptr [rbp+48h],rcx  
  00007FFD4D48146D  mov         rcx,7FFD4D549B20h  
  00007FFD4D481477  call        00007FFDACF1B500  
  00007FFD4D48147C  mov         qword ptr [rbp+58h],rax  
  00007FFD4D481480  mov         rcx,qword ptr [rbp+58h]  
  00007FFD4D481484  call        00007FFDACF45760  
  00007FFD4D481489  mov         qword ptr [rbp+40h],rax  
  00007FFD4D48148D  mov         rcx,qword ptr [rbp+50h]  
  00007FFD4D481491  mov         qword ptr [rbp+30h],rcx  
  00007FFD4D481495  mov         rcx,qword ptr [rbp+30h]  
  00007FFD4D481499  mov         rcx,qword ptr [rcx+8]  
  00007FFD4D48149D  mov         rdx,qword ptr [rbp+48h]  
  00007FFD4D4814A1  mov         r8,qword ptr [rbp+40h]  
  00007FFD4D4814A5  mov         r9,qword ptr [rbp+120h]  
  00007FFD4D4814AC  mov         rax,qword ptr [rbp+30h]  
  00007FFD4D4814B0  call        qword ptr [rax+18h]  
  00007FFD4D4814B3  nop  
          }
  ```

  

  当然DLR有其方便性，对于快速开发和脚本编写而言，DLR确实非常优秀。

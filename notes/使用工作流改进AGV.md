### 使用工作流改进AGV

AGV的工作流建模

这个是`TaskMove`子流程：

```mermaid
   stateDiagram-v2
    state fork_state <<fork>>
      [*] --> TaskStepMove
      TaskStepMove --> fork_state
      fork_state --> TaskStepWait
      fork_state --> TaskStepWait(Senconds)
      state join_state <<join>>
      TaskStepWait --> join_state
      TaskStepWait(Senconds) --> join_state
      join_state --> [*]
```

代码就很简单了：

``` python
async def flow():
    await task('TaskStepMove')
    await fork_and_merge(
        seq(task('TaskStepWait'),
        seq(task('TaskStepWait', 60),
    )
    await task('TaskMove')
```

`TaskMoveLift`:

``` mermaid
   stateDiagram-v2
      [*] --> TaskStepMoveFork
      TaskStepMoveFork --> [*]
```

`TaskPutOfCounterBlanced`:

``` mermaid
   stateDiagram-v2
    state fork_state <<fork>>
      [*] --> fork_state
      fork_state --> InitAdaptivePiling
      fork_state --> InitAdaptiveMode
      fork_state --> InitSpaceAdaptiveMode
      fork_state --> TaskStepMoveFork
      TaskStepMoveFork --> TaskStepMoveForkBeforeStop
      TaskStepMoveForkBeforeStop --> TaskStepMoveFork2
      state join_state <<join>>
      TaskStepMoveFork2 --> join_state
      InitSpaceAdaptiveMode --> join_state
            InitAdaptiveMode --> join_state
      InitAdaptivePiling --> join_state
      join_state --> [*]
```


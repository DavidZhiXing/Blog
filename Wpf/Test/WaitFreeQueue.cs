public class WaitFreeQueue<T>
{
    private readonly object _lock = new object();
    private readonly Queue<T> _queue = new Queue<T>();
    private readonly Action<T> _action;
    private readonly Func<T, bool> _filter;
    private readonly Func<T, Task> _asyncAction;
    private readonly Func<T, Task<bool>> _asyncFilter;
    private readonly TaskCompletionSource<bool> _tcs = new TaskCompletionSource<bool>();
    private bool _isCompleted;

    public WaitFreeQueue(Action<T> action)
    {
        _action = action;
    }

    public WaitFreeQueue(Func<T, bool> filter)
    {
        _filter = filter;
    }

    public WaitFreeQueue(Func<T, Task> asyncAction)
    {
        _asyncAction = asyncAction;
    }

    public WaitFreeQueue(Func<T, Task<bool>> asyncFilter)
    {
        _asyncFilter = asyncFilter;
    }

    public void Enqueue(T item)
    {
        lock (_lock)
        {
            if (_isCompleted)
            {
                throw new InvalidOperationException("Queue is completed");
            }

            if (_filter != null)
            {
                if (_filter(item))
                {
                    _queue.Enqueue(item);
                }
            }
            else if (_asyncFilter != null)
            {
                _asyncFilter(item).ContinueWith(t =>
                {
                    if (t.Result)
                    {
                        _queue.Enqueue(item);
                    }
                }, TaskContinuationOptions.ExecuteSynchronously);
            }
            else
            {
                _queue.Enqueue(item);
            }
        }
    }

    public void Complete()
    {
        lock (_lock)
        {
            if (_isCompleted)
            {
                throw new InvalidOperationException("Queue is completed");
            }

            _isCompleted = true;
            _tcs.SetResult(true);
        }
    }

    public void Wait()
    {
        Task.Run(async () =>
        {
            while (true)
            {
                T item;
                lock (_lock)
                {
                    if (_queue.Count == 0)
                    {
                        if (_isCompleted)
                        {
                            return;
                        }

                        await _tcs.Task;
                    }

                    item = _queue.Dequeue();
                }

                if (_action != null)
                {
                    _action(item);
                }
                else if (_asyncAction != null)
                {
                    await _asyncAction(item);
                }
            }
        });
    }

    public void Dispose()
    {
        lock (_lock)
        {
            _isCompleted = true;
            _tcs.SetResult(true);
        }
    }

    public void Clear()
    {
        lock (_lock)
        {
            _queue.Clear();
        }
    }

    public int Count
    {
        get
        {
            lock (_lock)
            {
                return _queue.Count;
            }
        }
    }

    public bool IsCompleted
    {
        get
        {
            lock (_lock)
            {
                return _isCompleted;
            }
        }
    }

    public bool IsEmpty
    {
        get
        {
            lock (_lock)
            {
                return _queue.Count == 0;
            }
        }
    }

    public T Dequeue()
    {
        lock (_lock)
        {
            return _queue.Dequeue();
        }
    }

    public T Peek()
    {
        lock (_lock)
        {
            return _queue.Peek();
        }
    }

    public void Enqueue(IEnumerable<T> items)
    {
        lock (_lock)
        {
            foreach (var item in items)
            {
                _queue.Enqueue(item);
            }
        }
    }

    public void Enqueue(params T[] items)
    {
        lock (_lock)
        {
            foreach (var item in items)
            {
                _queue.Enqueue(item);
            }
        }
    }

}
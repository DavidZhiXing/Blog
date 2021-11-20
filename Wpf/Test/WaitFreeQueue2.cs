public class WaitFreeQueue2<T>
{
    private Node<T> _head;
    private Node<T> _tail;

    public bool Dequeue(ref T value)
    {
        while (true)
        {
            Node<T> node = _head;
            if (node == null)
            {
                return false;
            }

            if (Interlocked.CompareExchange(ref _head, node.Next, node) == node)
            {
                value = node.Value;
                return true;
            }
        }
    }

    public void Enqueue(T value)
    {
        Node<T> node = new Node<T>(value);
        while (true)
        {
            Node<T> tail = _tail;
            node.Next = tail;
            if (Interlocked.CompareExchange(ref _tail, node, tail) == tail)
            {
                return;
            }
        }
    }

    public WaitFreeQueue2()
    {
        _head = _tail = new Node<T>();
    }

    public WaitFreeQueue2(IEnumerable<T> collection)
    {
        _head = _tail = new Node<T>();
        foreach (T value in collection)
        {
            Enqueue(value);
        }
    }

    private class Node<TData>
    {
        public TData Value;
        public Node<TData> Next;

        public Node(TData value)
        {
            Value = value;
        }
    }
}
并查集是一种数据结构，它可以用来解决一些元素分组的问题。它管理一系列不相交的集合，并支持两种操作：合并两个集合和查询一个元素属于哪个集合。

以下是一个 Python 的并查集实现示例：

```python
class UnionFind:
    def __init__(self, n):
        self.parent = list(range(n))
        self.rank = [0] * n

    def find(self, x):
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])
        return self.parent[x]

    def union(self, x, y):
        rootx = self.find(x)
        rooty = self.find(y)
        if rootx != rooty:
            if self.rank[rootx] > self.rank[rooty]:
                self.parent[rooty] = rootx
            else:
                self.parent[rootx] = rooty
                if self.rank[rootx] == self.rank[rooty]:
                    self.rank[rooty] += 1
```

在这个示例中，`UnionFind` 类有两个主要的方法：
- `find(x)`：找到元素 `x` 所在集合的代表元素，也就是根节点。
- `union(x, y)`：将元素 `x` 和 `y` 所在的两个集合合并为一个集合。


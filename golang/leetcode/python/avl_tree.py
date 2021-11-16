"""
Implementation of an auto-balanced binary search tree.
for doctest run following command:
    python3 -m doctest -v avl_tree.py
for testing run:
    python avl_tree.py

"""

from __future__ import annotations

import random
import math
from typing import Any

class my_queue:
    def __init__(self) -> None:
        self.data: list[Any] = []
        self.head: int = 0
        self.tail: int = 0

    def is_empty(self) -> bool:
            return self.head == self.tail

    def push(self, item: Any) -> None:
        self.data.append(item)
        self.tail += 1

    def pop(self) -> Any:
        if self.is_empty():
            raise Exception("Queue is empty")
        item = self.data[self.head]
        self.head += 1
        return item

    def count(self) -> int:
        return self.tail - self.head

    def print_queue(self) -> None:
        print(self.data)
        print("**************")
        print(self.data[self.head:self.tail])

class my_node:
    def __init__(self, data: Any) -> None:
        self.left: my_node = None
        self.right: my_node = None
        self.height: int = 1
        self.data = data

    def __str__(self) -> str:
        return f"{self.key}:{self.value}"

    def __repr__(self) -> str:
        return f"{self.key}:{self.value}"

    def get_data(self) -> Any:
        return self.data

    def get_left(self) -> my_node:
        return self.left

    def get_right(self) -> my_node:
        return self.right

    def set_data(self, data: Any) -> None:
        self.data = data

    def set_left(self, node: my_node) -> None:
        self.left = node

    def set_right(self, node: my_node) -> None:
        self.right = node

    def get_height(self) -> int:
        return self.height

    def set_height(self, height: int) -> None:
        self.height = height
        return

def get_height(node: my_node | None) -> int:
    if node is None:
        return 0
    return node.get_height()

def my_max(a: int, b: int) -> int:
    return a if a > b else b

def right_rotate(node: my_node) -> my_node:
    """
    Rotate node to right.

         A                           B
        / \                         / \
        B  C        -->            D   A
        / \                           / \
        D   E                         E   C

    """
    left_node = node.get_left()
    node.set_left(left_node.get_right())
    left_node.set_right(node)

    node.set_height(my_max(get_height(node.get_left()), get_height(node.get_right())) + 1)
    left_node.set_height(my_max(get_height(left_node.get_left()), get_height(left_node.get_right())) + 1)

    return left_node

def left_rotate(node: my_node) -> my_node:
    """
    Rotate node to left.

         A                           B
        / \                         / \
        B  C        -->            A   C
           / \                       / \
          D   E                     B   D

    """
    right_node = node.get_right()
    node.set_right(right_node.get_left())
    right_node.set_left(node)

    node.set_height(my_max(get_height(node.get_left()), get_height(node.get_right())) + 1)
    right_node.set_height(my_max(get_height(right_node.get_left()), get_height(right_node.get_right())) + 1)

    return right_node

def lr_rotate(node: my_node) -> my_node:
    """
    Left-Right rotation.

         A                           A                  Br
        / \         LR              / \     RR         / \
        B  C        -->            Br  C    -->       B   A
       / \                        / \                 /   / \
      Bl  Br                     B  UB              Bl   UB  C
           \                    /
           UB                  Bl
    """
    node.set_left(right_rotate(node.get_left()))
    return left_rotate(node)

def rl_rotate(node: my_node) -> my_node:
    """
    Right-Left rotation.

         A                           A                  Br
        / \         RL              / \     LR         / \
        B  C        -->            B   C    -->       A   B
           / \                    / \                 /   / \
          UB  Bl                 UB  Bl              C   Bl  Br
        """
    node.set_right(left_rotate(node.get_right()))
    return right_rotate(node)

def insert(node: my_node, data: Any) -> my_node:
    if node is None:
        return my_node(data)

    if data < node.get_data():
        node.set_left(insert(node.get_left(), data))
    else:
        node.set_right(insert(node.get_right(), data))

    node.set_height(my_max(get_height(node.get_left()), get_height(node.get_right())) + 1)

    balance = get_height(node.get_left()) - get_height(node.get_right())

    if balance > 1:
        if data < node.get_left().get_data():
            node = right_rotate(node)
        else:
            node = lr_rotate(node)
    elif balance < -1:
        if data > node.get_right().get_data():
            node = left_rotate(node)
        else:
            node = rl_rotate(node)

    return node

def get_right_most(node: my_node) -> my_node:
    while node.get_right() is not None:
        node = node.get_right()
    return node

def get_left_most(node: my_node) -> my_node:
    while node.get_left() is not None:
        node = node.get_left()
    return node

def delete(node: my_node, data: Any) -> my_node:
    if node is None:
        return node

    if data < node.get_data():
        node.set_left(delete(node.get_left(), data))
    elif data > node.get_data():
        node.set_right(delete(node.get_right(), data))
    else:
        if node.get_left() is None:
            return node.get_right()
        elif node.get_right() is None:
            return node.get_left()

        right_most = get_right_most(node.get_left())
        node.set_data(right_most.get_data())
        node.set_left(delete(node.get_left(), right_most.get_data()))

    node.set_height(my_max(get_height(node.get_left()), get_height(node.get_right())) + 1)

    balance = get_height(node.get_left()) - get_height(node.get_right())

    if balance > 1:
        if get_height(node.get_left().get_left()) >= get_height(node.get_left().get_right()):
            node = right_rotate(node)
        else:
            node = lr_rotate(node)
    elif balance < -1:
        if get_height(node.get_right().get_right()) >= get_height(node.get_right().get_left()):
            node = left_rotate(node)
        else:
            node = rl_rotate(node)

    return node

class AVLtrree:
    """
    An AVL tree doctest
    Example:
    >>> t = AVLtree()
    >>> t.insert(4)
    insert 4
    >>> print(str(t).replace("\\n", "\n"))
    4
    >>> t.insert(2)
    insert 2
    >>> print(str(t).replace("\\n", "\n"))
    2
    4
    2
    >>> t.insert(6)
    insert 6
    >>> print(str(t).replace("\\n", "\n"))
    2
    4
    6
    2
    6
    4
    >>> t.insert(1)
    insert 1
    >>> print(str(t).replace("\\n", "\n"))
    1
    2
    4
    6
    2
    6
    4
    1
    >>> t.insert(3)
    insert 3
    >>> print(str(t).replace("\\n", "\n"))
    1
    2
    3
    4
    6
    2
    6
    4
    1
    3
    >>> t.insert(5)
    insert 5
    >>> print(str(t).replace("\\n", "\n"))
    1
    2
    3
    4
    5
    6
    2
    6
    4
    1
    3
    5
    >>> t.insert(7)
    insert 7
    >>> print(str(t).replace("\\n", "\n"))
    1
    2
    3
    4
    5
    6
    7
    2
    6
    4
    1
    3
    5
    7
    >>> t.insert(8)
    insert 8
    >>> print(str(t).replace("\\n", "\n"))
    1
    2
    3
    4
    5
    6
    7
    8
    2
    6
    4
    1
    3
    5
    7
    8
    >>> t.insert(9)
    insert 9
    >>> print(str(t).replace("\\n", "\n"))
    1
    2
    3
    4

    """
    def __init__(self):
        self.root: my_node | None = None

    def insert(self, data: Any) -> None:
        self.root = insert(self.root, data)

    def __str__(self) -> str:
        output = ""
        queue = my_queue()
        queue.push(self.root)
        layer = self.get_height()
        if layer == 0:
            return output
        cnt = 0
        while not queue.is_empty():
            node = queue.pop()
            space = " " * (layer - cnt)
            output += space
            if node is not None:
                output += str(node.get_data())
                if cnt == layer:
                    output += str(node.get_data()) + "\n"
                    cnt = 0
                else:
                    queue.push(node.get_left())
                    queue.push(node.get_right())
                    cnt += 1
            else:
                output += "*"
                if cnt == layer:
                    output += "\n"
                    cnt = 0
                else:
                    queue.push(None)
                    queue.push(None)
                    cnt += 1

            output += space
            for i in range(100):
                output += " "
        return output

    def get_height(self) -> int:
        return get_height(self.root)

def _test() -> None:
    import doctest
    doctest.testmod()

def main() -> None:
    _test()
    t = AVLtrree()
    t.insert(4)
    t.insert(2)
    t.insert(6)
    t.insert(1)
    t.insert(3)
    t.insert(5)
    t.insert(7)
    t.insert(8)
    t.insert(9)
    print(t)
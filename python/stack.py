
from typing import Generic, List, TypeVar


T = TypeVar('T')

class StackOverflowError(Exception):
    pass

class StackUnderflowError(Exception):
    pass

class Stack(Generic[T]):
    def __init__(self, limit: int) -> None:
        self.stack: List[T] = []
        self.limit: int = limit

    def push(self, data: T) -> None:
        if len(self.stack) == self.limit:
            raise StackOverflowError
        self.stack.append(data)

    def pop(self) -> T:
        if len(self.stack) == 0:
            raise StackUnderflowError
        return self.stack.pop()

    def __len__(self) -> int:
        return len(self.stack)

    def __repr__(self) -> str:
        return f'Stack({self.stack})'
        
    def peek(self) -> T:
        if len(self.stack) == 0:
            raise StackUnderflowError
        return self.stack[-1]

    def is_empty(self) -> bool:
        return len(self.stack) == 0

    def is_full(self) -> bool:
        return len(self.stack) == self.limit

    def size(self) -> int:
        return len(self.stack)

    def __contains__(self, item: T) -> bool:
        return item in self.stack

    def __bool__(self) -> bool:
        return bool(self.stack)

# Test
def test_stack() -> None:
    stack = Stack(3)
    stack.push(1)
    stack.push(2)
    stack.push(3)
    assert stack.pop() == 3
    assert stack.pop() == 2
    assert stack.pop() == 1
    assert stack.is_empty()
    stack.push(1)
    stack.push(2)
    stack.push(3)
    assert stack.pop() == 3
    assert stack.pop() == 2
    assert stack.pop() == 1
    assert stack.is_empty()
    stack.push(1)
    stack.push(2)
    stack.push(3)
    assert stack.pop() == 3
    assert stack.pop() == 2
    assert stack.pop() == 1
    assert stack.is_empty()
    stack.push(1)
    stack.push(2)
    stack.push(3)
    assert stack.pop() == 3
    assert stack.pop() == 2
    assert stack.pop() == 1
    assert stack.is_empty()
    stack.push(1)
    stack.push(2)
    stack.push(3)
    assert stack.pop() == 3
    assert stack.pop() == 2
    assert stack.pop() == 1
    assert stack.is_empty()
    stack.push(1)
    stack.push(2)
    stack.push(3)
    assert stack.pop() == 3
    assert stack.pop() == 2
    assert stack.pop() == 1
    assert stack.is_empty()
    stack.push(1)
    stack.push(2)
    stack.push(3)
    assert stack.pop() == 3
    assert stack.pop() == 2
    
    try:
        stack.pop()
        assert False
    except StackUnderflowError:
        assert True
    try:
        stack.peek()
        assert False
    except StackUnderflowError:
        assert True

if __name__ == '__main__':
    test_stack()
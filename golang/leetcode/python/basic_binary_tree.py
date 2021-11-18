from __future__ import annotations

class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

def display(node):
    if node is None:
        return
    print(node.val, end=' ')
    display(node.left)
    display(node.right)

def depth(node):
    if node is None:
        return 0
    return 1 + max(depth(node.left), depth(node.right)) 

def is_full(node):
    if node is None:
        return True
    if node.left is None and node.right is None:
        return True
    if node.left is None or node.right is None:
        return False
    return is_full(node.left) and is_full(node.right)

def main():
    root = TreeNode(1)
    root.left = TreeNode(2)
    root.right = TreeNode(3)
    root.left.left = TreeNode(4)
    root.left.right = TreeNode(5)
    root.right.left = TreeNode(6)
    root.right.right = TreeNode(7)
    root.left.left.left = TreeNode(8)
    root.left.left.right = TreeNode(9)
    root.left.right.left = TreeNode(10)
    root.left.right.right = TreeNode(11)
    root.right.left.left = TreeNode(12)
    root.right.left.right = TreeNode(13)
    root.right.right.left = TreeNode(14)
    root.right.right.right = TreeNode(15)
    print('is_full: ', is_full(root))
    print('depth: ', depth(root))
    display(root)

if __name__ == '__main__':
    main()
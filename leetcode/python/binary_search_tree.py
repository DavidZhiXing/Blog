from pprint import pformat


class Node:
    def __init__(self, val):
        self.val = val
        self.left = None
        self.right = None

    def __repr__(self):
        from pprint import pprint
        if self.left is None and self.right is None:
            return str(self.val)
        return pformat({self.val: {'left': self.left, 'right': self.right}})

    def __str__(self):
        return str(self.val)

class BinarySearchTree:
    def __init__(self, root = None) -> None:
        self.root = root

    def insert(self, val):
        if self.root is None:
            self.root = Node(val)
        else:
            self.__insert(val, self.root)

    def __insert(self, val, node):
        if val < node.val:
            if node.left is None:
                node.left = Node(val)
            else:
                self.__insert(val, node.left)
        else:
            if node.right is None:
                node.right = Node(val)
            else:
                self.__insert(val, node.right)

    def search(self, val):
        if self.root is None:
            return False
        return self.__search(val, self.root)

    def __search(self, val, node):
        if node is None:
            return False
        if node.val == val:
            return True
        if val < node.val:
            return self.__search(val, node.left)
        else:
            return self.__search(val, node.right)

    def delete(self, val):
        if self.root is None:
            return False
        return self.__delete(val, self.root)

    def __delete(self, val, node):
        if node is None:
            return False
        if val < node.val:
            node.left = self.__delete(val, node.left)
        elif val > node.val:
            node.right = self.__delete(val, node.right)
        else:
            if node.left is None:
                return node.right
            elif node.right is None:
                return node.left
            else:
                node.val = self.__find_min(node.right)
        return node

    def __find_min(self, node):
        if node.left is None:
            return node.val
        return self.__find_min(node.left)

    def __str__(self):
        return str(self.root)

    def get_max(self, node=None):
        if node is None:
            node = self.root
        if node.right is None:
            return node.val
        return self.get_max(node.right)

    def get_min(self, node=None):
        if node is None:
            node = self.root
        if node.left is None:
            return node.val
        return self.get_min(node.left)

    def preorder_traverse(self, node):
        if node is not None:
            yield node
            yield from self.preorder_traverse(node.left)
            yield from self.preorder_traverse(node.right)

    def inorder_traverse(self, node):
        if node is not None:
            yield from self.inorder_traverse(node.left)
            yield node
            yield from self.inorder_traverse(node.right)

    def postorder_traverse(self, node):
        if node is not None:
            yield from self.postorder_traverse(node.left)
            yield from self.postorder_traverse(node.right)
            yield node

    def inorder(self, arr: list, node: Node):
        if node is not None:
            self.inorder(arr, node.left)
            arr.append(node.val)
            self.inorder(arr, node.right)   

    def find_kth_largest(self, k):
        arr = []
        self.inorder(arr, self.root)
        return arr[-k]

    def find_kth_smallest(self, k):
        arr = []
        self.inorder(arr, self.root)
        return arr[k-1]

def postorder(curr_node):
    node_list = []
    if curr_node is not None:
        node_list.extend(postorder(curr_node.left))
        node_list.extend(postorder(curr_node.right))
        node_list.append(curr_node.val)
    return node_list

def binary_search_tree():
    testList = [5, 3, 6, 2, 4, None, 7]
    t = BinarySearchTree()
    for i in testList:
        t.insert(i)

    print(t)

    if t.search(3):
        print("Found")
    else:
        print("Not Found")

    t.delete(3)
    print(t)

    print(t.get_max())
    print(t.get_min())

    print(t.find_kth_largest(2))
    print(t.find_kth_smallest(2))

    print(list(t.preorder_traverse(t.root)))
    print(list(t.inorder_traverse(t.root)))
    print(list(t.postorder_traverse(t.root)))

    print(postorder(t.root))

if __name__ == "__main__":
    import doctest
    doctest.testmod()

    #binary_search_tree()

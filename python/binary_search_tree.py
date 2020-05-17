
class Node:
    def __init__(self, key, value):
        self.key = key
        self.value = value
        self.left = None
        self.right = None


class BST:
    def __init__(self):
        self.start_node = Node('start', 'value')

    def put(self, key, value):
        self.put_in_space(self.start_node, key, value)

    def put_in_space(self, node, key, value):
        if node.left is None:
            node.left = Node(key, value)
        elif node.right is None:
            node.right = Node(key, value)
        else:
            self.put_in_space(node.left, key, value)


b = BST()
b.put('key1', 'value1')
b.put('key2', 'value2')
b.put('key3', 'value3')
b.put('key4', 'value4')
b.put('key5', 'value5')

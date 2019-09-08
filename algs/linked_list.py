class Node:
    def __init__(self, data):
        self.data = data
        self.next = None

    def add_node(self, data):
        while 

    def __iter__(self):
        return self

    def __next__(self):
        return self.next


start_node = Node('start_data')
second_node = start_node.add_node('second_node')
third_node = second_node.add_node('third_node')
node = start_node
seen_data = {node.data: True}
while node is not None:
    print(node.data)
    node = node.__next__()


class Stack:
    def __init__(self, first_element):
        self.data = []
        self.current_min = first_element

    def add(self, item: int):
        if item < self.current_min:
            self.current_min = item
        self.data.append(item)

    def pop(self) -> int:
        return self.data.pop()

    def min_(self) -> int:
        return self.current_min


stack = Stack(0)
stack.add(1)
stack.add(2)
print(stack.min_())  # -> 0
print(stack.pop())  # -> 2

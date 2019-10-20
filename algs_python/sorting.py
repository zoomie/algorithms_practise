# Selection sort
def get_smallest(ls):
    smallest = 1000
    index = None
    for i, num in enumerate(ls):
        if num < smallest:
            smallest = num
            index = i
    return ls.pop(index)


ls = [4, 2, 1, 3]
sorted_list = []
for _ in range(len(ls)):
    sorted_list.append(get_smallest(ls))


# Insertion sort

ls = [4, 2, 1]

if ls[1] < ls[0]:
    ls[0], ls[1] = ls[1], ls[0]

if ls[2] < ls[1]:
    pass

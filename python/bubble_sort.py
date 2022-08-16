from random import randint


def sort(arr):
    has_swapped = True
    while has_swapped:
        has_swapped = False
        for i in range(len(arr)-1):
            if arr[i] > arr[i+1]:
                arr[i], arr[i+1] = arr[i+1], arr[i]
                has_swapped = True
    return arr

arr = [randint(1, 100) for _ in range(20)]
print(sort(arr))
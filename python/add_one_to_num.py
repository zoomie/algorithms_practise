def plus_one(digits: list[int]) -> list[int]:
    for i in range(len(digits) -1, -1, -1):
        if digits[i] == 9:
            digits[i] = 0
        else:
            digits[i] += 1
            return digits              
    return [1] + digits

print(plus_one([1, 2, 3, 4])) #[1, 2, 3, 5]
print(plus_one([9, 9, 9, 9])) #[1, 0, 0, 0, 0] 
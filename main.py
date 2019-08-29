# Insert
# Remove
# Replace

from typing import List

str1 = 'man'
str2 = 'mun'


def change(str1: str, str2: str) -> List[str]:
    operations: List[str] = []
    for char1, char2 in zip(str1, str2):
        if char1 != char2:
            operations.append("replace")
    return operations


print(change(str1, str2))

# Insert
# Remove
# Replace

from typing import List


def edit_string(str1: str, str2: str) -> List[str]:
    operations: List[str] = []

    for char1, char2 in zip(str1, str2):
        if char1 != char2:
            operations.append("replace")
    return operations


str1 = 'man'
str2 = 'mun'
print(edit_string(str1, str2))

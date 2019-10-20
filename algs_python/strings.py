# 1 Does the string have unique characters


def unique_chars(string: str) -> bool:
    result = False
    values_seen = {}
    for char in string:
        if values_seen.get(char):
            result = True
            break
        values_seen[char] = True
    return result


# print(unique_chars('abca'))

# 3 Determine if one string is a permutation of the other

def string_permutation(sub_string: str, main_str: str) -> bool:
    length = len(sub_string)
    for i, _ in enumerate(main_str):
        temp_str = main_str[i: i + length]
        if temp_str == sub_string:
            return True


string_permutation('abc', 'abcde')

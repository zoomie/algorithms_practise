def split(nums):
    result = []
    pairs = ''
    for i in range(len(nums)):
        pairs += nums[i]
        if i+1 < len(nums) and nums[i+1] != nums[i]:
            result.append(pairs)
            pairs = ''
    result.append(pairs)
    return result

print(split('112233')) # ['11', '22', '33']



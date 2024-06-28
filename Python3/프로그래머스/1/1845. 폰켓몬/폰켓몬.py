def solution(nums):
    newDict = {}
    for num in nums:
        newDict[num] = True
    
    keys = 0
    for _ in newDict:
        keys += 1
        
    if len(nums) / 2 < keys:
        return len(nums) / 2
    else:
        return keys
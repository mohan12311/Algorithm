def solution(arr):
    answer = []
    prev = -1
    for number in arr:
        if number == prev:
            pass
        else:
            answer.append(number)
            prev = number
    return answer
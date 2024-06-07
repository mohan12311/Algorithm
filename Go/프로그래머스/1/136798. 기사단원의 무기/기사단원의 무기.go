// number 0 < number < 100_001
// 자기 번호의 약수 개수에 해당하는 공격력을 구매
// 상한치를 초과하면 => 함수에 따른 무기를 리턴
// checker = (약수개수) => {
//   if 약수개수 > limit { return power }
// }
// 공격력 1당 => 1kg 필요

import "math"


func solution(number int, limit int, power int) int {
    result := 0
    for i := 1; i <= number; i++ {
        toUpdate := measure(i)
        if toUpdate > limit {
            result += power
        } else {
            result += toUpdate
        }
    }
    return result
}

func measure(num int) int {
    count := 0
    sqrt := int(math.Sqrt(float64(num)))
    
    for i := 1; i <= sqrt; i++ {
        if num % i == 0 {
            if i == num / i {
                count++
            } else {
                count += 2
            }
        }
    }
    
    return count
}
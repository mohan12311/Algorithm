import (
	"strconv"
	"strings"
)

func solution(n int, k int) int {
	// Convert n to base k
	baseK := strconv.FormatInt(int64(n), k)
	
	// Split the string by '0' and count primes
	parts := strings.Split(baseK, "0")
	count := 0
	for _, part := range parts {
		if part != "" {
			num, _ := strconv.ParseInt(part, 10, 64)
			if isPrime(int(num)) {
				count++
			}
		}
	}
	
	return count
}

func isPrime(n int) bool {
	if n <= 1 {
		return false }
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
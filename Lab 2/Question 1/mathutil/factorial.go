package mathutil

func Fact(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

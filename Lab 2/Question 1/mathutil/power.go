package mathutil

func Pow(base, exponent int) int {
	if exponent < 0 {
		return -1
	}
	if exponent == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= exponent; i++ {
		result *= base
	}
	return result
}

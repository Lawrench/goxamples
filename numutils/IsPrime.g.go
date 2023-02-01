package numutils

import "math"

// IsPrime returns true if n is a prime number.
// Any number less than 2 is not a prime number.
// Any number greater than the square root of n is not a prime number.
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

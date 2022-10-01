/* ***** 007: 10_001st prime ******
 *
 * By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
 * we can see that the 6th prime is 13.
 *
 * What is the 10 001st prime number?
 *
 */

package basics

import (
	"math"
)

func isPrime(number int) bool {
	limit := math.Ceil(math.Sqrt(float64(number)))
	for i := float64(2); i <= limit; i++ {
		if number % int(i) == 0 {
			return false
		}
	}
	return true
}

func PrimeNumbers(limit int) int {
	primes := []int{2,}
	n := 2

	for len(primes) < limit {
		if isPrime(n) {
			primes = append(primes, n)
		}
		n++
	}

	return primes[limit-1]
}

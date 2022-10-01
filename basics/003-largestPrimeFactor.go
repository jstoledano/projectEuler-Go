/*
 * ***** PROBLEM 003: LARGEST PRIME FACTOR *****
 *
 * The prime factors of 13195 are 5, 7, 13 and 29.
 * What is the largest prime factor of the number 600851475143 ?
 *
 */

package basics

import "math"

func LargestPrimeFactor(limit uint64) uint64 {
	i := uint64(2)
	n := limit
	for (i*i <= n) {
		if (n % i != 0) {
			i++
		} else {
			n = uint64(math.Floor(float64(n) / float64(i)))
		}
	}
	return n;
}

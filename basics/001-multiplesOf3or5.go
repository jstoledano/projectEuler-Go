/*
 * ***** Problem 1 *****
 * ***** Multiples of 3 or 5 *****
 *
 * If we list all the natural numbers below 10 that are multiples of 3 or 5, we
 * get 3, 5, 6 and 9. The sum of these multiples is 23.
 *
 * Find the sum of all the multiples of 3 or 5 below 1000.
 *
 */

package basics

func byZero(a, b, c uint) bool {
	if a % b == 0 || a % c == 0 {
		return true
	}
	return false
}


func MultiplesOf(a, b, limit uint) uint {
	var sum uint = 0

	for i:= uint(0); i < limit; i++ {
		if byZero(i, a, b) {
			sum += i
		}
	}
	return sum
}

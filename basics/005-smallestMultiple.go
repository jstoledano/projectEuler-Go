/*
 * ***** PROBLEM 005: SMALLEST MULTIPLE *****
 *
 * 2520 is the smallest number that can be divided by each
 * of the numbers from 1 to 10 without any remainder.
 *
 * What is the smallest positive number that is evenly
 * divisible by all of the numbers from 1 to 20?
 *
 */

package basics

func SmallestMultiple(a, b int) int {
	multiple := false
	num := 1

	for (!multiple) {
		sum := 0
		num++
		for i := a; i <= b; i++ {
			sum += num % i
		}
		if sum == 0 {
			multiple = true
		}
	}
	return num
}

/*
 * ***** PROBLEM 004: LARGEST PALINDROME NUMBER *****
 *
 * A palindromic number reads the same both ways. The largest
 * palindrome made from the product of two 2-digit numbers is
 * 9009 = 91 × 99.
 *
 * Find the largest palindrome made from the product of two
 * 3-digit numbers.
 *
 */

package basics

import (
	"strconv"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func PalindromeProduct() uint64 {
	var product uint64 = 0
	var largestProduct uint64 = 0
	var stringProduct, reverseProduct string

	for a := 999; a >= 100; a-- {
		for b := 999; b >= 100; b-- {
			product = uint64(a) * uint64(b)
			stringProduct = strconv.Itoa(int(product))
			reverseProduct = reverse(stringProduct)
			if (stringProduct == reverseProduct && product > largestProduct) {
				largestProduct = product
			}
		}
	}
	return largestProduct
}

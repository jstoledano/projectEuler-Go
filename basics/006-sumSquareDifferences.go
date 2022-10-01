/* ***** SUM SQUARE DIFFERENCE *****
         *
         * The sum of the squares of the first ten natural numbers is,
         *                       1^2+2^2+3^2+...+10^2=385
         * The square of the sum of the first ten natural numbers is,
         *                     (1+2+3+...+10)^2 = 55^2 = 3025
         * Hence the difference between the sum of the squares of the
         * first ten natural numbers and the square of the sum is
         *                         3015-385=2640
         * Find the difference between the sum of the squares of the
         * first one hundred natural numbers and the square of the sum.
         *
         */

package basics

import (
	"fmt"
	"math"
)

func SumSquareDifferences(limit int) int {
	sumOfSquares := 0
	sumOfNumbers := 0
	squareOfSum  := 0
	difference   := 0

	for i := 1; i <= limit; i++ {
		sumOfSquares += int(math.Pow(float64(i), 2))
		sumOfNumbers += i
	}

	squareOfSum = int(math.Pow(float64(sumOfNumbers), 2))
	difference = squareOfSum - sumOfSquares

	fmt.Println("Suma de cuadrados: ", sumOfSquares)
	fmt.Println("Cuadrado de sumas: ", squareOfSum)
	fmt.Println("Diferencia: ", difference)

	return difference
}

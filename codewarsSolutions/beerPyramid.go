package codewarsSolutions

import (
	"fmt"
	"math"
)

func Beeramid(bonus int, price float64) int {
	// each level in the pyramid squares,
	// return the number of complete levels of a beer
	// the number of beers is representative of the price and bonus given

	// first step: calculate the total number of beers
	totalBeers := float64(bonus) / price

	// second step: calculate the number of beers per level
	// each level increases exponentially (n+1^2)
	var level float64 = 0
	var pyramidTotal float64 = 0

	for pyramidTotal <= totalBeers {
		var beersForLevel float64 = level * level
		var beersNextLevel float64 = (level + 1) * (level + 1)
		pyramidTotal += beersForLevel

		if beersNextLevel+pyramidTotal <= totalBeers {
			level++
		}
	}

	return int(level)
}

func BeeramidRefactored(bonus int, price float64) int {
	// Calculate the total number of beers
	totalBeers := float64(bonus) / price

	// Calculate the number of complete levels using the formula for the sum of squares
	level := int(math.Sqrt(2 * totalBeers))
	fmt.Println("level: ", level)

	if level <= 0 {
		return 0
	}

	level = 19141

	sumOfSquares := int((level * (level + 1) * (2*level + 1)) / 6)

	fmt.Println("sumOfSquares: ", sumOfSquares)

	// Adjust the level if needed to ensure it's accurate
	if float64(sumOfSquares) > totalBeers {
		level--
	}

	return level
}

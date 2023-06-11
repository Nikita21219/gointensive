// Package ex01 cool
package ex01

import (
	"math"
)

/*
Launch documentation web-server:
	godoc -http=:6060
Generate documentation:
	godoc -url http://localhost:6060/src/main/minCoins.go > doc.html
*/

// minCoins2 returns slice of int
func minCoins2(val int, coins []int) []int {
	if len(coins) == 0 {
		return []int{}
	}

	row := make([]int, val+1)
	prevs := make([]int, val+1)
	var result []int

	for i := 1; i <= val; i++ {
		row[i] = math.MaxInt
		prevs[i] = -1
	}

	for i := 1; i <= val; i++ {
		for _, bank := range coins {
			diff := i - bank
			if diff >= 0 && row[diff] < row[i] {
				row[i] = row[diff]
				prevs[i] = bank
			}
		}
		row[i]++
	}

	for val > 0 {
		if prevs[val] >= 0 {
			result = append(result, prevs[val])
			val -= prevs[val]
		} else {
			break
		}
	}
	return result
}

/*
minCoins2Optimized also returns a slice of int
In this function, I have optimized the memory allocation.
This time it allocates more memory in advance.
*/
func minCoins2Optimized(val int, coins []int) []int {
	if len(coins) == 0 {
		return []int{}
	}

	row := make([]int, val+1)
	prevs := make([]int, val+1)
	result := make([]int, 0, val)

	for i := 1; i <= val; i++ {
		row[i] = math.MaxInt
		prevs[i] = -1
	}

	for i := 1; i <= val; i++ {
		for _, bank := range coins {
			diff := i - bank
			if diff >= 0 && row[diff] < row[i] {
				row[i] = row[diff]
				prevs[i] = bank
			}
		}
		row[i]++
	}

	for val > 0 {
		if prevs[val] >= 0 {
			result = append(result, prevs[val])
			val -= prevs[val]
		} else {
			break
		}
	}
	return result
}

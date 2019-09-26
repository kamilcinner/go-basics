package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// Test sets defined as slices
	a1 := []int{-1, 3, 8, 2, 9, 5}
	a2 := []int{4, 1, 2, 10, 5, 20}
	aTarget := 24
	fmt.Println("test a:", closestSumPair(a1, a2, aTarget)) // should return {5, 20} or {3, 20}

	b1 := []int{7, 4, 1, 10}
	b2 := []int{4, 5, 8, 7}
	bTarget := 13
	fmt.Println("test b:", closestSumPair(b1, b2, bTarget)) // should return {4, 8}, {7, 7}, {7, 5}, or {10, 4}

	c1 := []int{6, 8, -1, -8, -3}
	c2 := []int{4, -6, 2, 9, -3}
	cTarget := 3
	fmt.Println("test c:", closestSumPair(c1, c2, cTarget)) // should return {-1, 4} or {6, -3}

	d1 := []int{19, 14, 6, 11, -16, 14, -16, -9, 16, 13}
	d2 := []int{13, 9, -15, -2, -18, 16, 17, 2, -11, -7}
	dTarget := -15
	fmt.Println("test d:", closestSumPair(d1, d2, dTarget)) // should return {-16, 2}, {-9, -7}
}

// a1 and a2 are the given arrays, and target is the target sum.
// It should return an array of two numbers as the result,
// one from each array.
func closestSumPair(a1, a2 []int, target int) [2]int {
	a1Sorted := make([]int, len(a1))
	copy(a1Sorted, a1)
	sort.Ints(a1Sorted)
	a2Sorted := make([]int, len(a2))
	copy(a2Sorted, a2)
	sort.Ints(a2Sorted)

	i := 0
	j := len(a2Sorted) - 1
	smallestDiff := iAbs(a1Sorted[0] + a2Sorted[0] - target)
	closestPair := [2]int{a1Sorted[0], a2Sorted[0]}

	for i < len(a1Sorted) && j >= 0 {
		v1 := a1Sorted[i]
		v2 := a2Sorted[j]
		currentDiff := v1 + v2 - target
		if iAbs(currentDiff) < smallestDiff {
			smallestDiff = iAbs(currentDiff)
			closestPair[0] = v1
			closestPair[1] = v2
		}

		if currentDiff == 0 {
			return closestPair
		} else if currentDiff < 0 {
			i++
		} else {
			j--
		}
	}

	return closestPair
}

func iAbs(n int) int {
	return (int)(math.Abs((float64)(n)))
}

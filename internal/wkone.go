package internal

import (
	"math"
)

// Single Number: Given a non-empty array of integers, every element appears twice except for one. Find that single one.
// Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?
func singleNumber(nums []int) int {
	// Handle null case
	if nums == nil {
		panic("Provided array cannot be nil.")
	}

	// Handle empty case
	if len(nums) == 0 {
		panic("Provided array cannot be empty.")
	}

	// Handle length 1 case
	if len(nums) == 1 {
		return nums[0]
	}

	// Create Map with key being the integer and the value being the frequency of the key in the array
	frequencyMap := make(map[int]int)

	// Update frequency count
	for _, value := range nums {
		prev := frequencyMap[value]

		frequencyMap[value] = prev + 1
	}

	// Find index with frequency of 1
	for index, freq := range frequencyMap {
		if freq == 1 {
			return index
		}
	}

	// Reach here if there is no element w/ a frequency of 1
	panic("There is no single element in the provided array")
}

/*
A happy number is a number defined by the following process:
Starting with any positive integer, replace the number by the sum of the squares of its digits,
and repeat the process until the number equals 1 (where it will stay),
or it loops endlessly in a cycle which does not include 1.
Those numbers for which this process ends in 1 are happy numbers.
 */
func isHappy(n int) bool {
	// Zero and negative numbers are not happy.
	if n < 1 {
		return false
	}

	// The only unhappy cycle in base-10 is 4 → 16 → 37 → 58 → 89 → 145 → 42 → 20
	switch n {
		case 4:
		case 16:
		case 37:
		case 58:
		case 89:
		case 145:
		case 42:
		case 20:
			return false
	}

	// Any power of 10 (10, 100, 1000, etc) is a happy number.
	if isPower(10, n) {
		return true
	}

	// Calculate new number to check.
	toCheck := 0
	for n > 0 {
		toCheck += int(math.Pow(float64(n % 10), 2))
		n /= 10
	}

	// Check new number
	return isHappy(toCheck)
}

// Returns true if y is a power of x, false otherwise
func isPower(x, y int) bool {
	// Negatives are impossible
	if y < 0 {
		return false
	}

	// Any number to the zero power is 1
	if y == 1 {
		return true
	}

	xf := float64(x)
	yf := float64(y)
	logyf := math.Log(yf)
	logxf := math.Log(xf)
	logi := int(logyf) / int(logxf)
	logf := logyf / logxf

	return float64(logi) == logf
}

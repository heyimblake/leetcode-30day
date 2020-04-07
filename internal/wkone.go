package internal

import (
	"math"
	"sort"
)

// Single Number:
// Given a non-empty array of integers, every element appears twice except for one. Find that single one.
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
Happy Number:
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
		toCheck += int(math.Pow(float64(n%10), 2))
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

// Maximum Subarray:
// Find the contiguous subarray (containing at least one number) which has the largest sum and return its sum
func maxSubArray(nums []int) int {
	// Null case
	if nums == nil {
		panic("Provided array cannot be nil.")
	}

	// Empty case
	if len(nums) == 0 {
		return 0
	}

	// 1 element case
	if len(nums) == 1 {
		return nums[0]
	}

	currentMax := nums[0]
	maxHere := nums[0]

	for i, v := range nums {
		if i == 0 {
			continue
		}

		// Calculate continuous sum
		sum := maxHere + v

		if v < sum {
			maxHere = sum
		} else {
			maxHere = v
		}

		// Update max if we current continuous sum is higher
		if maxHere > currentMax {
			currentMax = maxHere
		}
	}

	return currentMax
}

// Move Zeroes:
// Move all 0's to the end of it while maintaining the relative order of the non-zero elements.
// 1) You must do this in-place without making a copy of the array.
// 2) Minimize the total number of operations.
func moveZeroes(nums []int) {
	// Null case
	if nums == nil {
		panic("Provided array cannot be nil.")
	}

	// Empty/1 Case
	if len(nums) <= 1 {
		return
	}

	indexElementMap := make(map[int]int)
	zeroCount := 0

	// Calculate new indexes and put into map
	for _, num := range nums {
		if num != 0 {
			indexElementMap[len(indexElementMap)] = num
		} else {
			zeroCount++
		}
	}

	// No zeroes means no modifications needed!
	if zeroCount == 0 {
		return
	}

	// Set new values with their new indexes
	for i := 0; i < len(indexElementMap); i++ {
		nums[i] = indexElementMap[i]
	}

	i := len(nums) - zeroCount

	// Add zeroes to end
	for i < len(nums) {
		nums[i] = 0
		i++
	}
}

// Best Time to Buy and Sell Stock II:
// The ith element is the price of a given stock on day i
// Find the maximum profit
// You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times)
// You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again)
func maxProfit(prices []int) int {
	// Null Case
	if prices == nil {
		panic("Provided array cannot be nil.")
	}

	// Handle Empty/1 Case
	if len(prices) < 2 {
		return 0
	}

	// If the array is sorted (max value first), then max profit is 0
	if sort.SliceIsSorted(prices, func(x, y int) bool { return prices[x] >= prices[y] }) {
		return 0
	}

	// If the array is sorted (min value first), then max profit is buying on day 0 and selling on day n-1
	if sort.SliceIsSorted(prices, func(x, y int) bool { return prices[x] <= prices[y] }) {
		return prices[len(prices)-1] - prices[0]
	}

	canBuy := true
	boughtPrice := 0
	maxProfit := 0

	for i := 0; i < len(prices); i++ {
		pricei := prices[i]

		// Sell if we can
		if !canBuy {
			// Don't sell if we lose or don't make any money
			if pricei <= boughtPrice {
				continue
			}

			// Don't sell if we can get a better deal next time
			if i+1 < len(prices) && prices[i+1] > pricei {
				continue
			}

			canBuy = true
			maxProfit += pricei - boughtPrice
		}

		// See if we should buy (dont buy on last check)
		if canBuy && i < len(prices)-1 {
			// If next price is lower (or same) than current, don't buy
			if prices[i+1] <= pricei {
				continue
			}

			// If we reached here, we should buy
			canBuy = false
			boughtPrice = pricei
		}
	}

	return maxProfit
}

// Group Anagrams:
// Given an array of strings, group anagrams together.
// All inputs will be in lowercase.
// The order of your output does not matter.
func groupAnagrams(strs []string) [][]string {
	// Create map (Key: Sorted string. Value: array of associated strings)
	sortedStringsMap := make(map[string][]string)

	// Add to map
	for _, str := range strs {
		chars := []rune(str)
		sort.SliceStable(chars, func(i, j int) bool { return chars[i] < chars[j] })
		s := string(chars)
		sortedStringsMap[s] = append(sortedStringsMap[s], str)
	}

	// Convert map to double array
	arr := make([][]string, len(sortedStringsMap))
	i := 0

	for _, vals := range sortedStringsMap {
		for _, s := range vals {
			arr[i] = append(arr[i], s)
		}
		i++
	}

	return arr
}

// Counting Elements:
// Given an integer array arr, count element x such that x + 1 is also in arr.
// If there're duplicates in arr, count them seperately.
// Constraints:
// 1 <= arr.length <= 1000
// 0 <= arr[i] <= 1000
func countElements(arr []int) int {
	// Null Case
	if arr == nil {
		panic("Provided array cannot be nil.")
	}

	// Inputs of length 0/1 cannot have a matching element.
	if len(arr) < 2 {
		return 0
	}

	count := 0

	// Create map. (Key: element. Value: frequency)
	freqMap := make(map[int]int)

	// Sort Array
	sort.SliceStable(arr, func(i, j int) bool { return arr[i] < arr[j] })

	// Put elements and their frequencies into the map
	for _, element := range arr {
		freqMap[element] = freqMap[element] + 1
	}

	// Find matching pairs
	for element, freq := range freqMap {
		for freq > 0 {
			pairAmount := freqMap[element+1]

			if pairAmount > 0 {
				count++
				freq--
			} else {
				break
			}
		}
	}

	return count
}

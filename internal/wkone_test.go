package internal

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	inputs := [][]int{{2, 2, 1}, {4, 1, 2, 1, 2}}
	expected := []int{1, 4}

	for i, input := range inputs {
		actual := singleNumber(input)
		if actual != expected[i] {
			fmt.Printf("Expected %d but got %d.\n", expected[i], actual)
			t.Fail()
		}
	}
}

func TestIsHappy(t *testing.T) {
	expected := []bool{true, false}
	inputs := []int{19, 21}

	for i, b := range expected {
		actual := isHappy(inputs[i])
		if actual != b {
			fmt.Printf("Expected %t but got %t.\n", b, actual)
			t.Fail()
		}
	}
}

func TestIsPower(t *testing.T) {
	expected := []bool{true, true, false}
	actual := []bool{isPower(10, 10), isPower(10, 100), isPower(10, 20)}

	for i, b := range expected {
		if actual[i] != b {
			fmt.Printf("Expected %t but got %t.\n", b, actual[i])
			t.Fail()
		}
	}
}

func TestMaxSubArray(t *testing.T) {
	inputs := [][]int{{-2, -3, 4, -1, -2, 1, 5, -3}, {-2, 1, -3, 4, -1, 2, 1, -5, 4}, {-2, -1}}
	expected := []int{7, 6, -1}

	for i, input := range inputs {
		actual := maxSubArray(input)
		if actual != expected[i] {
			fmt.Printf("Expected %d but got %d.\n", expected[i], actual)
			t.Fail()
		}
	}
}

func TestMoveZeroes(t *testing.T) {
	inputs := [][]int{{0, 1, 0, 3, 12}, {0, 0}, {1, 3, 12}, {0, 0, 0, 1, 3, 12, 0}}
	expected := [][]int{{1, 3, 12, 0, 0}, {0, 0}, {1, 3, 12}, {1, 3, 12, 0, 0, 0, 0}}

	for i, input := range inputs {
		moveZeroes(input)
		if !reflect.DeepEqual(input, expected[i]) {
			fmt.Printf("Expected %d but got %d.\n", expected[i], input)
			t.Fail()
		}
	}
}

func TestMaxProfit(t *testing.T) {
	inputs := [][]int{{5, 2, 10}, {7, 1, 5, 3, 6, 4}, {1, 2, 3, 4, 5}, {7, 6, 4, 3, 1}, {5, 9, 3, 3, 6, 1}, {9, 2, 1, 5, 7}, {1, 7, 2, 3, 6, 7, 6, 7}, {1, 2, 3}, {8, 6, 4, 3, 3, 2, 3, 5, 8, 3, 8, 2, 6}}
	expected := []int{8, 7, 4, 0, 7, 6, 12, 2, 15}

	for i, input := range inputs {
		actual := maxProfit(input)
		if actual != expected[i] {
			fmt.Printf("Expected %d but got %d.\n", expected[i], actual)
			t.Fail()
		}
	}
}

func TestGroupAnagrams(t *testing.T) {
	inputs := [][]string{{"eat", "tea", "tan", "ate", "nat", "bat"}}
	expected := [][][]string{{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}}

	for i, input := range inputs {
		actual := groupAnagrams(input)
		if !reflect.DeepEqual(actual, expected[i]) {
			fmt.Printf("Expected %s but got %s.\n", expected[i], actual)
			t.Fail()
		}
	}
}

func TestCountElements(t *testing.T) {
	inputs := [][]int{{1, 2, 3}, {1, 1, 3, 3, 5, 5, 7, 7}, {1, 3, 2, 3, 5, 0}, {1, 1, 2, 2}, {8, 9, 10, 13, 18, 14, 14, 14, 15, 17, 19}, {21, 19, 18, 17}}
	expected := []int{2, 0, 3, 2, 8, 2}

	for i, input := range inputs {
		actual := countElements(input)
		if actual != expected[i] {
			fmt.Printf("Expected %d but got %d.\n", expected[i], actual)
			t.Fail()
		}
	}
}

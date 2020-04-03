package internal

import (
	"fmt"
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

package main

import (
	"reflect"
	"testing"
)

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

// version 1
// func SumAll(nums1 []int, nums2 []int) []int {
// 	nums1_sum := Sum(nums1)
// 	nums2_sum := Sum(nums2)
// 	return []int{nums1_sum, nums2_sum}
// }

// version 2
// func SumAll(numbersToSum ...[]int) []int {
// 	lengthOfNumbers := len(numbersToSum)
// 	sums := make([]int, lengthOfNumbers)

// 	for i, numbers := range numbersToSum {
// 		sums[i] = Sum(numbers)
// 	}

// 	return sums
// }

// version 3
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sumsTail []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sumsTail = append(sumsTail, 0)
		} else {
			sumsTail = append(sumsTail, Sum(numbers[1:]))
		}
	}
	return sumsTail
}

func TestSumAllTails(t *testing.T) {

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}

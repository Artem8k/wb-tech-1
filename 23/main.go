package main

import "fmt"

func deleteFromSlice(nums []int, i int) []int {
	var result []int

	result = append(result, nums[:i]...)
	result = append(result, nums[i+1:]...)

	return result
}

func main() {
	nums := []int{1, 5, 6, 2, 7, 14, 3, 9, 11, 19, 13, 15, 16, 18, 17, 4, 8, 10, 12, 20}
	fmt.Println(deleteFromSlice(nums, 3))
}

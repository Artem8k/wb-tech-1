package main

import (
	"fmt"
	"sort"
)

func binarySearch(nums []int, num int) int {
	low := 0
	high := len(nums) - 1

	// нахожу середину слайса, если искомое число больше числа по середине слайса то, делю слайс пополам и ищу в нужом из них
	// если число нашлось то возвращаю иднекс числа, если нет то -1
	for low <= high {
		middle := int(low + (high-low)/2)

		if nums[middle] == num {
			return middle
		} else if nums[middle] < num {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}

	return -1
}

func main() {

	nums := []int{1, 5, 6, 2, 7, 14, 3, 9, 11, 19, 13, 15, 16, 18, 17, 4, 8, 10, 12, 20}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	fmt.Println(binarySearch(nums, 15))
}

package main

import "fmt"

func quickSort(nums []int) []int {

	if len(nums) < 2 {
		return nums
	}
	// устанавлаю первый элемент (опору)
	firstElem := nums[0]

	lessThanFirst := []int{}
	greaterThanFirst := []int{}
	equalToFirstElem := []int{}

	// разбиваю на три слайса, слайс со значениями больше опоры, меньше опоры и равными опоре
	for _, n := range nums {

		if firstElem > n {
			lessThanFirst = append(lessThanFirst, n)
		} else if firstElem < n {
			greaterThanFirst = append(greaterThanFirst, n)
		} else {
			equalToFirstElem = append(equalToFirstElem, n)
		}

	}

	sorted := []int{}
	// делаю рекурсию для слайсов длина которых больше 0
	if len(lessThanFirst) > 0 {
		sorted = append(sorted, quickSort(lessThanFirst)...)
	}

	sorted = append(sorted, equalToFirstElem...)

	if len(greaterThanFirst) > 0 {
		sorted = append(sorted, quickSort(greaterThanFirst)...)
	}

	return sorted
}

func main() {
	nums := []int{1, 5, 6, 2, 7, 14, 3, 9, 11, 19, 13, 15, 16, 18, 17, 4, 8, 10, 12}

	sorted := quickSort(nums)

	fmt.Println(sorted)
}

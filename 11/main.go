package main

import "fmt"

//Если предположить что длина слайсов nums1 и nums2 одинаковая, то сложность O(n)
func findIntersection1(nums1 []int, nums2 []int) {
	var numsMap = make(map[int]int)
	var intersection []int
	// Добавляю все числа в map, использую key: число, значение: количество чисел в последовательности
	// использую констуркцию if numsMap[num] > 0, для того чтобы не задублировать количество чисел,
	// если вдруг в последовательности окажется несколько одинаковых элементов
	for _, num := range nums1 {
		if numsMap[num] > 0 {
			continue
		} else {
			numsMap[num] = 1
		}
	}

	// Тут также использую key: число, значение: количество чисел в последовательности
	// использую ту же конструкцию if numsMap[num] > 0
	// но уже для того чтобы найти пересечения
	for _, num := range nums2 {
		if numsMap[num] > 0 {
			numsMap[num] += 1
		} else {
			continue
		}
	}

	// из готовой Map'ы собираю слайс с числами попавшими в пересечение
	for num, count := range numsMap {
		if count > 1 {
			intersection = append(intersection, num)
		}
	}
	fmt.Println(intersection)
}

// O(n)^2
func findIntersection2(nums1 []int, nums2 []int) {
	var intersection []int
	var numsMap = make(map[int]int)

	// Прохожусь большим слайсом по меньшему с помощью циклов, нахожу пересечения добавляю в мапу
	if len(nums1) >= len(nums2) {
		for _, num1 := range nums1 {
			for _, num2 := range nums2 {
				if num1 == num2 {
					numsMap[num1] += 1
				}
			}
		}
	} else {
		for _, num2 := range nums2 {
			for _, num1 := range nums1 {
				if num2 == num1 {
					numsMap[num2] += 1
				}
			}
		}
	}

	// из готовой Map'ы собираю слайс с числами попавшими в пересечение
	for num, count := range numsMap {
		if count >= 1 {
			intersection = append(intersection, num)
		}
	}

	fmt.Println(intersection)
}

func main() {
	nums1 := []int{1, 6, 6, 3, 5, 8, 9, 9}

	nums2 := []int{3, 8, 2, 9, 1, 9}

	findIntersection1(nums1, nums2)
	findIntersection2(nums1, nums2)
}

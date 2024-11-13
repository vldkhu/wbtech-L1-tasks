package main

import "fmt"

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.


*/

func quickSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	pivot := a[len(a)/2] // выбор опорного элемента, в данном случае середина
	left := []int{}
	right := []int{}

	for _, v := range a {
		if v < pivot {
			left = append(left, v)
			// Элементы меньше опорного
		} else if v > pivot {
			right = append(right, v)
			// Элементы больше опорного
		}
	}
	return append(append(quickSort(left), pivot), quickSort(right)...)
}

func main() {
	unsortSlice := []int{34, 256, 234, 6, 34, 376, 787, 46, 12, 3, 66, 55, 64, 234, 878, 86, 7458, 745, 75, 93, 76}
	sortSlice := quickSort(unsortSlice)
	fmt.Println(sortSlice)
}

package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.

Первый варинт добавление в слайс через for
Временная сложность: O(n * m)
*/

func main() {

	sliceA := []int{1, 2, 3, 4, 5}
	sliceB := []int{4, 5, 6, 7, 8}

	fmt.Println(intersection(sliceA, sliceB))
}

func intersection(s1, s2 []int) []int {
	var slice []int
	for _, v := range s1 {
		for _, v2 := range s2 {
			if v == v2 {
				slice = append(slice, v)
			}
		}
	}
	return slice
}

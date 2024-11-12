package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.

Второй варинт добавление в map (более производительный вариант)
Временная сложность: O(n + m)
*/

func main() {
	// создаем два слайса
	sliceA := []int{1, 2, 3, 4, 5}
	sliceB := []int{4, 5, 6, 7, 8}

	fmt.Println(intersection(sliceA, sliceB))
}

func intersection(s1, s2 []int) []int {
	// создаем map для хранения элементов первого среза
	elemMap := make(map[int]struct{})
	for _, v := range s1 {
		elemMap[v] = struct{}{}
	}

	var slice []int
	// Проверяем элементы второго среза
	for _, v := range s2 {
		if _, exists := elemMap[v]; exists {
			slice = append(slice, v)
		}
	}
	return slice
}

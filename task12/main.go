package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
собственное множество.
*/

func main() {
	animals := [...]string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем множество с помощью карты
	uniqueAnimals := make(map[string]struct{})
	for _, animal := range animals {
		uniqueAnimals[animal] = struct{}{} // Используем пустую структуру для экономии памяти
	}

	// Выводим уникальные элементы
	for uniqueAnimal := range uniqueAnimals {
		fmt.Println(uniqueAnimal)
	}
}

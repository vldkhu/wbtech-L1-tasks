package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow»

2 способ: через билдер
*/

func reverseWords(s string) string {
	// Разделяем строку на слова
	words := strings.Fields(s)

	// Переворачиваем порядок слов
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Объединяем слова обратно в строку
	return strings.Join(words, " ")
}

func main() {
	str, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return
	}
	reverse := reverseWords(str)
	fmt.Println(reverse)
}

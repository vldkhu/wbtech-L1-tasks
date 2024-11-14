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

1 способ: через конкатеннацию строк
*/

func main() {
	str, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return
	}
	str = str[:len(str)-2]
	split := strings.Split(str, " ")
	res := ""

	for i := len(split) - 1; i >= 0; i-- {
		res += split[i]
		if i < len(split) {
			res += " "
		}
	}
	fmt.Println(res)
}

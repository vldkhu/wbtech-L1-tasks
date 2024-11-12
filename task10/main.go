package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
градусов. Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

func main() {
	array := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	val := make(map[int][]float64)

	for _, v := range array {
		// Определяем группу с шагом 10 градусов
		group := int(v/10) * 10
		val[group] = append(val[group], v)
	}

	// Выводим результат
	var builder strings.Builder
	for k, v := range val {
		builder.WriteString(strconv.Itoa(k))
		builder.WriteString(":{")
		for j := 0; j < len(v); j++ {
			builder.WriteString(fmt.Sprintf("%.1f", v[j]))
			if j+1 != len(v) {
				builder.WriteString(", ")
			}
		}
		builder.WriteString("}, ")
	}
	res := builder.String()
	res = res[0 : len(res)-2]
	fmt.Print(res)
}

package main

import (
	"fmt"
)

/*
Разработать программу, которая в рантайме способна определить тип
переменной: int, string, bool, channel из переменной типа interface{}.

SWITCH конструкция с ключевым словом type
*/

func typesInfo(values interface{}) {

	switch v := values.(type) {
	case int:
		fmt.Printf("type = \"%T\" and value = %v\n", v, v)
	case string:
		fmt.Printf("type = \"%T\" and value = %v\n", v, v)
	case bool:
		fmt.Printf("type = \"%T\" and value = %v\n", v, v)
	case chan int:
		fmt.Printf("type = \"%T\" and value = %v\n", v, v)
	default:
		fmt.Printf("unknows type... %T!\n", v)
	}
}

func main() {
	typesInfo(52)
	typesInfo("писятдва")
	typesInfo(true)
	typesInfo(52.52)
}

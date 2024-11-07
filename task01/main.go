package main

import (
	"errors"
	"fmt"
)

const adultAgeStep = 18

type Human struct {
	name   string
	age    int
	height int
}

type Action struct {
	*Human
}

func newHuman(name string, age int, height int) (*Human, error) {
	if name == "" {
		return nil, errors.New("имя не может быть пустым")
	}
	if age <= 0 {
		return nil, errors.New("age не может быть отрицательным или ноль")
	}
	if height <= 0 {
		return nil, errors.New("height не может быть отрицальным или ноль")
	}

	return &Human{
		name:   name,
		age:    age,
		height: height,
	}, nil

}

func (h Human) isAdult() bool {
	return h.age > adultAgeStep
}

func main() {
	human1, err := newHuman("Petr", 30, 190)
	if err != nil {
		return
	}
	human2, err := newHuman("Vlad", 17, 170)
	if err != nil {
		return
	}

	action1 := Action{Human: human1}
	action2 := Action{Human: human2}

	fmt.Println(action1.isAdult()) // true
	fmt.Println(action2.isAdult()) // false

}

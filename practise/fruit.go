package main

import "fmt"

type fruit []string

func newFruit() fruit {
	fruits := fruit{}

	fruitTypes := []string{"Apple", "Banana", "Grapes", "Mango", "Kiwi"}
	fruitValues := []string{"One", "Two", "Three", "Four", "Five"}

	for _, fruitType := range fruitTypes {
		for _, value := range fruitValues {
			fruits = append(fruits, value+" of "+fruitType)
		}
	}

	return fruits
}

func (d fruit) Print() {
	for i, fruit := range d {
		fmt.Println(i, fruit)
	}
}
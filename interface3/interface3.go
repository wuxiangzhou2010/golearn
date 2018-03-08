package main

import (
	"fmt"
	// "math"
)

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow"
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "????"
}

type JavaProgmamer struct {
}

func (i JavaProgmamer) Speak() string {
	return "Design patterns"
}

func main() {
	animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgmamer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}

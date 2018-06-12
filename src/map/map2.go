package main

import (
	"fmt"
)

func main() {

	// var colors map[string]string
	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#234343",
		"white": "#ffffff",
	}
	// colors["white"] = "#ffffff"
	// fmt.Println(colors)
	// delete(colors, "white")
	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex code for", color, "is", hex)
	}
}

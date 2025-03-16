package main

import "fmt"

func main() {
	// var colors map[string]string
	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	colors["white"] = "#ffffff"
	fmt.Println(colors)
	fmt.Println(colors["red"])
	delete(colors, "red")
	fmt.Println(colors)
	printMap(colors)
}

func printMap(colors map[string]string) {
	for k, v := range colors {
		fmt.Println(k, v)
	}
}

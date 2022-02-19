package main

import "fmt"

func main() {

	arr := []string{"Marcos", "Juan", "Alice", "Bob"}

	for _, name := range arr {
		switch name {
		case "Marcos":
			fmt.Println("Hello, Marcos")
		case "Alice":
			fmt.Println("Hello, Alice")
		default:
			fmt.Println("Hello, stranger")
		}
	}
}

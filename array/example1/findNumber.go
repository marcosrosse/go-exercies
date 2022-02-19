package main

import "fmt"

func arr(arr int) []int {
	array := []int{2, 3, 4, 3, 2, 1}
	return array
}
func main() {

	j := 0

	for _, i := range arr(3) {
		if i == 1 {
			j++
		}

	}
	fmt.Printf("Encontrado o numero 2 %d vezes\n", j)
}

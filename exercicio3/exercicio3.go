package main

import "fmt"

func main() {
	vetor := []int{1, 2, 3, 4, 5}

	var i, aux int
	aux = 1

	for i = 1; i < 5; i++ {
		aux = aux * (i + 1)
		vetor[i] = int(aux / 2)
	}
	fmt.Println(vetor[3])
}

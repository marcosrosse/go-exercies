package main

import "fmt"

func main() {

	vetor := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var i, j int

	var temValoresRepetidos bool

	vetor = append(vetor, 1)
	vetor = append(vetor, 2)
	vetor = append(vetor, 3)
	vetor = append(vetor, 4)
	vetor = append(vetor, 5)
	vetor = append(vetor, 6)
	vetor = append(vetor, 7)
	vetor = append(vetor, 8)
	vetor = append(vetor, 9)
	vetor = append(vetor, 10)

	// temValoresRepetidos = false

	for i = 1; i < 10; i++ {
		for j = 10; j > 1; j-- {
			if vetor[i] == vetor[j] {
				temValoresRepetidos = true
				fmt.Println("Valor repetido indice I: ", vetor[i])
				fmt.Println("Valor repetido indice J: ", vetor[j])
			}
		}
	}
	fmt.Println(temValoresRepetidos)

}

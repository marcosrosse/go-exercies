package main

import "fmt"

var resultado int

func funcao(x, y int) int {
	if y == 0 {
		return x
	} else {
		return funcao(y, x%y)
	}
}

func main() {

	resultado = funcao(9, 3)
	fmt.Println(resultado)
}

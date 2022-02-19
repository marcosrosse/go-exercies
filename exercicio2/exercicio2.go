package main

import "fmt"

func f(n int) int {
	if n < 2 {
		return n
	} else {

		return f(n-1) + f(n-2)
	}
}

func main() {
	fmt.Println(f(5))
}

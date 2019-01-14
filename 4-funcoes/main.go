package main

import "fmt"

// exemplo de função privada
func add(x int, y int) (int, int) {
	return x + y, 0
}

func main() {
	result, other := add(10,20)
	fmt.Println("O resultado de 10 + 20 é", result)
	fmt.Println("O parametro extra é: ", other)
}

package main

import (
	"fmt"
)

//exercicio criar slice com 5 valores e mostrar elas pelo range, com tipo tb

func main() {

	slice1 := []int{1, 2, 3, 4, 5}

	for _, valor := range slice1 {
		fmt.Printf("%v,%T", valor, valor)
	}
}

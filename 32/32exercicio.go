package main

import (
	"fmt"
)

// basicamente mostrar todos os anos desde o proprio nascimento usando FOR single condition inves de 
// for com inicialização;condição;ação
func main() {
	x := 2002
	for x <= 2024 {
		fmt.Printf("%v\n", x)
		x++
	}
}

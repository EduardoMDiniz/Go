package main

import (
	"fmt"
)

func main() {

	pizzas := map[string]int{
		"Calabresa":      56,
		"Pepperoni":      50,
		"Quatro Queijos": 45,
	}
	fmt.Println(pizzas)

	for key, value := range pizzas { // como sempre, as 2 variaveis de iniciação
		fmt.Println(key, value) // de for, recebem o indice e valor
	}

	//pra deletar algo do map simples

	delete(pizzas, "Pepperoni")
	
	fmt.Println(pizzas)
}

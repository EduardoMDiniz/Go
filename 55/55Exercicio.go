package main    

import (
	"fmt"
)

//Mesmo 54 mas add uma entrada a mais posterior de uma key:value

func main (){

	pessoas := map [string] string {

		"Sobrenome1": "Carros",
		"Sobrenome2": "Jogos",
		"Sobrenome3": "Futebol",
	}

	pessoas ["Sobrenome4"] = "Basket"
	
	for indice,valor := range pessoas{
		fmt.Println(indice,valor)
	}
}
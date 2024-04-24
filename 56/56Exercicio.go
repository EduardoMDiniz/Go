package main    

import (
	"fmt"
)

//Mesmo 54/55 mas add uma entrada a mais posterior de uma key:value

func main (){

	pessoas := map [string] string {

		"Sobrenome1": "Carros",
		"Sobrenome2": "Jogos",
		"Sobrenome3": "Futebol",
	}

	delete(pessoas,"Sobrenome1")

	for indice,valor := range pessoas{
		fmt.Println(indice,valor)
	}
}
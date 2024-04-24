package main     

import (
	"fmt"
)

/*
- Crie um map com key tipo string e value tipo []string.
    - Key deve conter nomes no formato sobrenome_nome
    - Value deve conter os hobbies favoritos da pessoa
- Demonstre todos esses valores e seus índices.
- Solução: https://play.golang.org/p/nD3TW8VQmH
*/

func main (){

	pessoas := map [string] string {

		"Sobrenome1": "Carros",
		"Sobrenome2": "Jogos",
		"Sobrenome3": "Futebol",
	}
	for indice,valor := range pessoas{
		fmt.Println(indice,valor)
	}
}
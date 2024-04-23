package main

import (
	"fmt"
)

/*
- Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
    - "Nome", "Sobrenome", "Hobby favorito"
- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
- Solução: https://play.golang.org/p/Gh81-d5tMi

*/

func main() {

	slice1 := [][]string{
		{"Nome","Sobrenome","Hobby"},
		{"Eduardo", "Surname", "Pelúcia"},
		{"Marcelo", "Sobrenome", "Lol"},
		{"Breno", "Cognome", "Valorant"},
	}

	for _, valor := range slice1 {
		fmt.Println(valor[0], valor[1], valor[2])

	}
}

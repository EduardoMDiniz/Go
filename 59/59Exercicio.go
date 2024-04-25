package main

import (
	"fmt"
)

/*
- Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
    - Nome
    - Sobrenome
    - Sabores favoritos de sorvete
- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
- Solução: https://play.golang.org/p/Pyp7vmTJfY  */

type saboresfav struct {
	saboresfavoritos []string //criei um tipo pra struct com atributo de slice
}

type pessoa struct {
	nome      string
	sobrenome string
	saboresfav
}

func main() {

	pessoa1 := pessoa{
		nome:      "Maria",
		sobrenome: "Silva",
		saboresfav: saboresfav{
			saboresfavoritos: []string{"Chocolate", "Banana"}, // como passar um atributo de slice
		},
	}

	pessoa2 := pessoa{

		nome:      "joão",
		sobrenome: "Moraes",
		saboresfav: saboresfav{
			saboresfavoritos: []string{"Creme", "Flocos"},
		},
	}

	for _, valor := range pessoa1.saboresfavoritos {
		fmt.Println(valor)

	}
	for _, valor := range pessoa2.saboresfavoritos {
		fmt.Println(valor)

	}

}

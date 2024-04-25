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

	meumapa := make(map[string]pessoa) //Criando o map usando make

	meumapa["Moraes"] = pessoa{
		nome:      "joão",
		sobrenome: "Moraes",
		saboresfav: saboresfav{
			saboresfavoritos: []string{"Creme", "Flocos"},
		},
	}
	meumapa["Garibaldi"] = pessoa{
		nome: "Anita",
		sobrenome: "Garibaldi",
		saboresfav: saboresfav{ 
			saboresfavoritos: []string{"Menta","Amendoas"},
		},
	}

	for _, valor := range meumapa{
		fmt.Println(valor)
	}
}

/*
- Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
- Demonstre os valores do map utilizando range.
- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
- Solução: https://play.golang.org/p/GLK11Q1_x8y */
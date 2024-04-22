package main    

import(
	"fmt"
)

//maps usam formato de chave valor map[key]value { key: value }
//para guardar varias coisas dentro de uma variavel de forma aleatoria

func main (){

	amigos := map[string] int { //estrutura básica, variavel := map [tipo da chave] tipo do valor ai abre e coloca{}
		"Lugano": 5,
		"raí": 10,
		"lucas moura": 7,
	}

	fmt.Printf("%v\n",amigos) 

	amigos ["Ceni"] = 01 // como adicionar uma chave valor ao map

	fmt.Println(amigos) 
}
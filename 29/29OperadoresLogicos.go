package main

import(
	"fmt"
)

// os operadores lógicos são:       &&    ||   !    
// é uma obviedade mas a gente usa pra colocar condições compostas que podem ou não concordar entre si

// se eu coloco condição 1 && condição 2 pra comparação lógica ser true eu preciso que os 2 sejam true

// se eu coloco condição 1 || condição 2 eu preciso que apenas 1 deles seja true

//exemplos

func main (){
	x := 10
	if x > 5 && x < 15{
		fmt.Printf("True po")
	}
}

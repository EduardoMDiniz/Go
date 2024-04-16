package main    

import(
	"fmt"
)
// for de statement

func main (){
	for x := 65; x <= 90; x++{
		fmt.Printf("%v",x)
		for y := 1; y <= 3; y++{   // aqui se usa o nested loop pra iterar outra variaval a cada vez q
			fmt.Printf("\t%#U",y)  // a primeira é iterada
		}
	}
}


// Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.

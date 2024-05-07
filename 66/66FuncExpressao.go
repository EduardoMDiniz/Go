package main  

import (
	"fmt"
)

//basicamente atribuir a uma variavel uma função

func main (){


	x := 100

	y := func (x int) int {         //int aqui pra tipar o tipo do nnosso retorno da func             
		fmt.Println(x,"vezes 30 é:")
		fmt.Println(x*30)
		return x
} 
	fmt.Println( y(x) )  //aqui eu passo a var Y com arg X 
}

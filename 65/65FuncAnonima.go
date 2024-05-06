package main  

import (
	"fmt"
)

//Funções descartáveis

func main (){

	x := 100

	func (x int) {                     //func q n passa o nome dela, so os argumentos
		fmt.Println(x,"vezes 30 é:")
		fmt.Println(x*30)
	} (x)                               //no final passamos a o argumennto que vamos usar na func
  
}
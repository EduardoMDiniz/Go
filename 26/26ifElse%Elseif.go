package main    

import(
	"fmt"
)

// O ELSE IF é o mais provavel de esquecer ent vou declarar: Ele é o que  se coloca entre o primeiro
// que sempre vai ser IF e o ultimo que sempre vai ser Else

//aqui vou demonstrar duas func com If e Else, uma onde o if "ganha"
// outra onde o Else "ganha", e outra onde o Else If "ganha"
var a int = 100
func main () {
	if a > 99{                      
		fmt.Println("A é maior que 99")
		sub ()
		sub2 ()
	}
}

func sub () {
	if b := 98; b > 99{             //tb posso declarar a e atribuir variavel com marmota na func
		fmt.Println("A é maior que 99")

	}else{
		fmt.Println("b é menor que 99")
	}
}

func sub2 (){
	if c := 95; c == 100{
		fmt.Println("C igual 100")
	}else if c == 95 {
		fmt.Println("C igual 150")
	}else{
		fmt.Println("C é qualquer outra coisa")
	}
}
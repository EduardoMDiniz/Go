package main     

import (
	"fmt"
)

//Recursividade é chamar uma função dentro dela mesma, isso é facilmennte aplicada em fatorial

func main(){
	
	fmt.Println(fatorial(5))
}

func fatorial (x int) int {

	if x == 1 {
		return x 
	}
	return x * fatorial(x-1)        // eu chamei a funcao que vai multiplicar com ela mesmo num valor inferior em 1 (x-1)
}
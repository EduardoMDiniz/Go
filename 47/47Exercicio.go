package main   

import (
	"fmt"
)

//exercicio criar uma slice com numero definido de valores e mostrar eles pelo range

func main (){
	
	slice1 := make([]int ,5 ,10)
	slice1 [0], slice1 [1], slice1[2], slice1[3], slice1[4] = 1,2,3,4,5

	for _, valor := range slice1 {
		fmt.Printf("%v, %T\n", valor, valor)
	}
}


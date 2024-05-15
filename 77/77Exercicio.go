package main 

import (
	"fmt"
)

func main (){


	z:= retorna_Funcao()
	
	fmt.Println(z(10,13))

}

func retorna_Funcao () func(int, int) int {

	return func (a int, b int) int {
		return a + b
	}
}
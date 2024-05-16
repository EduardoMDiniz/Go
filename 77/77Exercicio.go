package main 

import (
	"fmt"
)

/* - Crie uma função que retorna uma função.
- Atribua a função retornada a uma variável.
- Chame a função retornada. */

func main (){


	z:= retorna_Funcao()
	
	fmt.Println(z(10,13))

}

func retorna_Funcao () func(int, int) int {

	return func (a int, b int) int {
		return a + b
	}
}
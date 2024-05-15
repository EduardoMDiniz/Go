package main

import(
	"fmt"
)

//atribuir função a variavel, passar a função como uma expressão

func main (){
	y:= 10
	z:= 10
	var x = func (a int, b int) int{
		return a + b
	}
	fmt.Println(x(y,z))
}

//basicamennnte lembrar quando for passar o var x = func (x int) int{}
//passar os receivers e o retorno, sem nome etc
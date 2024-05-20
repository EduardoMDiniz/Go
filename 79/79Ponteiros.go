package main  

import(
	"fmt"
)

/*PONTEIROS e Dereference

Ponteiro é uma variável que aloca um endereço na memória
Dereference é a forma de acessar o valor que esta naquele endereço da memória */

func main (){

	x := 10

	y := &x

	fmt.Println(*y) //essa *y é a forma que podemos acessar o valor daquele ponteiro
					//O tipo da variável de ponteiro é *subtipo ou seja nesse caso, Ponteiro INT

	//eu posso fazer operações com o ponteiro se eu usar DEREFERENCE, então se eu iterar 

	*y++

	fmt.Println(y) //vai funcionar

}
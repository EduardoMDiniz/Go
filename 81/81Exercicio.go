package main 

import(
	"fmt"
)

/*- Crie um valor e atribua-o a uma variável.
- Demonstre o endereço deste valor na memória.
- Solução: https://play.golang.org/p/0jVt1yaoFL */

func main (){

	x:= 100

	fmt.Println(&x) //saída 0xc00000a0b8 endereço de x na memória

}
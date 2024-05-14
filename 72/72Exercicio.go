package main 

import(
	"fmt"
)

/* - Utilize a declaração defer de maneira que demonstre que sua execução só ocorre ao final do contexto ao qual ela pertence.
- Solução: https://play.golang.org/p/b5Ua2kNN8a */


func main (){

	defer fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
}
package main 

import(
	"fmt"
)

const x int = 10 //detalhe, se uma const n tiver seu tipo especificado, ela so ganha seu tipo	
            	// quando ela é usada, diferente de variavel que ganha quando é atribuida seu valor


const(                 //essa é outra forma de declarar constantes, posso abrir um () e colocar todas
	a int = 10         // que vai fazer parte do projeto
	b float64 = 10.1
	c string = "carro"
)
func main(){
	fmt.Println(x,a,b,c)
}
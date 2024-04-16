package main

import(
	"fmt"
)

type sorvete int 

var x sorvete 

var y int 

func main(){
	fmt.Printf("%v\n",x)
	fmt.Printf("%T\n",x)
	x = 42
	fmt.Printf("%v\n",x)

	y = int (x)          //"Convertendo" que na vdd é atribuindo o valor de X sorvete a variável Y int
	fmt.Printf("%v",y)   // A variável X continua sorvete, so que o valor dela agora vai pra Y
	fmt.Printf("%T",y)   // no tipo que a gente quer (INT)
	

}

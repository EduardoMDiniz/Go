package main     

import (
	"fmt"
)

func main (){

	for x := 10; x <= 100; x++{
		d := x % 4
		fmt.Printf("O resto de %v dividido por 4 é %v\n",x , d)
	}
}

//basicamente mostrar o resultado do resto da divisão por 4 de todos os numeros de 10 a 100
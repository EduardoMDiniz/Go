package main

import (
	"fmt"
)

const (       //basicamente o iota enumera em sequencia de numeros reais e inteiros constantes n
	a = iota  // declaradas
	b = iota 
	c = iota 
	d = iota 
)

func main() {
	fmt.Println(a, b, c, d)

}
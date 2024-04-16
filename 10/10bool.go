package main

import (
	"fmt"
)

var x bool

func main() {
	fmt.Printf("%v", x) //Aqui como n tem valor atribuido a ele, ele vai mostrar o valor 0 bool que é FALSE
	
	x = true
	fmt.Printf("%v", x)

	x = (10 < 0) // operador lógico > o que sempre gera valores booleanos
	fmt.Printf("%v", x)
	
	y := (1000 > 0)
	z := (23 > 0)
	fmt.Printf("%v", y)
	fmt.Printf("%v", z)

}

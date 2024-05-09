package main

import (
	"fmt"
)

/*
- Crie uma função que retorne um int
    - Crie outra função que retorne um int e uma string
    - Chame as duas funções
    - Demonstre seus resultados */

func main () {
	fmt.Println(retornar_int())
	fmt.Println(retornar_int_string ())
}

func retornar_int () int {
	var x int = 2
	var vezes2 int = x * 2

	return vezes2
}

func retornar_int_string () (int,string) {
	x := 10
	y := "casa"
	return x,y
}
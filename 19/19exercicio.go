package main

import (
	"fmt"
)

var a int = 10

func main() {
	fmt.Printf("%v - %b - %#x\n", a, a, a)

	b := a << 1

	fmt.Printf("%v - %b - %#x", b, b, b)

}

// aqui o exercicio era criar uma var int e mostrar decimal dela, binario e hexadecimal e dps
//deslocar 1 bit pra esquerda e alocar em outra variavel e mostrar ela da mesma forma q a outra
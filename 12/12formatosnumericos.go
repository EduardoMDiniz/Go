package main

import (
	"fmt"
)

var a int = 100

func main() {
	fmt.Printf("%d\n %b\n %#x", a, a, a)
}

//aqui basicamente criamos uma var int chamada A e pedimos pro Printf retornar ela em
// Decimal %d , binário %b e hexadecimal %#x


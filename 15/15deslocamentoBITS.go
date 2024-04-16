package main

import (
	"fmt"
)

func main() {
	x := 45
	y := x << 1                   // << e >>  é o operador de deslocamento de bits
	fmt.Printf("%v - %b\n", x, x) // o output vai entregar 101101 que é o binário de 45
	fmt.Printf("%v - %b\n", y, y) // e o Y vira 90 pq a gente moveu e o binarie vira 1011010
}

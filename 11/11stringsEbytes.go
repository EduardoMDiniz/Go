package main

import (
	"fmt"
)

var sb []byte // criei uma variável com valor de bytes, valor zero inclusive sem nd

func main() {
	s := "SAO PAULO FUTEBOL CLUBE" // S automaticamente é string
	sb = []byte(s)                 //nessa linha joguei o valor de S pra dentro de SB que ta em bytes

	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T", sb, sb)
}

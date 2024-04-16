package main

import (
	"fmt"
)

// O continue e o Break necessitam de uma estruturação condicional, então vou usar o IF e o ELSE

func main() {

	for x := 0; x < 10; x++ {
		if x == 5 {
			break
		}
		fmt.Println(x)
	}
}

// o contiue é a mesma coisa, só muda a ação mesmo

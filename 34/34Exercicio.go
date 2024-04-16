package main

import (
	"fmt"
)

var p int = 10

func main() {
	
	switch {
	case p == 10:
		fmt.Printf("%v TRUE 1", p)
	case p == 20:
		fmt.Printf("%v TRUE 2", p)
	case p == 11:
		fmt.Printf("%v TRUE 3", p)
	}
}

// Basicamente fazer um switch sem statement, ou seja, um switch que joga um True ali pra comparar


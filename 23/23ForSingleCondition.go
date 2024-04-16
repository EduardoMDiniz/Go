package main

import (
	"fmt"
)

var a int = 0
var b int = 1

func main() {

	for a < b {        // FOR e condição e dps ação
		a++
		fmt.Println(a)
	}
}

// FOR singles condition
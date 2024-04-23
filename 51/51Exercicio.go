package main

import (
	"fmt"
)

/*
- Comece com essa slice:
    - x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
- Utilizando slicing e append, crie uma slice y que contenha os valores:
    - [42, 43, 44, 48, 49, 50, 51]
*/

func main() {

	slice1 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	slice2 := slice1 [:3]
	slice3 := slice1 [6:]
	slicefinal := append(slice2, slice3...)

	fmt.Println(slicefinal)
}

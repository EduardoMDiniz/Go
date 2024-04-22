package main    

import(
	"fmt"
)

//mostrar fatias de uma slice

func main() {

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	slice2fat := slice1 [:4] [5:len(slice1)]
	slice3fat := slice1 [2:8]

	fmt.Println(slice2fat,slice3fat)
}
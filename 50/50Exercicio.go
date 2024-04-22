package main   

import (
	"fmt"
)
/*
- Começando com a seguinte slice:
    - x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
- Anexe a ela o valor 52;
- Anexe a ela os valores 53, 54 e 55 utilizando uma única declaração;
- Demonstre a slice;
- Anexe a ela a seguinte slice:
    - y := []int{56, 57, 58, 59, 60}
- Demonstre a slice x.
*/
func main (){

	slicex := []int {42,43,44,45,46,47,48,49,50,51}

	slicex = append(slicex, 52)
	fmt.Println(slicex)

	slicex = append(slicex,53,54,55)
	fmt.Println(slicex)

	slicey := []int {56,57,58,59,60}

	slicex = append (slicex,slicey...)

	fmt.Println(slicex)
	
}
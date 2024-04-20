package main

import (
	"fmt"
)


// O método make na hora de criar uma slice serve pra tu definir o tamanho e a capacidade máxima
// desse slice, serve pra quando vc ja tem ideia de quantos elementos terá que lidar dentro daquela
// slice, isso tem valor pelo fato de que como slice funciona por cima de um array, sempre que vc
// da append em uma slice o programa recria o array subjacente com um novo tamanho, e isso custa 
// desempenho dms

func main() {
	slice := make([]int, 5, 10) // slice := make ( []t, len, cap )
	
	slice[0], slice[1], slice[2], slice[3], slice[4] = 1, 2, 3, 4, 5
	
	slice = append(slice, 6)
	slice = append(slice, 7)
	slice = append(slice, 8)
	slice = append(slice, 9)
	slice = append(slice, 10)
	
	fmt.Println(slice, len(slice), cap(slice))
	
	slice = append(slice, 10)
	
	fmt.Println(slice, len(slice), cap(slice))

}
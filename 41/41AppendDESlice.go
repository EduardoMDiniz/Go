package main

import "fmt"

//Append básico de novos valores a uma slice

func main() {

	slice1 := []int{1, 2, 3, 4, 5}
	slice1 = append(slice1, 6, 7, 8, 9) //adicionando valores a 1 slide básico
	fmt.Println(slice1)

	sub()

}

// Apend básico de 1 slice a outra slice
func sub() {

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{11, 12, 13, 14}

	slice1 = append(slice1, slice2...)     //aqui para dar um append de uma slice em outra precisamos
	fmt.Println(slice1)                // desse ... no final do argumento pra descompactar o slice
}                                      // para os valores deles, pq o append so espera os valores
                                       // e não o slice
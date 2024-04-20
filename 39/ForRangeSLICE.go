package main

import (
	"fmt"
)


func main (){
	
	slice1 := []string{"eh mole","eh duro","eh paia"}
	for indice, valor := range slice1 {                //o range vai percorrer todo o slice
		fmt.Println(indice,valor)                      // e retornar o indice e o valor daquele indice
	}									   // com isso simplesmente alocamos nas VAR indice e valor
	                                          
	slice2 := append (slice1, "eh crucial")
		fmt.Println(slice2)                               
}

//se eu quiser fazer a adição de um valor ao slice eu tenho que fazer da seguinte maneira

// slice2 := append (slice1, "eh crucial")
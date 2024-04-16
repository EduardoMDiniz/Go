package main

import (
	"fmt"
)

var x int = 42
var y string = "james bond"
var z bool = true

func main(){
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z) 
	fmt.Println(s)                       
}

// aqui basicamente tinhamos que usar o Sprintf para atribuir todos as variáveis anteriores ao S
// o Sprint retorna todas as variáveis em String, com isso atribuimos o valor desse retorno no S 
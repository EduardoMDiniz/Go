package main

import (
	"fmt"
)

type hotdog int  //criando o tipo hodtdog

var b hotdog = 101      //Declarando e atribuindo valor a variável b do tipo Hotdog

func main(){
	x := 10                       //crinado a variável x do tipo int
	fmt.Printf("%v \n", x)
	
	x = int(b)                    //CONVERTINDO o tipo hotdog da variavel B para int usando o INT do x
	fmt.Printf("%v \n", x)  // que na vdd é so atribuir o valor da var B para X que é INT, ent ai a gente
	                        // tem o valor de B em INT so que agora na variavel X
	
	// ent a estrutura de conversão é
	// variavel qualquer do tipo q tu quer = tipo (variavel com o tipo que tu quer mudar)

}
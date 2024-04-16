package main

import (
	"fmt"
)

// Existem 2 switch, o default e o com Statement, O default ele recebe um True como comparação
// é como se ele recebesse o switch TRUE {...} e fosse comparando cada caso com o true, ou seja, ele
// ta avaliando se cada condicional de case é TRUE e volta a 1 que for TRUE.

// a Segunda é um com uma declaração especifica e ele vai escolher o caso que for igual aquele statement

// Ou seja, Switch Default sem nada na frente compara os case com true e o switch statement compara
// com o statement dado

//exemplos a baixo, primeiro default e segundo Statement

func main() { 

	x := 10
	switch { ////switch default

	case x < 5:                       //switch ja abre a função com {} e coloca os case dentro
		fmt.Println("x Menor que 5")  //Note que termina com : e não ponto e virgula ;
	case x == 5:                      // pq faz parte da sintese do switch, e não é so o final de
		fmt.Println("x Igual 5")      // uma instrução como no For x:= 0; x<10; x++ por exemplo
	case x > 5:
		fmt.Println("x Maior que 5")
	default:
		fmt.Println("Não é nenhuma das opções") // o default é um "tratamento de erro" é quando n 
	}                                          // atende nenhum dos cases
	sub() 
	sub2()
}

func sub() {        

	carro := "prata" 
	switch carro {      //switch statement

	case "vermelho":
		fmt.Println("Carro é vermelho")
	case "preto":
		fmt.Println("Carro é preto")
	case "prata":
		fmt.Println("Carro é prata")
	default:
		fmt.Println("Não é nenhuma das opções") // O default é um "tratamento de erro" é quando 
		                                        // Não atende nenhum dos cases
			
	}
}

// se eu quiser que se uma opção seja verdade outra automaticamente seja também
// eu preciso colocar fallthrough entre elas


// exemplo de pq é mais bonito fazer switch do que if 

func sub2() {
	carro := "prata"

	if carro == "preto" {
		fmt.Println("Carro é preto")
	} else {
		fmt.Println("Carro é prata")
	}
}

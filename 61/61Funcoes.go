package main

import (
	"fmt"
)

/*
Aqui eu falarei sobre as funções em específico, apesar de ter usados a função main pra me retornar
coisas até agora, agora eu vou documentar como usar uma func de forma integral em golang

A estrutura básica de uma func é  func (receiver) nome (parametros) (returns) { code }

    - Função básica.----------------------------------------------------------------------------

	A função básica é quando vc usa a estrutura dela apenas pra um codigo dentro de { code }
        - Go Playground: https://play.golang.org/p/FebJblBenP

    - Função que aceita um argumento. ----------------------------------------------------------

     A Função que vc pode passar argumentos pra o { codigo }  agir sob eles
        - Go Playground:  https://play.golang.org/p/CE6Ij3U4QB

    - Função com retorno. ----------------------------------------------------------------------
     A função pode te retornar um valor baseado no codigo e nos argumentos
        - Go Playground: https://play.golang.org/p/gKxwYe6btP

    - Função com múltiplos retornos e parâmetro variádico.--------------------------------------
     Função que retorna multiplos valores e tem quantidade infinita de argumentos

        - Go Playground: https://play.golang.org/p/OcQ1wXwM2c

    - Mais um: https://play.golang.org/p/8wc2TA9xH_  */

//  func (receiver) nome (parametros) (returns) { code }

func main() {

	basica()
	argumento("Tarde")
    chamando_soma ()
    chamando_soma_variadica()
}

func basica() {                    // Função Básica
	fmt.Println("Oi, Bom dia")
}

func argumento(s string) {          // Função com argumentos
	if s == "Manhã" {
		fmt.Println("Oi, Bom dia")
	} else if s == "Tarde" {
		fmt.Println("Boa tarde")
	} else {
		fmt.Println("boa noite")
	}
}


func soma(x, y int) int {     // A soma sendo a função com retorno "int" podendo ser mais de 1
	return x + y              // so colocar (int,int) pra ser 2 ints por exemplo
}


func chamando_soma () {    // a função q eu chamo a soma

	valor := soma(10, 10)

	fmt.Println(valor)
}


func soma_variadica (x ...int) (int, int,) { // inicializo a função passando x variadico com x....
	soma := 0                                //passo dps (int, int) pra retornar 2 valores
	for _, v := range x {                    // de resto é so um for pra somar os itens
		soma += v             
	}
	return soma, len(x)
}

func chamando_soma_variadica() {        //aqui eu chamo a func e passo os valores na frente

	total, quantos := soma_variadica(10, 10, 1, 2, 3, 5)

	fmt.Println(total, quantos)
}

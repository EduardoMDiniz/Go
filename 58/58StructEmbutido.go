package main      

import (
	"fmt"
)

/* Basicamente, da de vc passar um struct como caracteristica de outro, vc so precisa preencher
as caracteristicas dele como se fosse preenchendo aquela struct do lado de fora
exemplo: */

type pessoa struct{
	nome  string
	cpf   int 
}

type emprego struct{
	pessoa
	nomeEmprego  string 
	salario      int 
}

func main (){
	pessoa1 := emprego{
		pessoa: pessoa {
			nome: "Alice",
			cpf: 94034934,
		},

		nomeEmprego: "Secretaria",
		salario: 100000,

	}
	fmt.Println(pessoa1)

}
/* Aqui o que faço na função main é basicamente, criar o obj pessoa1 com tipo de obj (struct)
pessoa pra que ele herde os campos de pessoa, semelhante a o que ocorre na herença de classe

Discleimer importante do Copilot:

No seu exemplo, pessoa e emprego são tipos de estrutura (Obj). emprego inclui pessoa como um 
campo anônimo, o que significa que emprego herda todos os campos de pessoa. Isso é semelhante
a como  um objeto em uma linguagem orientada a objetos pode herdar campos de uma classe pai.
*/


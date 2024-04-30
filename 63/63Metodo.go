package main


import(
	"fmt"
)

/*O método é atribuir uma funcionalidade a um tipo, ou seja, é uma função que só funciona
com um tipo especifico de dado. 
E a gente passa o método pra uma função usando o receiver 

func (receiver) nome (parametros) (returns) { code } */

type pessoa struct{
	nome string 
	idade int 
}

func (p pessoa) oibomdia (){           //criei o método P e passei o tipo pessoa dei nome ao método
	fmt.Printf(p.nome,"diz bom dia")  // de oibomdia e no {code} escrevi o que ele faz
} 

func main (){

	mauricio := pessoa {  //criei uma variavel do tipo pessoa pra poder usar o método
		nome: "Mauricio", 
		idade: 20,
	}
	mauricio.oibomdia() //usei o método
}

/* o P é o receptor do tipo pessoa na função oibomdia, então no codigo quando eu for chamar
a variavel do tipo pessoa mauricio e passar o método oibomdia o mauricio é exatamente o P
mauricio.oibomdia // ou seja, o p é so uma forma de representar que ali vai vir uma variavel
do tipo pessoa */
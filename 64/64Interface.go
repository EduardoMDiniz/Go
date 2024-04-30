package main

import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
}

type dentista struct {
	pessoa
	salario     float64
	trabalhaCom string
}

type arquiteto struct {
	pessoa
	salario     float64
	trabalhaCom string
}

func (d dentista) oibomdia() {
	fmt.Printf("Oi, sou dentista e tenho o salário de %v", d.salario)
}

func (a arquiteto) oibomdia() {
	fmt.Printf("Oi, sou dentista e tenho o salário de %v", a.salario)
}

type gente interface {
	oibomdia()
}

func serhumano(g gente) {
	fmt.Println("\n-------Função serhumano-------")

	switch g.(type) {

	case dentista:
		g.oibomdia() //além de passar o método do dentista também posso passar outras coisas como esse fmt.println
		fmt.Println(g.(dentista).nome, "é alguém que trabalha com ", g.(dentista).trabalhaCom)

	case arquiteto:
		g.oibomdia() //além de passar o método do arquiteto também posso passar outras coisas como esse fmt.println
		fmt.Println(g.(arquiteto).nome, "é alguém que trabalha com ", g.(arquiteto).trabalhaCom)
	}
}

//criando os objetos de cada tipo já dentro da main

func main() {

	rodrigo := dentista{
		pessoa: pessoa{
			nome:  "rodrigo",
			idade: 25,
		},
		salario:     4000.0,
		trabalhaCom: "Dentes",
	}
	romario := arquiteto{
		pessoa: pessoa{
			nome:  "romario",
			idade: 35,
		},
		salario:     6000.0,
		trabalhaCom: "Predios e Casas",
	}

	fmt.Printf("Este fmt.Pritf é so pra utilizar os objetos %v e %v\n", rodrigo, romario)

	rodrigo.oibomdia() //forma de chamar comum direto pelo método
	serhumano(rodrigo) //forma de chamar a função que aplica a interface que contem o método oibomdia do tipo rodrigo

	fmt.Println("-------------------------------------------")

	romario.oibomdia() //forma de chamar comum direto pelo método
	serhumano(romario) //forma de chamar a função que aplica a interface que contem o método oibomdia do tipo romario

}

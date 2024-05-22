package main 

import (
	"fmt"
)

//tipos

type salario int 

type medico struct{
	nome string
	salario
}

type engenheiro struct{
	nome string 
	salario
}

type farmaceutico struct{
	nome string
	salario
}

//objetos

var medicoJoao medico = medico{
	nome: "joao",
	salario: 0,
}

var engenheiroRoberto engenheiro = engenheiro{
	nome: "Roberto",
	salario: 0,
}

var farmaceuticoDamiao farmaceutico = farmaceutico{
	nome: "Damião",
	salario: 0,
}

//metodo calcule_salario para cada tipo diferente

func (m *medico) calculo_salario (x salario, y salario) {

	m.salario = x * y
}
func (e *engenheiro) calculo_salario (x salario, y salario) {

	e.salario = x + y
}
func (f *farmaceutico) calculo_salario (x salario, y salario) {

	f.salario = x * y
}

//criação da interface

type salariototal interface{
	calculo_salario(x salario, y salario)
}

//função que implementa a interface, recebendo além do s salariototal, que é os tipos de dentro da interface, recebe também
//os argumentos que os métodos da interface necessitam x salario y salario.
func func_que_implementa_interface (s salariototal, x salario, y salario){
	s.calculo_salario(x,y)
}

func main (){

	//print pra mostrar os valores iniciais em 0
	fmt.Printf("Salario do Médico João: %v\nSalário do Engenheiro Roberto: %v\nSalário do Farmaceutico Damião: %v\n",medicoJoao.salario,engenheiroRoberto.salario,farmaceuticoDamiao.salario) 

	var x salario
	var y salario

	fmt.Println("Insira x e y para salario do médico")
	fmt.Scan(&x,&y)
	func_que_implementa_interface(&medicoJoao, x,y )

	fmt.Println(medicoJoao.salario)
	//

	fmt.Println("Insira x e y para salario do engenheiro")
	fmt.Scan(&x,&y)
	func_que_implementa_interface(&engenheiroRoberto, x,y )

	fmt.Println(engenheiroRoberto.salario)
	//

	fmt.Println("Insira x e y para salario do Farmaceutico")
	fmt.Scan(&x,&y)
	func_que_implementa_interface(&farmaceuticoDamiao, x,y )

	fmt.Println(farmaceuticoDamiao.salario)

	fmt.Println("Você atribuiu os salarios dos 3 profissionais, como prova, printarei eles dnv a baixo")

	fmt.Printf("Salario do Médico João: %v\nSalário do Engenheiro Roberto: %v\nSalário do Farmaceutico Damião: %v\n",medicoJoao.salario,engenheiroRoberto.salario,farmaceuticoDamiao.salario) 
}


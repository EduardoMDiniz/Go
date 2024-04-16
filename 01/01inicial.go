package main

import (
	"fmt"
)

func main()  {
	fmt.Print(Hello World)
	
} // estrutura basica

// Em GO uma variável não pode ser existir sem ser usada, então quando recebemos um valor e não queremos
// alocar espaço na memória para ele, usamos o underline _ para ignorar o valor.
// Exemplo: 

func main()  {
	_ , 02 := fmt.Print(Hello World, Germany)
}
// Nesse caso o valor de Hello World é ignorado e não é alocado espaço na memória para ele.

// Aqui a gente atribui valores pra uma variável usando o operador de atribuição :=

func main()  {
	x := 10
	y := 20
	fmt.Println(x, y)
}

// esse operador é curto, ou seja, ele faz a declaração e a atribuição ao mesmo tempo, ele
// declara o tipo da variavel automaticamente

// A declaração e atribuição usando o  :=  so funciona dentro de funções, fora de funções
// A gente usa o operador = para atribuir valores a variáveis

// o operador = só faz a atribuição, ele não declara a variável, ent serve para mudar o valor da variavel
//dps de ela estar declarada,

//atribuição de variáveis globais fora de funções é feita com o operador =, exemplo:

var x int = 10 

//Operadores mais chatos == Comparação de igualdade, != Diferente, > Maior, < Menor, >= Maior ou igual, <= Menor ou igual

//pra ver o tipo de variável basta dar um 
fmt.Println("%T", variavel)

// Declarando varias variáveis e vendo o tipo delas em seguida 

var a int = 5
var b float = 5.6
var c string = "oi"
var d bool = false 

func main () {
	fmt.print("%T",a)
	fmt.print("%T",b)
	fmt.print("%T",c)
	fmt.print("%T",d)
}

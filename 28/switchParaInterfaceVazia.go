package main  

import(
	"fmt"
)

//a ideia aqui é usar o switch para executar uma ação baseada no tipo da variavel

//variavel do tipo interface vazia e mudamos dentro do codeblock e usamos o switch
//para definir a ação baseada no novo tipo
//essa ideia da interface vazia é boa pra isso, quando precisamos declarar uma variavel antes do 
//codeblock mas n sabemos o tipo, e so iremos declarar ele quando atribuirmos o conteudo da variavel
//e o switch é perfeito pra essa situação pq é justamente por ele que iremos escolher a ação dependendo
//do tipo de variável que tivermos

var a interface{} 

func main (){
	a = 10
	switch a. (type){     //a estrutura é levemente diferente, precisamos do variavel.(type)
	case int:
		fmt.Println("A é Inteiro")
	case float64:
		println("A é float")
	case bool:
		println("A é booleano")
	}
		
}
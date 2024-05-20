package main

import (
	"fmt"
)

/* Arquivo para considerações sobre ponteiros

Já que o Ponteiro aponta pra um lugar na memoria x:= 10, ponteiro := &x quando a gente passar o ponteiro
Em uma função que altere o valor da variavel, o valor da variavel vai mudar onde quer q ela esteja
Isso não aconteceria normalmente, já que go é pass by value, sempre q passamos uma var ele faz uma copia
Dentro de onde a gente chamou aquela variável, Exemplo: */

func main() {

	x := 10
	fmt.Println(x) // saída 10

	recebe_a_var_normal(x) //saída 11 sendo uma cópia ja que no anterior o x continua 10
	
	recebe_o_ponteiro(&x)  //saída 11 alterando diretamente o valor de x
	fmt.Println(x)         //saída 11 provando que o valor de x foi alterado
}

func recebe_a_var_normal(x int) {
	x++
	fmt.Println(x)
}

func recebe_o_ponteiro(x *int) {
	*x++
	fmt.Println(*x)
}

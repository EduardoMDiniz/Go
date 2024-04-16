package main 

// O fmt "fâmiti" é o pacote que tu importa pra usar os 'prints' do GO

import (
	"fmt"
)

//exitem 3 principais tipos, os:
 fmt.Print
 fmt.Println
 fmt.Printf

// A diferença básica é, entre o Print e o Println é que o Println vai printar add espaços entre os
// valores

// Existem tb 3 categorias de print sendo:

fmt.Print // onde vai ter os Print/Println/Printf
fmt.Sprint // onde vai ter os Sprint/Sprintln/Sprintf
fmt.Fprint // onde vai ter o Fprint/Fprintln/Fprintf

// a diferença é o Print tem como função imprimir o valor na tela, equanto o Sprint o valor em forma
// de String independente da tipagem da variavel etc
// o Fprint é basicamente o print para writer interface, isto é qualquer coisa que receba bytes
// seja arquivos ou resposta de servidor etc

No geral na vida real a gente vai usar o Printf pra marioria pois
queremos sempre 

fmt.Printf("%v", x) isso é, usar o %v pra colocar o argumento X dentro da strinng "String"

Se a gennte quiser add espaços entre as stringss podemos simplesmente colocar uma strring " " ent
o println fica meio irrelevante
package main

import (
	"fmt"
	"sort"
)

func main (){

	ss := []string {"Água","Refrigerante","Cangica","Fanta"} 

	fmt.Println(ss)

	sort.Strings(ss)
	fmt.Println(ss)
}
	//estrutura - pacote.func(var) sempre

/* simples uso do método sort do pacote sort para ordenação. Nesse caso, usamos a func do pacote sort para string que retorna em ordem
alfabetica. */

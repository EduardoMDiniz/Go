package main

import (
	"fmt"
)

// Aqui eu demonstro um problema do slice quando se faz a seleção de uma fatia de um slice
// O programa so altera a ordem do primeiro slice e diminui o cap dele pra quantidade de itens
// pra poder criar o novo segundoslice, e dai quando vc printa ou usa o primeiroslice ele vai
// sair com os itens em ordem diferente de quando foi criado; Exemplo abaixo

func main() {

	primeiroslice := []int{1, 2, 3, 4, 5}
	
	fmt.Println(primeiroslice)
	
	segundoslice := append(primeiroslice[:2], primeiroslice[4:]...)

	fmt.Println(segundoslice)

	fmt.Println(primeiroslice)
	
	sub()
}

// A forma correta de fazer isso sem afeter a integridade do primeiro slice é
//inicializando o segundo slice com um slice vazio
// segundoslice := append ([] int {}, primeiroslice...) 
// ai dps vc faz o append dos valores do primeiro slice pro segundo 

func sub (){

	primeiroslice := []int{1, 2, 3, 4, 5}
	fmt.Println(primeiroslice) // output [1 2 3 4 5]

	segundoslice := append([]int{}, primeiroslice...) // inicializa o segundo slice tendo um slice vazio como base
	segundoslice = append(segundoslice[:2], segundoslice[4:]...) // agora sim fazemos o append que queremos

	fmt.Println(segundoslice) // output [1 2 5]
	fmt.Println(primeiroslice) // output [1 2 3 4 5] igual ao primeiro print


}

package main

import (
	"fmt"
)

const (
	_ = 2024 + iota    //aqui eu usei _ pra descartar o 2024 pois era so os proximos 4 anos e
	b                  // somei com a iota pra iterar 2024 + 1 da iota nas prox consts
	c                  // detalhe importante é que eu n coloquei = nas outras consts, pq na primeira
	d                  // const eu ja estou atribuindo o valor delas com a iteração 2024+1
) 

func main() {

	fmt.Println(b, c, d)
}

// a questão era pra imprimir os proximos 4 anos usando consts com iota

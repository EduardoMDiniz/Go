package main

import (
	"fmt"
	"sync"
)

/* 
o wg é a variavel de waitgroup do tipo Sync.WaitGroup que vai servir pra chamar as funções de GoRoutines, são elas:
wg.Add, que serve para colocar um numero N de funções para operarem de forma concorrente
wg.Wait que serve para colocar um marcador de espera entre agrupamento de funções concorrentes, no nosso caso duas
wg.Done é o marcador que se coloca no final da função pra sinalizar para o Wait que a função foi finalizada


 */

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	go func1()
	go func2()

	wg.Wait()

}

func func1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Func1 %v \n", i)
	}

	wg.Done()
}

func func2() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Func2 %v \n", i)
	}

	wg.Done()
}

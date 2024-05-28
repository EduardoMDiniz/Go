package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup  //waitgroup

func main (){

	contador := 0
	total_de_GoRoutines := 100

	wg.Add(total_de_GoRoutines) //iniciei o waitgroup

	for i := 0; i < total_de_GoRoutines; i++{ //for com func anonima dentro pra iterar 10x e emular 10 funções

		go func (){
			v := contador
			runtime.Gosched() //sleep time
			v++
			contador = v
			wg.Done()
		}()
			
	}
	wg.Wait()
	fmt.Println(contador)
	
}

/* essa função anonima itera 10x e simula 10 funções que tentariam acessar e utilizar da mesma variável que é o contador
nisso, o contador sempre vai ser diferente em cada execução pq essa concorrencia de varias "funções" pela mesma variavel
faz com que o append de +1 ocorra de forma desordenada.
Isso é uma condição de corrida 

Uma forma de descobrir se existe condição de corrida em um programa é quando for executar o go run programa.go, executar
go run -race programa.go

ele vai fazer uma análise pra ver se existe condição de corrida no programa

*/


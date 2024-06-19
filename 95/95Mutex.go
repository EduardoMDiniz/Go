package main 

import (
	"fmt"
	"runtime"
	"sync"
)

//Mutex é basicamennte exclusão mutua, ele faz com que apenas uma GoRoutine possa usar aquela variável até finalizar seu uso dela

var wg sync.WaitGroup  //waitgroup

var mu sync.Mutex //criamos uma var mu do tipo sync.mutex e ganhamos 2 métodos com esse tipo, lock e unnlock

func main (){

	contador := 0
	total_de_GoRoutines := 100

	wg.Add(total_de_GoRoutines) //iniciei o waitgroup

	for i := 0; i < total_de_GoRoutines; i++{ //for com func anonima dentro pra iterar 10x e emular 10 funções

		go func (){

			mu.Lock()        //usar lock no começo da func pra travar as variaveis que for usar 

			v := contador
			runtime.Gosched() //sleep time
			v++
			contador = v

			mu.Unlock()      //usar unlock pra soltar as variaveis da func nan hora que terminar a func

			wg.Done()
		}()
			
	}
	wg.Wait()
	fmt.Println(contador)
	
}

//resultado = 100 pq dessa vez a integridade da variavel se manteve
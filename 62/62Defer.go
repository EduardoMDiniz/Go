package main 

import(
	"fmt"
)

/*Defer é uma palavra chave em go que permite que tu deixa a execução de uma instrução pra dps
Se mais uma istrução tiver defer antes a ordem vai ser de quem teve seu defer lido primeiro
vai ser o ultimo */

func main (){
	defer fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
}
//com isso o output vai ser 3 4 2 1

/* ele é mt utilizado pra vc abrir um arquivo e logo em seguida dar a instrução de fechar usando
defer pra so fechar no final do codeblock, isso pq vc pode esquecer de colocar a instrução de 
fechar se vc deixar pra escrevar ela no final do codeblock, mas escrevendo a intrução de fechar
logo na hora que abre e deixando pro final com defer vc n corre risco de esquecer. */
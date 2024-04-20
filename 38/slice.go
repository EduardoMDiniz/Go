package main     

import(
	"fmt"
)

//o slide é basicament um array que tu pode declarar sem tamanho especificado e pode mudar o tamanho
// dele depois de criado


func main (){
	slice1 := []int{1,2,3,4,5}        //estrutura do slice basica, nome := [] tipo {valores}
	fmt.Println(slice1)
}
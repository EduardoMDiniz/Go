package main    

import (
	"fmt"
)

//retornar função

func main (){

	x := retornafuncao()
	y := x (30)            //pra chamar ela é so alocar ela em uma var e passar o argumento
	fmt.Println(y)
}

func retornafuncao () func(int) {   //aq no lugar de retonar um int ou outro tipo de dado como retornarfuncao () int { }
                                    //eu coloquei uma func (int) uma função que recebe um int, dentro eu faço a func
	return func(i int) int {        // dps pra chamar ela é so alocar ela em uma var e passar o argumento
			return i * 10 
	}
}
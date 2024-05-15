package main 

import(
	"fmt"
)
//criar função anonima

func main (){

	x := 200

	func (x int) {                     
		fmt.Println(x,"vezes 30 é:")
		fmt.Println(x*30)
	} (x)                              
  
}
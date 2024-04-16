package main

import (
	"fmt"
)

type hotdog int  //criando o tipo hodtdog

var x hotdog 

func main (){
	fmt.Printf("%T", x)
}
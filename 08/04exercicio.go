package main

import (
	"fmt"
)

type xereca int 

var x xereca 

func main(){
	fmt.Printf("%v\n",x)
	fmt.Printf("%T\n",x)
	x = 42
	fmt.Printf("%v\n",x)

}
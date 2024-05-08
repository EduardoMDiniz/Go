package main 

import(
	"fmt"
)

//- Closure é cercar ou capturar um scope para que possamos utilizá-lo em outro contexto

func main (){
	
	a := i()            //eu passando a função pra variavel como na aula anterior
	
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := i()            //eu passando a função pra variavel como na aula anterior

	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}

func i () func() int {

	x := 0              // a variavel q eu uso na func está dentro do scope da 1 func e não da de retorno
	return func() int {   // com isso eu posso reaproveitar a func pq ela sempre vai ter o resultado = var dela resetado quando 
		x++              // eu passar ela pra outrav var em outros scopes e etc  
		return x
	}
}

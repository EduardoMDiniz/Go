package main 

import(
	"fmt"
)

var x [10]int  //basicamente o array é uma estrutura de alocação de dados que o tipo vai ser o tamanho
               // [10] mais o tipo de dado no caso [10]int
			   // é importante lembrar que um array do tipo [10]int nn é a mesma coisa que um [11]int
	        	//ent n adianta ficar tentando converter usando y = [10]int(x)
func main (){
	x [0] = 100      //aqui eu estou atribuindo valor a um espaço do array, ja que ele estava vazio
	x [1] = 200
	fmt.Println(x)
	fmt.Println(len(x))    //aq to pedindoo tamanho do array com o len

}


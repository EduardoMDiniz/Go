package main

import (
	"fmt"
	"os"

)

//Demonstranndo a interface Writer do pacote IO, io.Writer.

func main (){

	file,_ := os.Create("file.txt")           //criamos apartir do método create um arquivo "file.txt"       
	n, err := file.Write([]byte("hello"))    //escrevemos nessa var file atráves de Write, passando o slice de bytes com hello

	fmt.Println(n,err)
	file.Close()
	
}

/*No código, os.Create("file.txt") retorna um valor do tipo *os.File, que você armazena na variável file. 
Como *os.File implementa io.Writer, você pode chamar file.Write([]byte("hello")) para escrever a string "hello" no arquivo.

Então, para esclarecer, *os.File implementa a interface io.Writer porque *os.File tem um método Write. 
Isso permite que você use um *os.File em qualquer lugar que um io.Writer seja esperado, como na chamada para file.Write no seu código.
*/

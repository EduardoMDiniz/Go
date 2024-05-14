package main 

import (
	"fmt"
)
/*- - Crie um tipo struct "pessoa" que contenha os campos:
    - nome
    - sobrenome
    - idade
- Crie um método para "pessoa" que demonstre o nome completo e a idade;
- Crie um valor de tipo "pessoa";
- Utilize o método criado para demonstrar esse valor.
- Solução: https://play.golang.org/p/GBZcnu0Ajp */

type pessoa struct {
	nome string 
	sobrenome string
	idade int
}

var joao pessoa = pessoa { //lembrar sempre quando criar uma var de tipo composto = struct.  De repetir o 
	nome: "João",          // var x tipo = tipo {  }
	sobrenome: "Silva",
	idade: 20,
}

func (p pessoa) demonstrar_pessoa (){
	fmt.Println(p.nome,p.sobrenome,p.idade)
}

func main(){
	
	joao.demonstrar_pessoa()

}

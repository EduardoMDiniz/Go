package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//Bcrypt basicamente implemta hash como em qualquer outra linguagem, n ta no standard package da linguagem, ent tem q importar


func main() {

	senha := "stringsenha"
	senhaerrada:= "stringsenha1"

	sbsenha, erro := bcrypt.GenerateFromPassword([]byte(senha), 10)  //bcrypt.GenerateFromPassword ([]slice(senha),int do custo comp)
	if erro != nil {
		fmt.Println(erro) //tratamennto de erro
	}

	fmt.Println(string(sbsenha)) //vai printar o hash do nosso sbsenha que é um slice de byte
	
	fmt.Println(bcrypt.CompareHashAndPassword(sbsenha, []byte(senhaerrada))) //print comparando a senha correta com a errada
	//retorna o erro hashedPassword is not the hash of the given password


	fmt.Println(bcrypt.CompareHashAndPassword(sbsenha,[]byte(senha)))  //print comparando com a senha correta
	// retorna nil, ou seja, nada.

}
/* repare que na função bcrypt.GenerateFromPassword ([]slice(senha),int do custo comp)
existe um argumento que recebe um int que significa o custo operacional da criação do hash, o default é 10, mas podeser
menos ou mais */
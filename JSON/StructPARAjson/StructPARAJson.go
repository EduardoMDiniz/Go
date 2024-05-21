package main

import (
	"encoding/json"
	"fmt"
)

/* DISCLAIMER 
Para exportar o nosso objeto para json tudo do objeto, isto é, o struct precisa precisa ser com letra Maiuscula desde o nome
aos atributos. */

type Pessoa struct{

	Nome string
	Sobrenome string 
	Idade int
	Est_civil string
}

func main () {

	joao := Pessoa {
		Nome: "João",
		Sobrenome: "Silva",
		Idade: 20,
		Est_civil: "Solteiro",
	}

	joaoJson, error := json.Marshal(joao)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(string(joaoJson))  // aq eu passo string na hora de mostrar a var em json
}

/*Simples demonstração de conversão de uma struct para json utilizando metodo marshal do pacote json que retorna 2 valores
sendo o primeiro o arquivo json, e o segundo um eventual erro.
O if é para caso a variavel para erro não seja vazia, ou seja, caso ela tenha algo, recebermos um println mostrando o erro. */
package main

import (
	"fmt"
)

/*
Struct é um tipo de variavel que serve pra guardar varios dados de tipos diferentes
Ele funciona como um objeto do JS por exemplo e outros, mas como n tem palavra reservada OBJ
aq, tu tem quem criar um tipo pra ele tal como qualquer tipo não primario
type tipoStruct struct {
	nome       string
	sobrenome  string
	cpf        float
}
percebe que eu passei o subtipo struct tal como se fosse qualquer outro tipo subjacente
*/

type TipoStruct struct {
	nome      string
	sobrenome string
	cpf       int
}

func main() {

	ronaldo := TipoStruct{
		nome:      "Ronaldo Nazário",
		sobrenome: "De Lima",
		cpf:       02443541622,
	}
	fmt.Println(ronaldo)
}

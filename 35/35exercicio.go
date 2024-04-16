package main  

import(
	"fmt"
)

func main (){
	esportefav := "futebol"
	switch esportefav {
	case "volei":
		fmt.Printf("Meu esporte favorito é volei")
	case "basket":
		fmt.Printf("Meu esporte favorito é basket")
	case "ping pong":
		fmt.Printf("Meu esporte favorito é Ping Pong")
	case "futebol":
		fmt.Printf("Meu esporte favorito é Futebol")
	}
}

//basicamente fazer um switch com statement, dessa vez eu passo algo pra ele encontrar o case
// que tenha o valor daquele algo
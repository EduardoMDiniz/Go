package main

import (
	"fmt"
)

func main() {
	for horas := 0; horas <= 12; horas++ {
		fmt.Println("Hora", horas)
		for min := 0; min < 60; min++ {
			fmt.Println("Minutos", min)
		}
	}
}

//basicamente Nested Loop = loop dentro do outro, aqui era o loop hora itera 1 e dps a cada 1 iterada
// ele vai iterar o min 60x e dps vai voltar e iterar a hora e até atender a condicional
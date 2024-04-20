package main

import "fmt"

// fatiar uma slice, basicamente pegar uma parte de uma slice


func main (){

	sabores := []string {"calabresa","pepperoni","queijo","bacon"}
	sabores_preferidos := sabores [0:2]
	fmt.Println(sabores_preferidos)  //resultado: calabresa, pepperoni isso pq o segundo indice é
	sub ()							 // onde vc traça a linha e descarta o valor dele
}                                    


// Como deletar um item de uma slice
// Basicamente eu crio uma slice nova ou dou um novo valor a slice existente com os itens de 
// antes e depois do item q eu quero jogar fora usando append

func sub (){
	sabores := []string {"calabresa","pepperoni","queijo","bacon"}
	sabores1 := append(sabores[:2], sabores [3:]...)
	fmt.Println(sabores1)
}

// o [3:]... é pra quando eu n puder eu n quiser definir o final ou inicio da seleção
// com o ... a linguagem entende que é tudo pós o 3
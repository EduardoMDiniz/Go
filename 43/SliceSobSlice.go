package main

import (
	"fmt"
)

//Aqui é basicamente como fazer slices multidimensionais, com colunas e linhas, tal como uma matriz

func main() {

	ss := [][]int{

	//  0   1   2
		{1, 2, 3}, //0
		{4, 5, 6}, //1
		{7, 8, 9}, //2    // existe um indice tanto pras linhas quanto pras colunas
	}                       
	fmt.Println(ss[1][2]) //ent se eu quiser colocar o 6 por exemplo, preciso colocar os 2 indices
}                         //começando pela coluna / linha 

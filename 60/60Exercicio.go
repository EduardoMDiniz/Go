package main      

import(
	"fmt"
)

/*
- Crie um novo tipo: veículo
    - O tipo subjacente deve ser struct
    - Deve conter os campos: portas, cor
- Crie dois novos tipos: caminhonete e sedan
    - Os tipos subjacentes devem ser struct
    - Ambos devem conter "veículo" como struct embutido
    - O tipo caminhonete deve conter um campo bool chamado "traçãoNasQuatro"
    - O tipo sedan deve conter um campo bool chamado "modeloLuxo"
- Usando os structs veículo, caminhonete e sedan:
    - Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
    - Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
- Demonstre estes valores.
- Demonstre um único campo de cada um dos dois.
- Solução: https://play.golang.org/p/3eoGb0kxzT */

type veiculo struct{
	portas int
	cor string 
}

type caminhonete struct{
	veiculo
	tracao4x4 bool
}

type sedan struct{
	veiculo
	modelo_luxo bool 
}

func main (){

	hilux := caminhonete{
		veiculo:veiculo{
			portas: 4,
			cor: "Branco",
		},
		tracao4x4: true,
	}
	corolla := sedan{
		veiculo:veiculo{
			portas: 4,
			cor:"Preto",
		},
		modelo_luxo: false,
	}

	fmt.Println(hilux,corolla)
	fmt.Printf("\n %v e %v", hilux.cor , corolla.cor)
}
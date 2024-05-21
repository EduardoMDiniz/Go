package main

import (
	"encoding/json"
	"fmt"
)
type TipoEmpresa struct {
	Empresa      string        `json:"empresa"`     // essas crases om uma versão daquele atributo escrito em json é uma tag
	AnoFundacao  int64         `json:"anoFundacao"` //tag json, serve para especificar na hora de converter para struct, qual é
	Ativo        bool          `json:"ativo"`       //o campo que vai receber aquele atributo json, é especialmente válido de usar
	Enderecos    Endereco    `json:"enderecos"`    //quanndo o campo da sua struct tem um nome diferente do campo do json mas 
	Funcionarios Funcionario `json:"funcionarios"` //precisa receber aquele valor
}

type Endereco struct {
	Rua    string `json:"rua"`   
	Numero int64  `json:"numero"`
	Cidade string `json:"cidade"`
	Estado string `json:"estado"`
	Pais   string `json:"pais"`  
}

type Funcionario struct {
	Nome    string  `json:"nome"`   
	Cargo   string  `json:"cargo"`  
	Salario float64 `json:"salario"`
}

func main (){

	sb := []byte (`{
        "empresa": "Tech Corp",
        "anoFundacao": 2005,
        "ativo": true,
        "enderecos": [
            {
                "rua": "Rua Principal",
                "numero": 123,
                "cidade": "São Paulo",
                "estado": "SP",
                "pais": "Brasil"
            },
            {
                "rua": "Avenida Secundária",
                "numero": 456,
                "cidade": "Rio de Janeiro",
                "estado": "RJ",
                "pais": "Brasil"
            }
        ]
    }`)
	var empresa TipoEmpresa

	erro := json.Unmarshal(sb, &empresa)
	if erro != nil{
		fmt.Println(erro)
	}
	fmt.Println(empresa.AnoFundacao)
}

/*simples demonstração de uma conversão de json para struct, inciialmennte, temos que ter os tipos que existem dentro daquele json
então, usamos https://jsonformatter.org/json-to-go para pegar esse tipos mais rápidos, dps disso alocamos o toda a estrutura json
dentnro de uma []byte, slice de bytes usando sb := []byte (`{json}`), dps criamos uma variavel do tipo principal do json
e iniciamos o método unmarshal que retorna 1 erro e recebe 2 argumentos, sendo o slice de bytes e o ponteiro com a localização de 
onde a gennte vai salvar o objto, dps é o tratamento de erro com if */

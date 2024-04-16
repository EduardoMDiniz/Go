package main     

import(
	"fmt"
)

func main (){                        //estrutura basica do loop for

	for x := 0; x < 10; x++ {   // FOR primeiro, variavel declarada e atribuida dps o ; (ponto e virgula)
		fmt.Println(x)          // dps do ; vem a condição de repetição x < 10 e dnv ; ai a ação
	}                           // for; declarar var se n existir ; condicional; ação {}
}

// também da de fazer coma  variavel ja declarada e atribuida fora do codeblock
// é so tirar ela de dps do FOR, ent ficaria tipo

// var x int = 0

// ... for; x < 10; x++


//estrutura novamente

// for; variavel; condicional; ação {}
// ou
// for; condicional; ação {}

// note que ao final sempre tem um ; que basicamente separa uma instrução de outra
// isso pq o go usa o ; pra finalizar todas instruções no codigo, apenas a gente n vê
// e pode escolher colocar ou não
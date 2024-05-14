package main 

import(
	"fmt"
)

/*- Crie um tipo "quadrado"
- Crie um tipo "círculo"
- Crie um método "área" para cada tipo que calcule e retorne a área da figura
    - Área do círculo: 2 * π * raio
    - Área do quadrado: lado * lado
- Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
- Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
- Crie um valor de tipo "quadrado"
- Crie um valor de tipo "círculo"
- Use a função "info" para demonstrar a área do "quadrado"
- Use a função "info" para demonstrar a área do "círculo"
- Solução: https://play.golang.org/p/qLY-q3vffQ */

var pi float64 = 3.14

//criando os tipos 

type quadrado struct{

	lado1 float64
	lado2 float64
}

type circulo struct {
	raio float64 
}

func (q quadrado) calc_area () float64 { //método pra calcular area do quadrado

	return q.lado1 * q.lado2 
}

func (c circulo) calc_area () float64 {  //método pra calcular area do circulo

	return 2 * pi * c.raio
}

type figura interface{ //criando a interface
	calc_area() float64
}

//criando func pra receber qualquer tipo que esteja no metodo figura

func info (f figura) float64 {
	fmt.Println("\n-------Função info-------")
	
	switch f.(type) {

	case quadrado:
		return f.calc_area()   
	
	case circulo:
		return f.calc_area()
	}
	return 0.0
}

//criando os obj struct dos tipos quadrado e circulo

var quadradasso quadrado = quadrado {

	lado1: 10,
	lado2: 15,
}

var circulozao circulo = circulo{
	raio: 30,
}

// func main

func main () {
	fmt.Println(info(quadradasso))
	fmt.Println(info(circulozao))
}
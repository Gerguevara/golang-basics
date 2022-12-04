package main

import (
	"fmt"
	"math"
)

func main() {
	// Rectángulo
	baseRectangulo := 20
	alturaRectangulo := 10

	areaRectangulo := baseRectangulo * alturaRectangulo

	fmt.Println("El Area del Rectángulo es :", areaRectangulo)

	// Circulo : AreaCirculo = pi por radio al cudrado
	const PI float64 = 3.14 // Constant
	var radioCirculo float64 = 10

	areaCirculo := PI * radioCirculo * radioCirculo

	fmt.Println("El Area del Circulo es :", areaCirculo)

	// Trapecio
	var baseUno float64 = 6
	var baseDos float64 = 15
	var alturaTrapecio float64 = 25

	areaTrapecio := ((baseUno + baseDos) * alturaTrapecio) / 2

	fmt.Println("El Area del Trapecio es :", areaTrapecio)

	main2()
	basicas()
}

func main2() {
	// area del rectangulo
	base := 4
	altura := 2
	println("Area del rectangulo ", base*altura)

	// area del trapecio
	base = 3
	baseAbajo := 2
	areaTrapecio := altura * (base + baseAbajo) / 2
	println("Area del trapecio ", areaTrapecio)

	// area del circulo
	radio := 2
	areaCirculo := math.Pi * math.Pow(float64(radio), 2)
	println(areaCirculo)
}

func basicas() {
	//Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)

	x := 10
	y := 50

	//Suma
	result := x + y
	fmt.Println("Suma:", result)

	//Resta
	result = y - x
	fmt.Println("Resta:", result)

	//Multiplicacion
	result = x * y
	fmt.Println("Multiplicació:", result)

	//División
	result = y / x
	fmt.Println("División:", result)

	//Modulo
	result = y % x
	fmt.Println("Modulo:", result)

	//Incremental
	x++
	fmt.Println("Incremetal:", x)

	//Decremental
	x--
	fmt.Println("Decremental:", x)
}

// incrementar x++
//decrementar x--
//modulo y % x

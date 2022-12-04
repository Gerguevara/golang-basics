// en go siempre se debe declarar un package en este cao es el archivo principal asi que usa main
package main

// al guardar la extension de go nos importa automaticamente los paquetes utilizados
import "fmt"

func main() {
	//Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.16
	fmt.Println("Pi:", pi)
	fmt.Println("Pi2:", pi2)

	//Declaracion de varaibles
	base := 12          //Primera forma, := indica que la variable no esta declarada asi que de un solo la declara y asigna, y automaticamente toma su tipo
	var altura int = 14 //Segunda forma
	var area int        //Go no compila si las variables no son usadas

	fmt.Println(base, altura, area)
	//Zero values
	//Go asigna vaalores a variables vacías por defecto el no  asigna null
	var a int     // es 0
	var b float64 // es 0
	var c string  // es ''
	var d bool    // es	false

	fmt.Println(a, b, c, d)

	//Ejercicio
	//Calcular el áera del cuadrado
	const baseCuadrado = 10
	// areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("El área del cuadrado es:", baseCuadrado*baseCuadrado)
}

//para ejecutar en desarrollo
//go run src/main.go

// para compilar
// go build src/main.go
//para ejecutar
//./main

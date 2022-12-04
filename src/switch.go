package main

import "fmt"

func main() {
	getDia(7)
	isOdd()
}

//switch

func isOdd() {
	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println(" No Es par")

	}
}

//switch anidado, no se declara condicion junto al switch siuno que se deja abierto
// y dentro ponemos las condiciones

func getTrimestre() {
	var mes int8 = 10
	switch {
	case mes <= 3:
		fmt.Println("Primer Trimestre")
	case mes > 3 && mes <= 6:
		fmt.Println("Segundo Trimestre")
	case mes > 6 && mes <= 9:
		fmt.Println("Tercer Trimestre")
	case mes > 9 && mes <= 12:
		fmt.Println("Cuarto Trimestre")
	default:
		fmt.Println("Este no es un mes valido")
	}
}

func getDia(dayNumber int) {

	switch dayNumber {
	case 1:
		fmt.Println("Lunes")
		break
	case 2:
		fmt.Println("Martes")
		break
	case 3:
		fmt.Println("Miercoles")
		break
	case 4:
		fmt.Println("Jueves")
		break
	case 5:
		fmt.Println("Viernes")
		break
	case 6:
		fmt.Println("Sabado")
		break
	case 7:
		fmt.Println("Domingo")
		break
	default:
		fmt.Println("Ese no es un día valido de la semana!")
	}
}

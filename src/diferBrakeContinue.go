package main

import "fmt"

func main() {
	watchDefer()
	continueAndBreak()
}

//uso de defer

//defer hace que la funcion que tenga la linea hace que esa funcioin sea ejecutado
//justo antes de terminar la ejecucion y no en orden de linea

// se puede usar deffecr para cerrar conexiones  addgb por ejemplo

// es buena opractica usar un solo defer por funcion

func watchDefer() {
	defer fmt.Println("me ejecuto con defer")
	fmt.Println("yo me ejecuto antes pero me declararon despues")
}

// uso de continue y break (para un cliclo for)

func continueAndBreak() {
	var i uint8 = 0
	for i < 10 {
		// Continue
		if i == 2 {
			fmt.Println("¡El número que es 2!, no se necesario seguir buscando")
			i++
			// corta el codigo aqui y valta a la siguiente  recorrido
			continue
		}

		if i == 8 {
			fmt.Println("¡Break! haz llegaod muy lejos")
			// rompe el codigo aqui y ya no sigue a las siguientes lineas y termina el recorrido
			break
		}
		fmt.Println(i)
		i++

	}
}

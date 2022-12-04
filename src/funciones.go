package main

import "fmt"

func main() {
	// sendMEssage("Roberto", 9)
	fmt.Println(textBuilder("Gerardo", " Guevara"))
	a, b := dobleValueReturn(2)
	// para ignorar el valor en una funcion se utiliza _
	_, c := dobleValueReturn(3)
	fmt.Println(a, b, c)
}

func sendMEssage(message string, hora int) {
	fmt.Println(message, hora)
}

// buena practica agrupar funciones del mismo tipo
func textBuilder(nombre, apellido string) string {
	return " hola " + nombre + apellido
}

// go permite que una funcion retorne un doble valor
func dobleValueReturn(a int) (c, d int) {
	return a, a * 2
}

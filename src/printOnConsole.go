package main

import "fmt"

const a = "1000000000"
const c = 3

func main() {
	// a := 10
	// b := 45
	// fmt.Println(a + b + c)
	print()
}

func print() {
	nombre := "Udemy"
	cursos := 1000
	//la buena practica es darle el nombre de s o d segun su tipo de variable
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)

	// otro uso de fmtSprintF es la de almacenar esta concatenacion
	var message string = fmt.Sprintf("%v tiene más de %v cursos\n", nombre, cursos)
	fmt.Println(message)

	//asi mismo esta funcion nos ayuda a saber el tipo de variable que tenemos  agregando %T
	fmt.Printf("La variable 'nombre' es de tipo : %T\n", nombre)
}

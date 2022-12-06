package main

import "fmt"

func main() {
	//los array aw declaran son sun length  en este caso 4
	// y no puede cambiar es inmutabl
	var array [4]int
	//se asigna valor por posiciones, si no defecto toma el value
	//que go asigna por defecto basado en su tipo de dato
	array[0] = 1
	array[1] = 2
	// len indica el numero de elemntos en el array (length)
	// cap indica la capacidad maxima del array
	fmt.Println(array, len(array), cap(array))

	// Slice, a diferencia de los arrays no se le indica su length
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	// su cap y length son dinamicos
	fmt.Println(slice, len(slice), cap(slice))

	// Métodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])  //imprime hasta el indice 3 (indice 3 no incluido)
	fmt.Println(slice[2:4]) // imprime desde 2 hasta el 4 (indices no incluidos)
	fmt.Println(slice[4:])  // de el 4 en adelante (indice incluido)
	// los indice se cuentan desde el 1

	// Append, hace un push al slice
	slice = append(slice, 11)
	fmt.Println(slice)

	// Append
	newSlice := []int{12, 13, 14}
	// combinar  2 slice en uno
	// el ... indica que es como si hicera un for y luego push al slice
	//que se desea agregar
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}

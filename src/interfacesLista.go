package main

import "fmt"

func main() {
	//imitando un poco el comportamiento de javscript
	//un arreglo con muchos tipos diferentes, el primer {} se deja vacio
	myInterface := []interface{}{"hola", 123, false}
	fmt.Println(myInterface)
}

//lista de interfaces

package main

import "fmt"

type Car struct {
	brand string
	model int
}

func main() {
	myCar := Car{brand: "Honda", model: 2015}
	fmt.Println(myCar)
	// instanciacion 2da forma
	var myCar2 Car
	myCar2.brand = "Honda"
	//si no se especifica toma el valor por defecto
	//myCar2.model = 2014
	fmt.Println(myCar)
}

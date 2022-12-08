package main

import "fmt"

type figura2D interface {
	//un metodo que retornara valor
	area() float64
}

type cuadrado struct {
	base float64
}
type rectangulo struct {
	base   float64
	altura float64
}

// esta funcion pertenece a cuadrado
func (c cuadrado) area() float64 {
	return c.base * c.base
}

// esta funcion pertenece a reactangulo y aplica sobre carga de metods
func (r rectangulo) area() float64 {
	return r.base * r.altura
}

// esta funcion solo toma como entrada la interfaz dada
// y la fuerza a ejecutar su funcion de area
func calcular(f figura2D) {
	fmt.Println(f.area())
}

// notas: en no es necesario decir que un struct implementa directamente
// una interfaz bas con que tenga un mismo metodo o propiedad la cual
// tiene la interfaz y la funcio que esta utilizara

func main() {
	cuadrado := cuadrado{base: 40}
	rectangulo := rectangulo{altura: 50, base: 40}

	calcular(cuadrado)
	calcular(rectangulo)
}

//lo ideal es usar interface condo los stuts comparte una misma funcion

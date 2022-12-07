package main

import "fmt"

func main() {
	a := 50
	// con el & indicamos que es un puntero de memoria de a (paso por referenca)
	//sino se usa y solo se pone a simplemente toma el valor primitivo
	b := &a
	fmt.Println(a)
	// esta asignacion indica que apuntara a la  posicion de meoria
	//donde esta a
	fmt.Println(b)
	// con * accedemos al valor de ese paso por referencia
	fmt.Println(*b)
	// si modificamos un valor utilizando * afectamos al orgina
	//*b = 100
	//fmt.Println(a) // a sera 100
	ejercicioPunteros()
}

//los punteros en go no son mas que acceso a la memoria

type pc struct {
	ram   int
	disk  int
	brand string
}

// funciones con acceso a propiedades de un struct
func (myPc pc) ping() {
	fmt.Println("brand: ", myPc.brand)
}

func (myPc *pc) duplicateRam() {
	myPc.ram = myPc.ram * 2
}

// estas funciones puede ser usadas de geters y seter tambien
func ejercicioPunteros() {
	myPc := pc{ram: 16, disk: 200, brand: "apple"}
	fmt.Println(myPc)
	myPc.ping()
	fmt.Println("ram actua", myPc.ram)
	myPc.duplicateRam()
	fmt.Println("ram actual dyplicada", myPc.ram)
	//  ver en consola de forma mas optimizada
	myPC2 := pc{ram: 20, disk: 500, brand: "lenovo"}
	fmt.Println(myPC2)
}

//IMPRIMER EN CONSOLA DE FORMA MAS , similar a sobreescribir el metodo toString de java

func (myPc pc) String() string {
	return fmt.Sprintf("PC: %s con %dGB RAM y SSD %dGB", myPc.brand, myPc.ram, myPc.disk)
}

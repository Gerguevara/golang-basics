package mypackage

import "fmt"

// en go el nombrado con  inicial mayuscula denota
// que es una propiedad publica que puede ser consumida fuera del pauqete

// lo contrario si la primera letra es minuscula es private
type CarPublic struct {
	Year  int
	Brand string
}

type privateCar struct {
	year  int
	brand string
}

func PrintMessage() {
	fmt.Println("hola")
}

//en go siempre se debe declarar un package en este cao es el archivo principal asi que usa main
package main

// al guardar la extension de go nos importa automaticamente los paquetes utilizados
import "fmt"

func main() {
	fmt.Println("Hello, miau 🐈!")
}

//para ejecutar en desarrollo
//go run src/main.go

// para compilar
// go build src/main.go
//para ejecutar
//./main

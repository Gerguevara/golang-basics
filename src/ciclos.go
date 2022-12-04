package main

import "fmt"

func main() {
	// recorrerCicloCondicional(10)
	recorrerCicloWhile(100)
}

//recorre hasta  llegar al numero indicado
func recorrerCicloCondicional(numero int) {
	for i := numero; i >= 0; i-- {
		fmt.Println(i)
	}

}
func recorrerCicloWhile(counter int) {

	count := 0

	for count < counter {
		fmt.Println(count)
		count += count + 10
	}

}

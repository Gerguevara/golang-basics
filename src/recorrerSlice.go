package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	fmt.Println(isPalindromoEasy("anna"))
}

func recorridoBasico() {
	slice := []string{"hello", "world", "dude"}
	for i, value := range slice {
		fmt.Println(i, value)
	}

	//tambien podemos ignora algunas propiedades
	slice2 := []string{"hello", "world", "dude"}
	for _, value := range slice2 {
		fmt.Println(value)
	}

	//si solo interesa el indice no es necesario escapar el 2do valor
	slice3 := []string{"hello", "world", "dude"}
	for i := range slice3 {
		fmt.Println(i)
	}
}

func isPalindromoEasy(text string) bool {
	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		//hace un parse de a string pr sino lo que devuelve es el cod ascii
		textReverse += string(text[i])
	}

	return text == textReverse
}

// palindromo es palabra que se lee igual derecho o al reves
// otra forma es recorrer la palbra en sentido inverso y rearmarla
// y luego compararla
func isPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

// aprovechamos lo metros del paquete string y reflect
func isPalindromoWithFunctions(word string) bool {
	//corta el string
	splitWord := strings.Split(strings.ToUpper(word), "")
	//hace una comparcion pero a nivel de objeto y ve el conenido
	// sino hari una comparacion a nivel de puntero en memoria
	return reflect.DeepEqual(splitWord, reverseString(splitWord))
}

func reverseString(input []string) []string {
	if len(input) == 0 {
		return input
	}
	//conrecurividad
	return append(reverseString(input[1:]), input[0])
}

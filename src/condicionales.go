package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	if isPair(6) {
		fmt.Println("Number is pair")
	} else {
		fmt.Println("Number is odd")
	}
	// example login
	if isValidUser("Alpha5", "MyPassword") {
		fmt.Println("Credentials are valid")
	} else {
		fmt.Println("Credentials aren't valid")
	}
	parseNumber("asdasd")
}

func isPair(num int) bool {
	return num%2 == 0
}

func isValidUser(userName, pass string) bool {
	return userName == "Alpha" && pass == "MyPassword"
}

// los operadores en este lenguaje son los mismos que en todos  &&  || > !true  etc

func parseNumber(numberInString string) {
	//esta liberira hace conversione automaticas
	value, err := strconv.Atoi(numberInString)
	if err != nil {
		log.Fatalf(err.Error())
	} else {
		fmt.Println("Valor convertido: ", value)
	}
}

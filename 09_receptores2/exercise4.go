package main

import "fmt"

func esPar(number int) bool {
	if number%2 == 0 {
		return true
	}
	return false
}

func main() {
	//programa que indique si el numero ingresado es par o impar
	var numero int
	fmt.Print("Ingrese un numero entero: ")
	fmt.Scan(&numero)
	fmt.Println("Es par: ", esPar(numero))
}

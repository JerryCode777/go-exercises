package main

import "fmt"

func main() {
	var n, m, intentos int
	n = 10
	m = 15
	intentos = 0

	for true {
		var numero int
		fmt.Print("Adivina un numero entre n y m: ")
		fmt.Scan(&numero)
		intentos++
		if numero > n && numero < m {
			fmt.Println("Adivinaste! El numero si pertenece al rango")
			break
		}
	}
}

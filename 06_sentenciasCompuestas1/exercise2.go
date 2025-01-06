package main

import "fmt"

func main() {
	var numUsuario, intPos, intNeg int
	for {
		fmt.Print("Ingresa un numero entero(0 para finalizar): ")
		fmt.Scan(&numUsuario)

		if numUsuario > 0 {
			intPos++
		} else if numUsuario < 0 {
			intNeg++
		} else {
			break
		}
	}

	fmt.Println("Se han ingresado ", intPos, " enteros positivos y ", intNeg, " enteros negativo")
}

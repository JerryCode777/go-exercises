package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//juego del porfiado
	var numero int
	fmt.Print("Ingrese un numero entero del 1 al 10: ")
	fmt.Scan(&numero)

	if numero > 9 || numero < 1 {
		fmt.Println("Ingrese un numero valido")
		return
	}

	var numeroPc int
	for true {
		numeroPc = rand.Intn(9) + 1
		if numeroPc >= numero {
			break
		}
	}

	fmt.Println("Numero jugador: ", numero)
	fmt.Println("Numero PC: ", numeroPc)

	if numero > numeroPc {
		fmt.Println("Ganar Jugador!")
	} else if numero == numeroPc {
		fmt.Println("Empate")
	} else {
		fmt.Println("Gana Pc")
	}

}

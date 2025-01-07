package main

import "fmt"

func main() {
	var pntos, nroAutos int
	fmt.Print("Ingrese el numero de autos que ingresaron: ")
	fmt.Scan(&nroAutos)

	var Mayor, Menor int
	Menor = 100
	for i := 0; i < nroAutos; i++ {
		fmt.Print("Ingrese el numero de puntos contaminantes del auto ", i+1, " (0-100):")
		fmt.Scan(&pntos)
		if Mayor < pntos {
			Mayor = pntos
		}
		if Menor > pntos {
			Menor = pntos
		}
	}
	fmt.Println("Puntos del auto que mas contamino: ", Mayor)
	fmt.Println("Puntos del auto que menos contamino: ", Menor)
}

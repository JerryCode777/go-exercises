package main

import "fmt"

func main() {
	var i, nota, nroNotas float64

	for true {
		fmt.Print("Ingrese la nota del alumno(-1 para finalizar) : ")
		fmt.Scan(&nota)
		if nota == -1 {
			break
		}
		i = i + nota
		nroNotas++
	}
	promedio := i / nroNotas
	fmt.Println("el promedio es: ", promedio)
}

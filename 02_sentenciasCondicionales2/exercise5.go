package main

import "fmt"

func main() {
	var notaTeoria, notaLaboratorio int
	fmt.Print("Ingrese la nota de teoria: ")
	fmt.Scan(&notaTeoria)
	fmt.Print("Ingrese la nota de laboratorio: ")
	fmt.Scan(&notaLaboratorio)

	var notaFinal float32

	if notaLaboratorio >= 11 && notaTeoria >= 11 {
		fmt.Println("El estudiante aprueba!")
		notaFinal = float32(notaLaboratorio)*0.25 + float32(notaTeoria)*0.75
	} else {
		if notaLaboratorio > notaTeoria {
			notaFinal = float32(notaTeoria)
		} else {
			notaFinal = float32(notaLaboratorio)
		}
	}

	fmt.Println("Su nota final es: ", notaFinal)

}

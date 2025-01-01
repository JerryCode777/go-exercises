package main

import "fmt"

func main() {
	var notaTeoria, notaLaboratorio float32
	fmt.Print("Ingrese la nota de Teoria: ")
	fmt.Scan(&notaTeoria)
	fmt.Print("Ingrese la nota de laboratorio: ")
	fmt.Scan(&notaLaboratorio)

	var notaFinal float32

	if notaLaboratorio >= 13 {
		notaFinal = 0.75*notaTeoria + 0.25*notaLaboratorio
	} else {
		if notaLaboratorio > notaTeoria {
			notaFinal = notaTeoria
		} else {
			notaFinal = notaLaboratorio
		}
	}

	notaFinal = notaFinal + 0.5

	if notaFinal >= 13 {
		fmt.Println("El alumno esta aprobado!, promedio: ", notaFinal)
	} else {
		fmt.Println("El alumno esta desaprobado, promedio: ", notaFinal)
	}

}

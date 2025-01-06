package main

import "fmt"

func main() {
	var nota float64
	fmt.Print("Ingrese la nota: ")
	fmt.Scan(&nota)

	switch {
	case nota < 4.9:
		fmt.Println("Reprobado, repite el semestre")
	case nota < 10.4:
		fmt.Println("Reprobado, pasa a subsanacion")
	case nota < 15.9:
		fmt.Println("Aprobado")
	case nota <= 20:
		fmt.Println("Aprobado con distincion maxima")
	default:
		fmt.Println("Dato invalido, la nota maxima es 20")
	}
}

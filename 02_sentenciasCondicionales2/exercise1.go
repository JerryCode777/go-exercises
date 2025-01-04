package main

import "fmt"

func main() {
	var nota1 float64
	fmt.Print("Ingrese nota 1 del alumno: ")
	fmt.Scan(&nota1)

	var nota2 float64
	fmt.Print("Ingrese nota 2 del alumno: ")
	fmt.Scan(&nota2)

	var nota3 float64
	fmt.Print("Ingrese nota 3 del alumno: ")
	fmt.Scan(&nota3)

	promedio := (nota1 + nota2 + nota3) / 3

	if promedio >= 13 {
		fmt.Println("El estudiante esta aprobado!, su promedio es: ", promedio)
	} else {
		fmt.Println("El estudiante esta desaprobado!, su promedio es: ", promedio)
	}

}

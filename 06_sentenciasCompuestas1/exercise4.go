package main

import "fmt"

func main() {
	var nroAlumnos, nota, notaMenor, suma int
	fmt.Print("Ingrese el numero de alumnos: ")
	fmt.Scan(&nroAlumnos)

	for i := 0; i < nroAlumnos; i++ {
		fmt.Print("Ingrese al nota del alumno ", i+1, ": ")
		fmt.Scan(&nota)
		suma += nota
		notaMenor = 20
		if notaMenor > nota {
			notaMenor = nota
		}
	}
	fmt.Println("El promedio es: ", float64(suma)/float64(nroAlumnos))
	fmt.Println("La menor nota es: ", notaMenor)

}

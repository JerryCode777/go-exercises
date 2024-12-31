package main

import "fmt"

func main() {
	var salario float64
	fmt.Print("Ingrese el salario del empleado: ")
	fmt.Scan(&salario)

	if salario < 1000 {
		salario = salario * 1.2
	} else {
		salario = salario * 1.1
	}

	fmt.Println("El empleado gana: ", salario)
}

package main

import "fmt"

func main() {
	var salario float64
	fmt.Print("Ingrese el salario del empleado: ")
	fmt.Scan(&salario)

	if salario < 1000 {
		salario = salario * 1.2
	} else if salario < 5000 {
		salario = salario * 1.1
	} else {
		salario = salario * 1.05
	}

	fmt.Println("Se le debe pagar: ", salario)
}

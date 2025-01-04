package main

import "fmt"

func main() {
	//cuanto se le debe pagar a un empleado si gana menos de 1000 se le dara un bono del 10%
	var salario float64
	fmt.Print("Ingrese el salario del empleado: ")
	fmt.Scan(&salario)

	if salario < 1000 {
		salario = salario * 1.1
	}

	fmt.Println("El empleado gana: ", salario)

}

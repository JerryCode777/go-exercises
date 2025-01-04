package main

import "fmt"

func main() {
	var tiempo int
	var salario, utilidad float64

	fmt.Print("Ingrese su tiempo de antiguedad(anios): ")
	fmt.Scan(&tiempo)
	fmt.Print("Ingrese su salario: ")
	fmt.Scan(&salario)

	switch {
	case tiempo < 1:
		utilidad = salario * 0.05
	case tiempo < 2:
		utilidad = salario * 0.07
	case tiempo < 5:
		utilidad = salario * 0.10
	case tiempo < 10:
		utilidad = salario * 0.15
	default:
		utilidad = salario * 0.20
	}

	fmt.Println("La utilidad es: ", utilidad)
}

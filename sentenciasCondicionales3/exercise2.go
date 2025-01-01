package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Print("Ingrese el primer numero entero: ")
	fmt.Scan(&num1)
	fmt.Print("Ingrese el segundo numero entero: ")
	fmt.Scan(&num2)

	switch {
	case num1 == num2:
		fmt.Println("La multiplicacion es: ", num1*num2)
	case num1 > num2:
		fmt.Println("La resta es: ", num1-num2)
	default:
		fmt.Println("La suma es: ", num1+num2)
	}
}

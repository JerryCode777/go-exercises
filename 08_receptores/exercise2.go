package main

import "fmt"

func factorial(a int) int {
	if a < 0 {
		return -1
	}
	factorial := 1
	for i := 1; i <= a; i++ {
		factorial *= i
	}
	return factorial
}
func productoFactoriales(a, b int) (int, error) {
	factA := factorial(a)
	factB := factorial(b)
	if factA == -1 || factB == -1 {
		return 0, fmt.Errorf("no se puede calcular el factorial de numeros negativos")
	}
	return factA * factB, nil
}

func main() {
	var a, b int

	fmt.Println("Usted ingresara 2 numeros enteros: ")
	fmt.Print("Ingrese el valor de a: ")
	if _, err := fmt.Scan(&a); err != nil {
		fmt.Println("error")
		return
	}

	fmt.Print("Ingrese el valor de b: ")
	if _, err := fmt.Scan(&b); err != nil {
		fmt.Println("error")
		return
	}

	result, err := productoFactoriales(a, b)
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println("El producto de los factoriales es", result)
	}

}

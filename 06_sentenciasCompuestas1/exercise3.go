package main

import "fmt"

func main() {
	var n, m, numero, intentos int
	n = 12
	m = 14

	for {
		fmt.Println("Ingrese entre 1 y 20: ")
		fmt.Scan(&numero)
		intentos++
		if numero >= n && numero <= m {
			break
		}
	}
	fmt.Println("Acertaste en ", intentos, " intentos")
}

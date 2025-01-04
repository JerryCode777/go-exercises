package main

import "fmt"

func main() {
	var n, m int
	fmt.Print("Ingrese el valor de n: ")
	fmt.Scan(&n)
	fmt.Print("Ingrese el valor de m: ")
	fmt.Scan(&m)

	//sumatoria de los numeros que este en el rango [n,m]
	suma := 0
	for i := n + 1; i < m; i++ {
		suma = suma + i
	}

	fmt.Println("la suma de los numeros entre <n,m> es: ", suma)
}

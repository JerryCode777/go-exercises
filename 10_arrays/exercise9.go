package main

import (
	"fmt"
	"log"
)

func main() {
	var rows, cols int
	fmt.Println("Ingrese las dimensiones de la matriz: ")
	fmt.Print("Numero de filas: ")
	fmt.Scan(&rows)
	fmt.Print("Numero de columnas: ")
	fmt.Scan(&cols)

	A := readMatriz(rows, cols)
	fmt.Println("Matriz A")
	mostrarMatriz(A)

	fmt.Println("Transpuesta de A")
	B := transpuesta(A)
	mostrarMatriz(B)

}

func readMatriz(rows, cols int) [][]int {
	matriz := make([][]int, rows)
	fmt.Println("Ingrese valores: ")
	for i := 0; i < rows; i++ {
		matriz[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			fmt.Printf("[%d][%d]: ", i, j)
			if _, err := fmt.Scan(&matriz[i][j]); err != nil {
				log.Fatal(err)
			}
		}
	}
	return matriz
}

func transpuesta(matriz [][]int) [][]int {
	m := len(matriz)
	n := len(matriz[0])
	transpuesta := make([][]int, n)
	for i := 0; i < n; i++ {
		transpuesta[i] = make([]int, m)
		for j := 0; j < m; j++ {
			transpuesta[i][j] = matriz[j][i]
		}
	}
	return transpuesta
}

func mostrarMatriz(matrix [][]int) {
	for _, rows := range matrix {
		fmt.Println(rows)
	}
}

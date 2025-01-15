package main

import (
	"fmt"
	"log"
)

func main() {
	var rowsA, colsA, rowsB, colsB int
	fmt.Println("Ingrese el tamano de la matriz A: ")
	fmt.Print("filas: ")
	fmt.Scan(&rowsA)
	fmt.Print("columnas: ")
	fmt.Scan(&colsA)
	fmt.Println("Ingrese el tamano de la matriz B: ")
	fmt.Print("filas: ")
	fmt.Scan(&rowsB)
	fmt.Print("columnas: ")
	fmt.Scan(&colsB)

	fmt.Println("Valores de A")
	A := readMatriz(rowsA, colsA)
	fmt.Println("Valores de B")
	B := readMatriz(rowsB, colsB)

	fmt.Println("matriz A: ")
	mostrarMatriz(A)
	fmt.Println("matriz B: ")
	mostrarMatriz(B)

	fmt.Println("=======================================")
	fmt.Println("AxB")
	result, err := multiplicacion(A, B)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Resultado: ")
	mostrarMatriz(result)

}

func readMatriz(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			fmt.Printf("[%d][%d]: ", i, j)
			if _, err := fmt.Scan(&matrix[i][j]); err != nil {
				log.Fatalf("Error al leer el elemento [%d] [%d]: %v", i, j, err)
			}
		}
	}
	return matrix
}

func multiplicacion(A, B [][]int) ([][]int, error) {
	if len(A[0]) != len(B) {
		return nil, fmt.Errorf("El numero de columnas de A debe coincidir con el numero de filas de B")
	}
	rowsA := len(A)
	colsB := len(B[0])
	colsA := len(A[0])
	result := make([][]int, rowsA)
	for i := 0; i < rowsA; i++ {
		result[i] = make([]int, colsB)
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				result[i][j] = result[i][j] + A[i][k]*B[k][j]
			}
		}
	}
	return result, nil
}

func mostrarMatriz(matrix [][]int) {
	for _, rows := range matrix {
		fmt.Println(rows)
	}
}

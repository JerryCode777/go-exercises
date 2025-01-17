package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isValidNumber(num string) bool {
	if len(num) != 9 {
		return false
	}
	for _, ch := range num {
		if !('0' <= ch && ch <= '9') {
			return false
		}
	}
	return true
}

func main() {
	var numbers []string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingresa números de celular (9 dígitos). Escribe 'q' para salir:")

	for len(numbers) < 100 {
		fmt.Print("Número: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		if input == "q" {
			break
		}

		if isValidNumber(input) {
			numbers = append(numbers, input)
		} else {
			fmt.Println("Número inválido. Debe tener 9 dígitos y solo contener números.")
		}
	}

	fmt.Println("Números ingresados:")
	for _, number := range numbers {
		fmt.Println(number)
	}
}

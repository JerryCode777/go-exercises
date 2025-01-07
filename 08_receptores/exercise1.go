package main

import "fmt"

// mayor2 devuelve el mayor de dos numeros reales
func mayor2(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func mayor8(a, b, c, d, e, f, g, h float64) float64 {
	return mayor2(a, mayor2(b, mayor2(c, mayor2(d, mayor2(e, mayor2(f, mayor2(g, h)))))))
}

func main() {
	var a, b, c, d, e, f, g, h float64
	fmt.Print("Ingrese el valor de a: ")
	fmt.Scan(&a)
	fmt.Print("Ingrese el valor de b: ")
	fmt.Scan(&b)
	fmt.Print("Ingrese el valor de c: ")
	fmt.Scan(&c)
	fmt.Print("Ingrese el valor de d: ")
	fmt.Scan(&d)
	fmt.Print("Ingrese el valor de e: ")
	fmt.Scan(&e)
	fmt.Print("Ingrese el valor de f: ")
	fmt.Scan(&f)
	fmt.Print("Ingrese el valor de g: ")
	fmt.Scan(&g)
	fmt.Print("Ingrese el valor de h: ")
	fmt.Scan(&h)
	mayor := mayor8(a, b, c, d, e, f, g, h)

	fmt.Println("El mayor de los 8 numeros reales es: ", mayor)

}

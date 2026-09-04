package main

import "fmt"

func Calcular(a, b float64, operador string) float64 {
	switch operador {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			fmt.Println("No se puede dividir entre cero")
			return 0
		}
		return a / b
	default:
		fmt.Println("Operador no válido")
		return 0
	}
}

func minicalculadora() {
	fmt.Println("10 + 5 =", Calcular(10, 5, "+"))
	fmt.Println("10 / 0 =", Calcular(10, 0, "/"))
}
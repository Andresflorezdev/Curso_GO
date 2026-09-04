package main

import (
	"fmt"
)

func calculadora() {
	for {
		fmt.Println("\n=== CALCULADORA EN GO ===")
		fmt.Println("1. Suma (+)")
		fmt.Println("2. Resta (-)")
		fmt.Println("3. Multiplicación (*)")
		fmt.Println("4. División (/)")
		fmt.Println("5. Salir")
		fmt.Print("Seleccione una opción (1-5): ")

		var opcion int
		fmt.Scanln(&opcion)

		if opcion == 5 {
			fmt.Println("¡Hasta luego!")
			break
		}

		if opcion < 1 || opcion > 5 {
			fmt.Println("Opción no válida. Intente de nuevo.")
			continue
		}

		// Solicitar la cantidad de números a operar
		var cantidad int
		for {
			fmt.Print("¿Cuántos números desea ingresar? (Mínimo 2): ")
			fmt.Scanln(&cantidad)
			if cantidad >= 2 {
				break
			}
			fmt.Println("Debe ingresar al menos 2 números.")
		}

		// Leer los números
		numeros := make([]float64, cantidad)
		for i := 0; i < cantidad; i++ {
			fmt.Printf("Ingrese el número %d: ", i+1)
			fmt.Scanln(&numeros[i])
		}

		// Procesar la operación seleccionada
		resultado := numeros[0]
		errorDivisao := false

		switch opcion {
		case 1:
			for _, num := range numeros[1:] {
				resultado += num
			}
			fmt.Printf("\nResultado de la suma: %.2f\n", resultado)

		case 2:
			for _, num := range numeros[1:] {
				resultado -= num
			}
			fmt.Printf("\nResultado de la resta: %.2f\n", resultado)

		case 3:
			for _, num := range numeros[1:] {
				resultado *= num
			}
			fmt.Printf("\nResultado de la multiplicación: %.2f\n", resultado)

		case 4:
			for _, num := range numeros[1:] {
				if num == 0 {
					errorDivisao = true
					break
				}
				resultado /= num
			}
			if errorDivisao {
				fmt.Println("\nError: No se puede dividir entre cero.")
			} else {
				fmt.Printf("\nResultado de la división: %.2f\n", resultado)
			}
		}
	}
}
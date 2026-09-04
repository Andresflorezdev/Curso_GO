package main

import "fmt"
func repeticiones()  {
	/*var suma int = 0;
	for i:=0; i<=100; i++{
		if i %2 != 0 {
			suma = suma + i
		}
	}
	fmt.Println("la suma fue", suma) */

	miMapa := map[string]string {
		"medellin":"antioquia",
		"bogota":"cundinamarca",
		"manizales":"caldas",
	}
	for ciudad, departamento := range miMapa {
		fmt.Printf("Ciudad: %s, Departamento: %s\n", ciudad, departamento)
	}
}

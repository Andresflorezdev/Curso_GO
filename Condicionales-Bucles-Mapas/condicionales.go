package main

import "fmt"
func condicionales()  {
	var edad int = 39;
	if edad > 65 {
		fmt.Println("Estas jubilado")
	} else if edad >= 18 {
		fmt.Println("la persona esta activa")
	} else {
		fmt.Println("Eres menor")
	}
	fmt.Println("finalizando programa")
}
package ejercicios

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func TablaMultiplicarConsola (){
	
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingresa el numero con el que vas a construir una tabla de multiplicar")
	if scanner.Scan(){
		numero, err := strconv.Atoi(scanner.Text())
		if err != nil{
			fmt.Println("El dato ingresado no es un numero",err.Error())
		}
		fmt.Println()
		fmt.Println("******** Tabla de multiplicar del",numero, "********")
		fmt.Println()
		for i := 1; i <= 10; i++ {
			fmt.Println(numero, "x", i, "=", numero*i)
		}
	}
}

func TablaMultiplicar ()string{
	var texto string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingresa el numero con el que vas a construir una tabla de multiplicar")
	if scanner.Scan(){
		numero, err := strconv.Atoi(scanner.Text())
		if err != nil{
			fmt.Println("El dato ingresado no es un numero",err.Error())
		}
		fmt.Println()
		fmt.Println("******** Tabla de multiplicar del",numero, "********")
		fmt.Println()
		for i := 1; i <= 10; i++ {
			texto +=fmt.Sprintln(numero, "x", i, "=", numero*i)
		}
	}
	return texto
}
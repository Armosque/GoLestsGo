package teclado

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresarNumeros(){
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese un numero:")
	
	if scanner.Scan(){
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil{
			fmt.Println("El dato ingresado no es un numero",err.Error())
		}
	}
	fmt.Println("Ingrese otro numero:")
	if scanner.Scan(){
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil{
			fmt.Println("El dato ingresado no es un numero",err.Error())
		}
	}
	
	fmt.Println("Ingrese una leyenda:")
	if scanner.Scan(){
		leyenda = scanner.Text()
	}
	fmt.Println(leyenda, numero1*numero2)
}
	

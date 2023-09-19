package middleweres

import (
"fmt"
)

func sumar(a int, b int) int {
	return a + b
}
func restar(a int, b int) int {
	return a - b
}
func multiplicar(a int, b int) int{
	return a*b
}

func MiMiddlewere(){
	fmt.Println("Inicio")

	result:= operacionesMidd(sumar)(1,2)
	fmt.Println(result)

	result= operacionesMidd(restar)(10,4)
	fmt.Println(result)

	result= operacionesMidd(multiplicar)(7,9)
	fmt.Println(result)
}

func operacionesMidd(f func(int,int) int) func(int, int) int {
	return func (a int, b int ) int{
		fmt.Println("Inicio de Operaciones")
		return f(a, b)
	}
}
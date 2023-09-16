package arreglos_slices

import (
	"fmt"
)

var tabla [10]int = [10]int{20, 5, 19}
var matriz [20][30]int
func MostrarArreglos(){
	tabla[7] = 9
	tabla[8] = 10

	tabla2 :=[10]string{"Uno", "Dos", "Tres"}

	fmt.Print(tabla)
	fmt.Println(tabla2)

	for i:=0; i<len(tabla); i++{
		
		fmt.Println(tabla[i])
	}
	for i:=0; i<len(tabla); i++{
		
		fmt.Println(tabla2[i])
	}
	matriz[7][24]= 15
	fmt.Println(matriz)
}
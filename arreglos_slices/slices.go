package arreglos_slices

import (
	"fmt"
)

var tablaS []int = []int{20, 5, 19}
var arreglo [10]int = [10]int{4,8,14,23,35,46,57,68,71,77}
func MostrarSlices(){
	fmt.Println(tablaS)

	porcion:=arreglo[3:]
	fmt.Println(porcion)
	porcion2:=arreglo[:6] // No incluye 
	fmt.Println(porcion2)
	porcion3:=arreglo[1:6] // No incluye 
	fmt.Println(porcion3)

}

func Capacidad(){
	elementos:= make([]int, 5, 20)
	fmt.Printf("Capacidad: %d\nLongitud: %d\n", cap(elementos), len(elementos))

	nums:= make([]int, 0,0)
	for i:=0; i<100; i++{
		nums = append(nums, i)
	}
	fmt.Printf("Capacidad: %d\nLongitud: %d\n", cap(nums), len(nums))
}
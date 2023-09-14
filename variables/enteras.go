package variables

import "fmt"

func MostrarEnteros(){
	var intcomun int
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("intcomun es = ", intcomun)
	fmt.Println("intde32 es = ", intde32)
	fmt.Println("intde64 es = ", intde64)
}
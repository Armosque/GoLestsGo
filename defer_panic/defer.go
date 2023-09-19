package defer_panic

import (
	
	
	"fmt"

)

func Defer(){
	fmt.Println("Vamos a hacer un defer. Primer mensaje")
	defer fmt.Println("Este es un mensaje final gracias a defer")

	fmt.Println("Segundo mensaje")
}
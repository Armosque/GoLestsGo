package ejerc_interfaces

import (
	"fmt"
	"github.com/armosque/Godesde0/interfaces"
)

func HumanosRespirando(h interfaces.Humano){
	h.Respirar()
	fmt.Printf("Soy un/a %s y estoy respirando\n", h.Genero())

}
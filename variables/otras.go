package variables

import (
	"fmt"
	"time"
	"strconv"
)

var Nombre string
var Sueldo float32
var Estado bool
var Fecha time.Time
func OtrasVariables(){
	Nombre = "Juan"
	Sueldo = 1000.0
	Estado = true
	Fecha = time.Now()

	fmt.Println(Nombre, Sueldo, Estado, Fecha)
}

func ConvertirATexto(numero int)(bool,string){
	
	texto := strconv.Itoa(numero)
	return true, texto

}
package goroutines

import (
	"fmt"
	"time"
	"strings"

)

func MiNombreLento(nombre string, canal1 chan bool){
	letras:= strings.Split(nombre, "")
	for _,letra := range letras{
		time.Sleep(100*time.Millisecond)
		fmt.Println(letra)
	}
	canal1 <- true
}
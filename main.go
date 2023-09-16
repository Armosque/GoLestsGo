package main

import (
	"github.com/armosque/Godesde0/files"
	//"github.com/armosque/Godesde0/iteraciones"
	//"github.com/armosque/Godesde0/teclado"
	//"github.com/armosque/Godesde0/variables"
	//"github.com/armosque/Godesde0/ejercicios"
	//"fmt"
	//"runtime"
)

func main (){
	/*estado, text := variables.ConvertirATexto(1550)
	fmt.Println(estado)
	fmt.Println(text)

	os := runtime.GOOS
	if os == "windows" {
		fmt.Println("Esto es Windows")
	} else {
		fmt.Println("Esto es Linux")
	}

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Windows")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s", os)
		
	}
	numero, texto := ejercicios.ConvertirANumero("500")
	fmt.Println(numero)
	fmt.Println(texto)

	teclado.IngresarNumeros()

	iteraciones.Iterar()*/

	//fmt.Println(ejercicios.TablaMultiplicar())

	//files.GrabarTabla()

	//files.AgregarTablas()

	files.LeerArchivo()
}

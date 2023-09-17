package main

import (
	e "github.com/armosque/Godesde0/ejerc_interfaces"
	"github.com/armosque/Godesde0/modelos"
	
	//"github.com/armosque/Godesde0/users"
	//"github.com/armosque/Godesde0/arreglos_slices"
	//"github.com/armosque/Godesde0/mapas"
	//"github.com/armosque/Godesde0/funciones"
	//"github.com/armosque/Godesde0/files"
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

	//files.LeerArchivo()

	//funciones.Calculos()

	//funciones.LlamarClosure()

	//funciones.Exponencial(20)

	//arreglos_slices.MostrarArreglos()

	//arreglos_slices.Capacidad()

	//mapas.MostrarMapas()

	//users.NuevoUsuario()
	
	Pedro:=new(modelos.Hombre)
	e.HumanosRespirando(Pedro)

	Ana:=new(modelos.Mujer)
	e.HumanosRespirando(Ana)
}

package files

import (
	"bufio"
	"os"

	"github.com/armosque/Godesde0/ejercicios"

	//"ioutil"
	"fmt"
)

var fileName string = "./files/txt/tabla.txt"
func GrabarTabla(){
	var texto string = ejercicios.TablaMultiplicar()
	
	archivo, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Error al crear el archivo", err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}


func AgregarTablas(){
	var texto string = ejercicios.TablaMultiplicar()
	if !Append(fileName,texto){
		fmt.Println("Error al agregar la tabla")
	}
}

func Append(fileName string, texto string) bool{
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el Append", err.Error())
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el WriteString", err.Error())
		return false
	}
	arch.Close()
	return true
	
}

func LeerArchivo(){
	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error al leer el archivo", err.Error())
		return
	}
	scanner:= bufio.NewScanner(archivo)
	for scanner.Scan(){
		fmt.Println(">",scanner.Text())
	}
	archivo.Close()
}

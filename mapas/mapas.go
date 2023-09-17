package mapas

import "fmt"


func MostrarMapas(){
	paises:= make(map[string]string)
	fmt.Println(paises)

	paises["Argentina"]="Buenos Aires"
	paises["Colombia"]="Bogota"
	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	campeonato:= map[string]int{
		"Barcelona": 39,
		"Real Madrid": 38,
		"Valencia": 37,
		"Sevilla": 36,
	}
	fmt.Println(campeonato)

	/*for equipo, puntaje:= range campeonato{
		fmt.Println(equipo, puntaje)
		fmt.Printf("El equipo %s tiene un puntaje de %d\n", equipo, puntaje)
	
	}*/
	delete(campeonato, "Sevilla")
	
	_,existe:= campeonato["Valencia"]
	fmt.Printf("El equipo se llama %s y existe %t\n","Valencia",existe)

	puntaje,existe:= campeonato["Juventus"]
	fmt.Printf("El puntaje de %s es %d y existe %t\n", "Juventus", puntaje, existe)
}
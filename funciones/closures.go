package funciones

import "fmt"

func tabla(valor int)func()int{
	tabladel:=2
	secuencia :=0
	return func()int{
		secuencia++
		return tabladel*secuencia
	}
}

func LlamarClosure(){
	tabladel:=2
	MiTabla:=tabla(tabladel)

	for i:=0;i<10;i++{
		fmt.Println(MiTabla())
	}
}
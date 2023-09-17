package users

import (
	"time"
	"fmt"
	"github.com/armosque/Godesde0/modelos"
)

func NuevoUsuario(){
	usuario:= new (modelos.User)
	usuario.AddUser(1, "Juan", time.Now(), true)
	fmt.Println(usuario)
}
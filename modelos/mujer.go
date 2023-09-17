package modelos

type Mujer struct {
	Hombre
	EsMadre bool

}

func (mujer *Mujer) Respirar() {mujer.respirarando = true}
func (mujer *Mujer) Comer() {mujer.comiendo = true}
func (mujer *Mujer) Pensar() {mujer.pensando = true}
func (mujer *Mujer) Genero() string {return "Mujer"}

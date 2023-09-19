package modelos

type Hombre struct {
	edad         int
	altura       float32
	peso         float32
	respirarando bool
	comiendo     bool
	pensando     bool
	vivo         bool
}

// EstaVivo implements interfaces.Humano.


func (hombre *Hombre) Respirar()      { hombre.respirarando = true }
func (hombre *Hombre) Comer()         { hombre.comiendo = true }
func (hombre *Hombre) Pensar()        { hombre.pensando = true }
func (hombre *Hombre) Genero() string { return "Hombre" }

package interfaces

type animal interface {
	Comer()
	Respirar()
	EsCarnivoro() bool
	EstaVivo() bool
}
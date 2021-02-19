package abstractFactory

type factory interface {
	CreatePC() computer
	CreatePad() pad
}
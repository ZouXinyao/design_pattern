package flyweight

type dress interface {
	GetColor() string
}

type RedDress struct {
	color string
}

func (r *RedDress) GetColor() string {
	return r.color
}

func NewRedDress() *RedDress {
	return &RedDress{color: "Red"}
}

type BlackDress struct {
	color string
}

func (b *BlackDress) GetColor() string {
	return b.color
}

func NewBlackDress() *BlackDress {
	return &BlackDress{color: "Black"}
}
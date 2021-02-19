package abstractFactory

// DellFactory dell工厂
type DellFactory struct {
}

// CreatePC 创建dellPc
func (fac *DellFactory) CreatePC() *dellComputer {
	return &dellComputer{brand: "Dell"}
}

func (fac *DellFactory) CreatePad() *dellPad {
	return &dellPad{brand: "Dell"}
}

// LenovoFactory Lenovo工厂
type LenovoFactory struct {
}

// CreatePC 创建LenovoPc
func (fac *LenovoFactory) CreatePC() *lenovoComputer {
	return &lenovoComputer{brand: "Lenovo"}
}

func (fac *LenovoFactory) CreatePad() *lenovoPad {
	return &lenovoPad{brand: "Lenovo"}
}

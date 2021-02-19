package simplefactory

// CreatePCFromSimple 简单工厂创建对象
func CreatePCFromSimple(brand string) computer {
	switch brand {
	case "Dell":
		return &dellComputer{brand: brand}
	case "Lenovo":
		return &lenovoComputer{brand: brand}
	}
	return nil
}
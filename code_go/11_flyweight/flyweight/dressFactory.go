package flyweight

const (
	//TerroristDressType terrorist dress type
	RedType = "Red"
	//CounterTerrroristDressType terrorist dress type
	BlackType = "Black"
)

type dressFactory struct {
	dressMap map[string]dress
}

var (
	dressFactorySingleInstance = &dressFactory{
		dressMap: map[string]dress{
			RedType: NewRedDress(),
			BlackType: NewBlackDress(),
		},
	}
)

func (d *dressFactory) GetDress(dressType string) dress {
	return d.dressMap[dressType]
}

func newDressFactory() *dressFactory {
	return dressFactorySingleInstance
}

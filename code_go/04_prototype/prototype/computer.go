package prototype

type computer struct {
	brand string
	screen string
	Pc *mainframe
}

func (c *computer) SetBrand(brand string) {
	c.brand = brand
}

func (c *computer) GetBrand() string {
	return c.brand
}

func (c *computer) SetScreen(screenBrand string) {
	c.screen = screenBrand
}

func (c *computer) GetScreen() string {
	return c.screen
}

func (c *computer) SetMainframe(pcBrand *mainframe) {
	c.Pc = pcBrand
}

func (c *computer) GetMainframe() *mainframe {
	return c.Pc
}

func (c *computer) Clone() *computer {
	newComputer := computer{}
	newComputer.screen = c.screen
	newComputer.Pc = c.Pc.clone()
	return &newComputer
}

func NewComputer(brand string) *computer {
	return &computer{brand: brand}
}
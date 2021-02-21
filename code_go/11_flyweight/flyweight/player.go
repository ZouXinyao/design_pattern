package flyweight

type Player struct {
	dress      dress
	playerName string
	gun        string
}

func NewPlayer(dressType string, playerName string) *Player {
	dress := newDressFactory().GetDress(dressType)
	return &Player{
		playerName: playerName,
		dress:      dress,
	}
}

func (p *Player) SetGun(gun string) {
	p.gun = gun
}

func (p *Player) GetGun() string {
	return p.gun
}

func (p *Player) GetDressColor() string {
	return p.dress.GetColor()
}

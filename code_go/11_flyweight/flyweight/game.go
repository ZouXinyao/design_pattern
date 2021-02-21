package flyweight

import (
	"fmt"
	"strings"
)

type Game struct {
	RedRable []*Player
	BlackRable []*Player
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Add(p *Player) {
	if strings.EqualFold(p.dress.GetColor(), RedType) {
		g.RedRable = append(g.RedRable, p)
	} else {
		g.BlackRable = append(g.BlackRable, p)
	}
}

func (g *Game) GetAllRedMember() {
	for _, player := range g.RedRable {
		fmt.Println(player.playerName)
	}
}

func (g *Game) GetAllBlackMember() {
	for _, player := range g.BlackRable {
		fmt.Println(player.playerName)
	}
}


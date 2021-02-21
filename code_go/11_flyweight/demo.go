package main

import (
	"github.com/ZouXinyao/design_pattern/code_go/11_flyweight/flyweight"
)

func main() {
	game := flyweight.NewGame()

	game.Add(flyweight.NewPlayer(flyweight.RedType, "redPlayer01"))
	game.Add(flyweight.NewPlayer(flyweight.RedType, "redPlayer02"))
	game.Add(flyweight.NewPlayer(flyweight.BlackType, "BlackPlayer01"))
	game.Add(flyweight.NewPlayer(flyweight.BlackType, "BlackPlayer02"))

	game.GetAllBlackMember()
	game.GetAllRedMember()
}
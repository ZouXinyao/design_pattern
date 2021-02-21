package main

import "github.com/ZouXinyao/design_pattern/code_go/16_state/state"

func main() {
	player01 := state.NewRole()
	player01.GetFlower()
	player01.GetFlower()
	player01.GetMushroom()
	player01.UnderLife()
	player01.GetFlower()
	player01.ShowInfo()

	player02 := state.NewRole()
	player02.GetMushroom()
	player02.GetFlower()
	player02.UnderLife()
	player02.UnderLife()
	player02.UnderLife()
	player02.ShowInfo()



}

package state

import "fmt"

type Role struct {
	currentState State
	life int
	score int
	haveGun bool
}

func NewRole() *Role {
	return &Role{
		currentState: NewChildState(),
		life: 1,
		score: 0,
		haveGun: false,
	}
}

func (r * Role) setState(s State) {
	r.currentState = s
}

func (r *Role) GetFlower() {
	r.currentState.GetFlower(r)
}

func (r *Role) GetMushroom() {
	r.currentState.GetMushroom(r)
}

func (r *Role) UnderLife() {
	r.currentState.UnderLife(r)
}

func (r *Role) ShowInfo() {
	fmt.Println("life: ", r.life)
	fmt.Println("score: ", r.score)
}
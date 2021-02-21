package state

import "fmt"

type State interface {
	GetFlower(h *Role)
	GetMushroom(h *Role)
	UnderLife(h *Role)
}

type ChildState struct {
}

func NewChildState() State {
	return &ChildState{}
}

func (c *ChildState) GetFlower(h *Role) {
	h.score += 1000
}

func (c *ChildState) GetMushroom(h *Role) {
	h.score += 2000
	h.life++
	h.setState(NewHumanState())
}

func (c *ChildState) UnderLife(h *Role) {
	h.score = 0
	h.life--
	fmt.Println("You have died")
}

type HumanState struct {
}

func NewHumanState() State {
	return &HumanState{}
}

func (m *HumanState) GetFlower(h *Role) {
	h.score += 1000
	h.life++
	h.haveGun = true
	h.setState(NewSupermanState())
}

func (m *HumanState) GetMushroom(h *Role) {
	h.score += 1000
}

func (m *HumanState) UnderLife(h *Role) {
	h.life--
	h.setState(NewChildState())
}

type SupermanState struct {
}

func NewSupermanState() State {
	return &SupermanState{}
}

func (m *SupermanState) GetFlower(h *Role) {
	h.score += 1000
}

func (m *SupermanState) GetMushroom(h *Role) {
	h.score += 1000
}

func (m *SupermanState) UnderLife(h *Role) {
	h.life--
	h.haveGun = false
	h.setState(NewHumanState())
}
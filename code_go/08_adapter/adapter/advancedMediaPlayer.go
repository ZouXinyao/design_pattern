package adapter

import "fmt"

type AdvancedMediaPlayer interface {
	PlayWMA(fileName string)
}

type WMAPlayer struct {
}

func (WMA *WMAPlayer) PlayWMA(fileName string) {
	fmt.Println("Playing WMA file. File name is " + fileName)
}

func NewWMAPlayer() *WMAPlayer {
	return &WMAPlayer{}
}
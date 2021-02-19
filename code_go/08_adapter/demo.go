package main

import (
	"github.com/ZouXinyao/design_pattern/code_go/08_adapter/adapter"
)

func main() {
	player := adapter.NewAudioPlayer()

	player.Play("mp3", "song01.mp3")
	player.Play("WMA", "song02.WMA")
	player.Play("mp4", "song03.mp4")

}
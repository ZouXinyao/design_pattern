package adapter

import (
	"fmt"
	"strings"
)

type MediaPlayer interface {
	Play(audioType string, fileName string)
}

type AudioPlayer struct {
	*MediaAdapter
}

func NewAudioPlayer() *AudioPlayer {
	return &AudioPlayer{}
}

func (player *AudioPlayer) Play(audioType string, fileName string) {
	if strings.EqualFold(audioType, "mp3") {
		fmt.Println("Playing mp3 file. File name is " + fileName)
	} else if strings.EqualFold(audioType, "WMA") {
		player.MediaAdapter = NewMediaAdapter(audioType)
		player.MediaAdapter.Play(audioType, fileName)
	} else {
		fmt.Println("Invalid media. " + audioType + " format not supported")
	}
}
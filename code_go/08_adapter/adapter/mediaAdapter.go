package adapter

import "strings"

type MediaAdapter struct {
	AdvancedMediaPlayer
}

func NewMediaAdapter(audioType string) *MediaAdapter {
	if strings.EqualFold(audioType, "WMA") {
		return &MediaAdapter{AdvancedMediaPlayer: NewWMAPlayer()}
	}
	return &MediaAdapter{}
}

func (adapt *MediaAdapter) Play(audioType string, fileName string) {
	adapt.PlayWMA(fileName)
}
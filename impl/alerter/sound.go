package alerter

import (
	"os/exec"
)

type SoundPlayer struct {
	soundPath string
}

func NewSoundPlayer(soundPath string) *SoundPlayer {
	return &SoundPlayer{soundPath: soundPath}
}

func (p *SoundPlayer) Alert(message string) error {
	cmd := exec.Command("bash", "-c", "afplay", p.soundPath)
	return cmd.Run()
}

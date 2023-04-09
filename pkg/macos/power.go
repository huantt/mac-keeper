package macos

import (
	"bytes"
	"os/exec"
	"strings"
)

const checkPowerSourceCommand = "pmset -g batt | head -n 1 | cut -d \\' -f2"
const ACPowerSource = "AC Power"
const BatteryPower = "Battery Power"

func GetPowerSource() (*string, error) {
	cmd := exec.Command("bash", "-c", checkPowerSourceCommand)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	powerSource := out.String()
	powerSource = strings.TrimSpace(powerSource)
	return &powerSource, nil
}

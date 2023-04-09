package macos

import (
	"fmt"
	"testing"
)

func TestGetPowerSource(t *testing.T) {
	if powerSource, err := GetPowerSource(); err != nil {
		t.Error(err)
	} else {
		fmt.Println(*powerSource)
	}
}

package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"mac-keeper/cmd"
)

func main() {
	command := cobra.Command{}
	command.AddCommand(cmd.MonitorPowerCommand())
	err := command.Execute()
	if err != nil {
		fmt.Println(err)
	}

}

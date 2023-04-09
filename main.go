package main

import (
	"fmt"
	"github.com/huantt/mac-keeper/cmd"
	"github.com/spf13/cobra"
)

func main() {
	command := cobra.Command{}
	command.AddCommand(cmd.MonitorPowerCommand())
	err := command.Execute()
	if err != nil {
		fmt.Println(err)
	}

}

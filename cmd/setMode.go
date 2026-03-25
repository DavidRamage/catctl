/*
Copyright © 2025 David Ramage
*/
package cmd

import (
	"catctl/catfunctions"
	"fmt"

	"github.com/spf13/cobra"
)

// setModeCmd represents the setMode command
var setModeCmd = &cobra.Command{
	Use:   "mode",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		serial, radio := catfunctions.GetConf()
		command := catfunctions.GetRadioData(radio, "commands", "setmode")
		bandcode := catfunctions.GetRadioData(radio, "bandtable", args[0])
		out := catfunctions.SendCommand(serial, command+bandcode)
		fmt.Println(out)
	},
}

func init() {
	setCmd.AddCommand(setModeCmd)
}

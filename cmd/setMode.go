/*
Copyright © 2025 David Ramage
*/
package cmd

import (
	"catctl/catfunctions"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// setModeCmd represents the setMode command
var setModeCmd = &cobra.Command{
	Use:   "mode",
	Short: "Sets the mode of the radio",
	Long:  `catctl set mode lsb`,
	Run: func(cmd *cobra.Command, args []string) {
		serial, radio, err := catfunctions.GetConf()
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(-1)
		}
		command, err := catfunctions.GetRadioData(radio, "commands", "setmode")
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(-1)
		}
		bandcode, err := catfunctions.GetRadioData(radio, "bandtable", args[0])
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(-1)
		}
		out, err := catfunctions.SendCommand(serial, command+bandcode)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(-1)
		}
		fmt.Println(out)
	},
}

func init() {
	setCmd.AddCommand(setModeCmd)
}

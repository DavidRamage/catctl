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

// setVfoCmd represents the setVfo command
// this does not work yet
var setVfoCmd = &cobra.Command{
	Use:   "catctl set vfo",
	Short: "Set the current variable frequency oscillator (VFO)",
	Long:  `catctl set vfo 0`,
	Run: func(cmd *cobra.Command, args []string) {
		serial, radio, err := catfunctions.GetConf()
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(-1)
		}
		command, err := catfunctions.GetRadioData(radio, "commands", "setvfo")
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(-1)
		}
		out, err := catfunctions.SendCommand(serial, command)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(-1)
		}
		fmt.Println(out)

	},
}

func init() {
	setCmd.AddCommand(setVfoCmd)
}

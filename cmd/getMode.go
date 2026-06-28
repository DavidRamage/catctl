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

// getModeCmd represents the getMode command
var getModeCmd = &cobra.Command{
	Use:   "catctl get mode",
	Short: "Get the current mode",
	Long: `catctl get mode
LSB`,
	Run: func(cmd *cobra.Command, args []string) {
		serial, radio, err := catfunctions.GetConf()
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(-1)
		}
		command, err := catfunctions.GetRadioData(radio, "commands", "getmode")
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(-1)
		}
		modeKey, err := catfunctions.SendCommand(serial, command)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(-1)
		}
		mode, err := catfunctions.GetRadioData(radio, "modetable", modeKey)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(-1)
		}
		fmt.Println("Mode:", mode)
	},
}

func init() {
	getCmd.AddCommand(getModeCmd)
}

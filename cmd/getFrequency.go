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

// getFrequencyCmd represents the getFrequency command
// We need to mess with this because get frequency is per vfo
var getFrequencyCmd = &cobra.Command{
	Use:   "frequency",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		serial, radio, err := catfunctions.GetConf()
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(-1)
		}

		command, err := catfunctions.GetRadioData(radio, "commands", "getfreqvfo"+vfo)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(-1)
		}
		cmdOut, err := catfunctions.SendCommand(serial, command)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(-1)
		} else {
			fmt.Println(cmdOut)
		}
	},
}

func init() {
	getCmd.AddCommand(getFrequencyCmd)
}

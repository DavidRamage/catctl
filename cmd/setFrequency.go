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

// setFrequencyCmd represents the setFrequency command
// we need to split this out per vfo

var setFrequencyCmd = &cobra.Command{
	Use:   "frequency",
	Short: "Sets the frequency of the current VFO",
	Long:  `catctl set frequency 7.2m`,
	Run: func(cmd *cobra.Command, args []string) {
		serial, radio, err := catfunctions.GetConf()
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(-1)
		}
		command, err := catfunctions.GetRadioData(radio, "commands", "setfreqvfo"+vfo)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(-1)
		}
		hz, err := catfunctions.ToHz(args[0])
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(-1)
		}
		out, err := catfunctions.SendCommand(serial, command+hz)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(-1)
		}
		fmt.Println(out)
	},
}

func init() {
	setCmd.AddCommand(setFrequencyCmd)
}

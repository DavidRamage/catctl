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

// not natively supported by the CAT protocol, but can be used to get the band from the radio and then use that to determine which commands to send for other functions like getmode, getvfo, etc.
// getBandCmd represents the getBand command
var getBandCmd = &cobra.Command{
	Use:   "catctl get band",
	Short: "Get the current band the radio is using",
	Long: `catctl get band
80m`,
	Run: func(cmd *cobra.Command, args []string) {
		serial, radio, err := catfunctions.GetConf()
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(-1)
		}
		command, err := catfunctions.GetRadioData(radio, "commands", "getband")
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
	getCmd.AddCommand(getBandCmd)
}

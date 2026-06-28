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

// getVfoCmd represents the getVfo command
var getVfoCmd = &cobra.Command{
	Use:   "vfo",
	Short: "Get the current VFO",
	Long: `catctl get vfo
VS0;.`,
	Run: func(cmd *cobra.Command, args []string) {
		serial, radio, err := catfunctions.GetConf()
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(-1)
		}
		command, err := catfunctions.GetRadioData(radio, "commands", "getvfo")
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
	getCmd.AddCommand(getVfoCmd)
}

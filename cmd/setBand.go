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

// setBandCmd represents the setBand command
var setBandCmd = &cobra.Command{
	Use:   "band",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		serial, radio := catfunctions.GetConf()
		commandPrefix := catfunctions.GetRadioData(radio, "commands", "setband")
		bandcode := catfunctions.GetRadioData(radio, "bandtable", args[0])
		out, err := catfunctions.SendCommand(serial, commandPrefix+bandcode)
		if err != nil {
			os.Exit(-1)
		} else {
			fmt.Println(out)
		}
	},
}

func init() {
	setCmd.AddCommand(setBandCmd)
}

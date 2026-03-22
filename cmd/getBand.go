/*
Copyright © 2025 David Ramage
*/
package cmd

import (
	"catctl/catfunctions"
	"fmt"

	"github.com/spf13/cobra"
)

// not natively supported by the CAT protocol, but can be used to get the band from the radio and then use that to determine which commands to send for other functions like getmode, getvfo, etc.
// getBandCmd represents the getBand command
var getBandCmd = &cobra.Command{
	Use:   "band",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		serial, radio := catfunctions.GetConf()
		command := catfunctions.GetCommand(radio, "getband")
		fmt.Println(command)
		cmdOut := catfunctions.SendCommand(serial, command)
		fmt.Println(cmdOut)
	},
}

func init() {
	getCmd.AddCommand(getBandCmd)
	//fmt.Println("Command to send:", command)
	//_, radio := catfunctions.GetConf()
	//command := catfunctions.GetCommand(radio, "getBand")
}

/*
Copyright © 2025 David Ramage
*/
package cmd

import (
	"catctl/catfunctions"
	"fmt"

	"github.com/spf13/cobra"
)

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
		fmt.Println("getBand called")
		_, radio := catfunctions.GetConf()
		fmt.Println(radio)
		command := catfunctions.GetCommand(radio, "setBand")
		fmt.Println(command)
	},
}

func init() {
	getCmd.AddCommand(getBandCmd)
	//fmt.Println("Command to send:", command)
	//_, radio := catfunctions.GetConf()
	//command := catfunctions.GetCommand(radio, "getBand")
}

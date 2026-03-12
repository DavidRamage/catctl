/*
Copyright © 2025 David Ramage 
*/
package cmd

import (
	"catctl/catfunctions"
	"fmt"

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
		_, radio := catfunctions.GetConf()
		commandPrefix := catfunctions.GetCommand(radio, "setband")
		command := commandPrefix + " " + args[0]
		fmt.Println(command)
	},
}

func init() {
	setCmd.AddCommand(setBandCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setBandCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setBandCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

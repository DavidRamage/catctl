/*
Copyright © 2025 David Ramage
*/
package cmd

import (
	"catctl/catfunctions"
	"fmt"

	"github.com/spf13/cobra"
)

// setFrequencyCmd represents the setFrequency command
// we need to split this out per vfo
var setFrequencyCmd = &cobra.Command{
	Use:   "frequency",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		_, radio := catfunctions.GetConf()
		command := catfunctions.GetCommand(radio, "setfrequency")
		fmt.Println(command)
	},
}

func init() {
	setCmd.AddCommand(setFrequencyCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setFrequencyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setFrequencyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

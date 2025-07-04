/*
Copyright © 2025 David Ramage
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// aboutCmd represents the about command
var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "About catctl",
	Long:  "Catctl is a command line tool for interacting with your radio.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("about called")
	},
}

func init() {
	rootCmd.AddCommand(aboutCmd)
	about()
}

func about() {
	about := `Catctl is a command line tool for interacting with your radio.`
	fmt.Println(about)
}

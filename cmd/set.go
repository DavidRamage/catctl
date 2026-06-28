/*
Copyright © 2025 David Ramage
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set various things about the radio",
	Long: `catctl set frequency
catctl set mode
etc.`,
}

func init() {
	rootCmd.AddCommand(setCmd)
}

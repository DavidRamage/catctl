/*
Copyright © 2025 David Ramage
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get status information from the radio",
	Long: `get frequency
catctl get mode
etc.`,
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getBandCmd)
}

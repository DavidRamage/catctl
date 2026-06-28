/*
Copyright © 2025 David Ramage
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "catctl get",
	Short: "Get status information from the radio",
	Long: `catctl get frequency
catctl get mode
etc.`,
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getBandCmd)
}

/*
Copyright © 2025 David Ramage
*/
package cmd

import (
	"catctl/catfunctions"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// getFrequencyCmd represents the getFrequency command
// We need to mess with this because get frequency is per vfo
var getFrequencyCmd = &cobra.Command{
	Use:   "catctl get frequency",
	Short: "Returns the frequency of the current VFO",
	Long: `A longer description that spans multiple lines and likely contains examples
catctl get frequency
7170.90 khz`,
	Run: func(cmd *cobra.Command, args []string) {
		serial, radio, err := catfunctions.GetConf()
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(-1)
		}

		command, err := catfunctions.GetRadioData(radio, "commands", "getfreqvfo"+vfo)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(-1)
		}
		prefixlen_str, err := catfunctions.GetRadioData(radio, "formatting", "prefixlen")
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(-1)
		}
		prefixlen, err := strconv.Atoi(prefixlen_str)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(-1)
		}
		suffixlen_str, err := catfunctions.GetRadioData(radio, "formatting", "suffixlen")
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(-1)
		}
		suffixlen, err := strconv.Atoi(suffixlen_str)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(-1)
		}
		cmdOut, err := catfunctions.SendCommand(serial, command)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(-1)
		} else {
			//TODO: make a preference-specific config file
			freqFmt, err := catfunctions.FmtFrequencyOut(cmdOut, "khz", prefixlen, suffixlen)
			if err != nil {
				fmt.Println("Error: ", err)
				os.Exit(-1)
			}
			fmt.Println(freqFmt)
		}
	},
}

func init() {
	getCmd.AddCommand(getFrequencyCmd)
}

/*
Copyright © 2021 Kyle Nguyen <nvietthu@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// tgCmd represents the tg command
var tgCmd = &cobra.Command{
	Use:   "tg -i [entry price] -d [+- number (%)]",
	Short: "Target - Calculate final value though a delta of Percentage",
	Long: `Target - Calculate final value though a delta of Percentage`,
	Run: func(cmd *cobra.Command, args []string) {
		  entry,_ := cmd.Flags().GetFloat64("in")
		  deltaPercent,_ := cmd.Flags().GetFloat64("delta")
		  exit := ((entry * deltaPercent) / 100) + entry
		  delta := exit - entry
		  color := ""
		  if exit < entry {
			color = "\033[31m"
			fmt.Printf("Entry: %0.2f\nDelta: %0.2f%%\n%sExit: %0.2f\nLossed: %0.2f\n",entry, deltaPercent, color, exit, delta)
		  } else {
			color = "\033[32m"
			fmt.Printf("Entry: %0.2f\nDelta: %0.2f%%\n%sExit: %0.2f\nGained: %0.2f\n",entry, deltaPercent, color, exit, delta)
		  }
	},
}

func init() {
	rootCmd.AddCommand(tgCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tgCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tgCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	tgCmd.Flags().Float64P("in", "i", 0.00, "Entry price")
	tgCmd.Flags().Float64P("delta", "d", 0.00, "Delta +- in percent")
}

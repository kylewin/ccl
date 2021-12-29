/*
Copyright © 2021 Kyle Nguyen <nvietthu@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pcCmd represents the pc command
var pcCmd = &cobra.Command{
	Use:   "pc [entry] [exit]",
	Short: "Percentage - Calculate Percentage",
	Long: `Percentage - Calculate Delta in Percentage`,
	Run: func(cmd *cobra.Command, args []string) {
              entry,_ := cmd.Flags().GetFloat64("in")
              exit,_ := cmd.Flags().GetFloat64("out")
              delta := ((exit - entry) / entry) * 100
              color := ""
              if delta < 0 {
                color = "\033[31m"
              } else {
                color = "\033[32m"
              }
              fmt.Printf("Entry: %0.2f\nExit: %0.2f\n%sChange: %0.2f%%\n",entry, exit, color, delta)
	},
}

func init() {
	rootCmd.AddCommand(pcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	pcCmd.Flags().Float64P("in", "i", 0.0, "Entry price")
	pcCmd.Flags().Float64P("out", "o", 0.0, "Exit price")
}

/*
Copyright © 2021 Kyle Nguyen <nvietthu@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// auCmd represents the au command
var auCmd = &cobra.Command{
	Use:   "au -d [Gold Price in USD] -v [1 usd to vnd]",
	Short: "Gold price",
	Long:  `Gold price from USD/Ounce to VND/L`,
	Run: func(cmd *cobra.Command, args []string) {
		gold_usd, _ := cmd.Flags().GetFloat64("gold_usd")
		usd_to_vnd, _ := cmd.Flags().GetFloat64("usd_to_vnd")
		if usd_to_vnd == 0 {
			usd_to_vnd = 24275
		}

		gold_vnd := (0.0375 * (gold_usd * usd_to_vnd)) / 0.0283
		color := "\033[33m"

		p := message.NewPrinter(language.English)
		p.Printf("1USD: %0.2f VND\n", usd_to_vnd)
		p.Printf("1L: %s%0.2f VND\n", color, gold_vnd)
	},
}

func init() {
	rootCmd.AddCommand(auCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// auCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	auCmd.Flags().Float64P("gold_usd", "d", 0.0, "Gold in USD/Ounce")
	auCmd.Flags().Float64P("usd_to_vnd", "v", 0.0, "1 USD in VND")
}

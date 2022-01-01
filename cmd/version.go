/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	goVersion "go.hein.dev/go-version"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var (
	shortened  = false
	version    = "dev"
	commit     = "none"
	date       = "unknown"
	output     = "json"
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Version will output the current build information",
		Long:  ``,
		Run: func(_ *cobra.Command, _ []string) {
			var response string
			versionOutput := goVersion.New(version, commit, date)
			response = versionOutput.ToShortened()
			fmt.Printf("%+v", response)
			return
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

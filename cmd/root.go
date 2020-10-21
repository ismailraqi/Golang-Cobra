package cmd

import (
	"fmt"
	"gorecover/config"
	"os"

	"github.com/common-nighthawk/go-figure"

	"github.com/spf13/cobra"
)

var logDTFormat string = "Jan 02 2006 03:04:05 PM"

var rootCmd = &cobra.Command{
	Use:     config.AppName,
	Short:   config.ShortDesc,
	Long:    config.LongDesc,
	Version: config.Version,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	println("hello graviton")
	// },
}

//Execute is the first func
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
func init() {
	// Display the app ASCII logo
	myFigure := figure.NewFigure(config.AppDisplayName, "", true)
	myFigure.Print()
}

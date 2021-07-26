package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "gocopy",
	Short:   "A lightweight automated backup file software.",
	Long:    `A lightweight automated backup file software.`,
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		//  // Do Stuff Here
		fmt.Println("Hello Cobra CLI application")
		//Run: func(cmd *cobra.Command, args []string) {fmt.Println("Hello CLI")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

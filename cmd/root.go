/*
Copyright © 2023 Joshua Brewer
*/
package cmd

import (
	"os"

	"github.com/Joshbwr/tclone/pkg/util"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tclone",
	Short: "Clone Locally Stored Starter Templates",
	Long: `Clone Locally Stored Starter Templates`,

	Run: func(cmd *cobra.Command, args []string) { 
		util.Init()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


package cmd

import (
	"github.com/spf13/cobra"
	"github.com/webjaba/ctp/internal/app"
)

var colorsCmd = &cobra.Command{
	Use:   "colors",
	Short: "show possible colors",
	Long:  `this command can show you all possible colors and how they looks`,
	Run:   app.Colors,
}

func init() {
	rootCmd.AddCommand(colorsCmd)
}

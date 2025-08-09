package cmd

import (
	"github.com/spf13/cobra"
	"github.com/webjaba/ctp/internal/app"
)

var long = "example with short name:\n\nctp print some text -c=g\n\033[32msome text\033[0m\n\nyou can use wide name of color\n"

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "colored print, use -c flag to set color",
	Long:  long,
	Run:   app.Print,
}

func init() {
	rootCmd.AddCommand(printCmd)

	printCmd.Flags().StringP("color", "c", "", "choose a color")
}

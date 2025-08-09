package app

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/webjaba/ctp/internal/formatting"
)

func Colors(cmd *cobra.Command, args []string) {
	fmt.Println(formatting.GetColorSet())
}

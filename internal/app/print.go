package app

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/webjaba/ctp/internal/formatting"
)

func Print(cmd *cobra.Command, args []string) {
	colorFlag := cmd.Flag("color")

	res := formatting.Colorize(
		strings.Join(args, " "),
		colorFlag.Value.String(),
	)

	fmt.Println(res)
}

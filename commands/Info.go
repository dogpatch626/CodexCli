package Commands

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:     "info",
	Aliases: []string{"info"},
	Short:   "CodexCli is a in development cli",
	Run: func(cmd *cobra.Command, args []string) {
		color.Red("CodexCli is a in-development cli made in go as a project to learn what is even going on with this lang.")
	},
}

func init() {

	rootCmd.AddCommand(infoCmd)

}

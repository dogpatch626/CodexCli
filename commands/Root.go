package Commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "Codex",
	Short: "Codex is a virus total cli",
	Long:  "Codex is a virus total cli built by dogpatch626 in Go",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("idk")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
}

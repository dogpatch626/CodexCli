package Commands

import (
	"CodexCli/banner"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Codex",
	Short: "Codex is a virus total cli",
	Long:  "Codex is a virus total cli built by dogpatch626 in Go",
	Run: func(cmd *cobra.Command, args []string) {
		banner.Banner()

		ran := exec.Command("./CodexCli", "--help")

		out, err := ran.Output()

		if err != nil {
			fmt.Println("Could not run: --help", err)
		}

		fmt.Println("\n", string(out))

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

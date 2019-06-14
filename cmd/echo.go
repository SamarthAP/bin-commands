package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "Output the ARGS",
	Long:  `Output the ARGS`,
	Run: func(cmd *cobra.Command, args []string) {
		n := cmd.Flag("n")
		if n.Value.String() != "" {
			fmt.Print(n.Value.String()+" ", strings.Join(args, " "))
		} else {
			fmt.Println(strings.Join(args, " "))
		}
	},
}

func init() {
	rootCmd.AddCommand(echoCmd)
	echoCmd.Flags().StringP("n", "n", "", "If -n is specified, the trailing newline is suppressed")
}

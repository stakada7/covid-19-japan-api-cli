package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version string
	Revision string
)

func init()  {
	rootCmd.AddCommand(newVersionCmd())
}

func newVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "version",
		Short: "Print version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("covid-19-japan-api-cli version: %s, revision: %s\n", Version, Revision)
		},
	}

	return cmd
}
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

const VersionNumber = "0.0.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: VersionNumber,
	Long:  VersionNumber,
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("gostj version %s\n", VersionNumber)
	},
}

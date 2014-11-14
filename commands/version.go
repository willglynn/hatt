package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the version number",
	Long:  "Well, sometimes you gotta know what version you've got.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("hatt version 0.1\n")
	},
}

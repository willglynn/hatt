package commands

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "hatt",
	Short: "Hatt is a file hashing tool",
	Long:  `hatt: hash all the things!`,
}

func init() {
	RootCmd.AddCommand(HashCmd)
	RootCmd.AddCommand(VersionCmd)
}

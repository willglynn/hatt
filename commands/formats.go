package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/willglynn/hatt/interop"
)

var FormatsCmd = &cobra.Command{
	Use:   "formats",
	Short: "List supported import/export formats",
	Long:  "hatt supports reading/writing formats written by other tools.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Listing formats supported for Input and Output:\n\n")

		fmt.Printf("%-16s  %2s  %s\n", "Name", "IO", "Description")
		fmt.Printf("%-16s  %2s  %s\n", "----------------", "--", "---------------------------------")
		for _, f := range interop.Formats {
			i, o := "-", "-"
			if _, canRead := f.(interop.Importer); canRead {
				i = "I"
			}
			if _, canWrite := f.(interop.Exporter); canWrite {
				o = "O"
			}

			fmt.Printf("%-16s  %s%s  %s\n", f.Name(), i, o, f.Description())
		}

		fmt.Printf("\n")
	},
}

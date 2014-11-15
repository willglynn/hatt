package commands

import (
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/willglynn/hatt/interop"
	"github.com/willglynn/hatt/manifest"
)

var ExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a manifest to another format",
	Long:  `hatt supports exporting to other formats. (See the 'formats' subcommand for a list.)`,

	Run: export.run,
}

type exportOperation struct {
	manifestPath string
	outputPath   string
	formatName   string

	m *manifest.Manifest
}

var export exportOperation

func (export *exportOperation) run(cmd *cobra.Command, args []string) {
	if export.manifestPath == "" {
		cmd.Usage()
		log.Fatal("manifest file must be specified")
		return
	}

	// find an output format
	var format interop.Format
	if export.formatName != "" {
		format = interop.FindByName(export.formatName)
	} else {
		format = interop.FindByExtension(export.outputPath)
	}

	if format == nil {
		if export.formatName == "" {
			log.Fatalf("specify a format with -f, check `hatt formats` for a list")
		} else {
			log.Fatalf("no such format %q, check `hatt formats`", export.formatName)
		}
	}

	exporter, ok := format.(interop.Exporter)
	if !ok {
		log.Fatalf("%s format does not support exporting", format.Name())
	}

	// open the manifest, ignorning not found errors
	if m, err := manifest.ReadFromFile(export.manifestPath); err != nil {
		log.Fatalf("error reading manifest file %q: %v", export.manifestPath, err)
	} else {
		export.m = m
	}

	// get an output writer
	var w io.Writer
	if export.outputPath == "-" {
		w = os.Stdout
	} else if file, err := os.OpenFile(export.outputPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666); err != nil {
		log.Fatalf("error opening output file: %v", err)
	} else {
		defer file.Close()
		w = file
	}

	// do the export
	if err := exporter.Export(export.m, w); err != nil {
		log.Fatalf("error exporting: %v", err)
	}
}

func init() {
	ExportCmd.PersistentFlags().StringVarP(&export.manifestPath, "manifest", "m", "", "path to manifest file")
	ExportCmd.PersistentFlags().StringVarP(&export.outputPath, "output", "o", "-", "path to output file, - is stdout")
	ExportCmd.PersistentFlags().StringVarP(&export.formatName, "format", "f", "", "format name (see 'hatt formats')")
}

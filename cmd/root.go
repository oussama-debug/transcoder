package cmd

import (
	"io"
	"os"

	"github.com/oussama-debug/transcoder/internal/graphics"
	"github.com/spf13/cobra"
)

const (
	appName      = "transcoder"
	shortAppDesc = "A graphical CLI for transcoder service."
	longAppDesc  = "Transcoder is a CLI to convert your powerpoint files and videos."
)

var (
	// Output writer for testing
	out io.Writer = os.Stdout

	rootCmd = &cobra.Command{
		Use:   appName,
		Short: shortAppDesc,
		Long:  longAppDesc,
		RunE:  run,
	}
)

func init() {
	rootCmd.AddCommand(versionCmd())
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func run(cmd *cobra.Command, args []string) error {
	graphics.Start()
	return nil
}

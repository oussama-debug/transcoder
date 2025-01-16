package cmd

import (
	"fmt"

	"github.com/oussama-debug/transcoder/internal/color"
	"github.com/spf13/cobra"
)

const (
	version = "0.0.1"
	cmdLogo = `
     ____.                                        .__ 
    |    |____  _____    ____  __ _______    ____ |__|
    |    \__  \ \__  \  /  _ \|  |  \__  \  /    \|  |
/\__|    |/ __ \_/ __ \(  <_> )  |  // __ \|   |  \  |
\________(____  (____  /\____/|____/(____  /___|  /__|
              \/     \/                  \/     \/    
`
)

func colorizeMessage(msg string, c color.Color) string {
	if c == -1 {
		return msg
	}
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", c, msg)
}

func printLogo(c color.Color) {
	fmt.Fprintln(out, colorizeMessage(cmdLogo, c))
}

func printVersion(short bool) {
	const _format = "%-20s %s\n"
	var outputColor color.Color = color.Cyan

	if short {
		outputColor = -1
	} else {
		outputColor = color.Cyan
		printLogo(outputColor)
	}

	fmt.Fprintf(out, _format, "Version:", colorizeMessage(version, outputColor))
}

func versionCmd() *cobra.Command {
	var _short bool
	command := &cobra.Command{
		Use:   "version",
		Short: "Print version info",
		Long:  "Print version information",
		RunE: func(cmd *cobra.Command, args []string) error {
			printVersion(_short)
			return nil
		},
	}

	command.Flags().BoolVarP(&_short, "short", "s", false, "Print transcoder version info in short format")
	return command
}

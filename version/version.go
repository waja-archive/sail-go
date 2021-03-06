package version

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/runabove/sail/internal"
)

var versionNewLine bool

// Cmd version
var Cmd = &cobra.Command{
	Use:     "version",
	Short:   "Display Version of sail: sail version",
	Long:    `sail version`,
	Aliases: []string{"v"},
	Run: func(cmd *cobra.Command, args []string) {
		version := fmt.Sprintf("\"%s\"", internal.VERSION)
		internal.FormatOutput([]byte(version), prettyFormater)
	},
}

func prettyFormater(data []byte) {
	fmt.Fprintf(os.Stderr, "Version sail: %s\n", internal.VERSION)
	internal.ReadConfig()
}

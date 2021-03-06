package update

import (
	"fmt"
	"net/http"
	"os"

	"github.com/runabove/sail/internal"

	"github.com/inconshreveable/go-update"
	"github.com/spf13/cobra"
)

// used by CI to inject architecture (linux-amd64, etc...) at build time
var architecture string

// used by CI to inject url for downloading with sail update.
// value of urlUpdate injected at build time
// full URL update is constructed with architecture var :
// urlUpdate + sail- + architecture, sail is the binary
var urlUpdateRelease string
var urlUpdateSnapshot string

func init() {
	if urlUpdateSnapshot != "" {
		Cmd.AddCommand(cmdUpdateSnapshot)
	}
}

// Cmd update
var Cmd = &cobra.Command{
	Use:     "update",
	Short:   "Update sail to the latest release version : sail update",
	Long:    `sail update`,
	Aliases: []string{"up"},
	Run: func(cmd *cobra.Command, args []string) {
		doUpdate(urlUpdateRelease, architecture)
	},
}

var cmdUpdateSnapshot = &cobra.Command{
	Use:     "snapshot",
	Short:   "Update sail to latest snapshot version : sail update snapshot",
	Long:    `sail update snapshot`,
	Aliases: []string{"snap"},
	Run: func(cmd *cobra.Command, args []string) {
		doUpdate(urlUpdateSnapshot, architecture)
	},
}

func doUpdate(baseurl, architecture string) {
	url := fmt.Sprintf("%s/sail-%s", baseurl, architecture)

	if internal.Verbose {
		fmt.Fprintf(os.Stderr, "Debug: updating from %s\n", url)
	}

	resp, err := http.Get(url)
	if err != nil {
		internal.Exit("Error when downloading sail\n", err)
	}
	defer resp.Body.Close()
	err = update.Apply(resp.Body, update.Options{})
	if err != nil {
		internal.Exit("Error when updating sail\n", err)
	}
}

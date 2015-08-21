package container

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"stash.ovh.net/sailabove/sailgo/internal"
)

func init() {
	Cmd.AddCommand(cmdContainerList)
	Cmd.AddCommand(cmdContainerInspect)
	Cmd.AddCommand(cmdContainerAttach)

}

var Cmd = &cobra.Command{
	Use:     "container",
	Short:   "Container commands : sailgo container --help",
	Long:    `Container commands : sailgo container <command>`,
	Aliases: []string{"c", "containers"},
}

var cmdContainerAttach = &cobra.Command{
	Use:   "attach",
	Short: "Attach to a container console : sailgo container attach <applicationName>/<containerId>",
	Long: `Attach to a container console : sailgo container attach <applicationName>/<containerId>
	\"example : sailgo container attach myApp myContainerId"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Invalid usage. sailgo container attach <applicationName>/<containerId>. Please see sailgo container attach --help")
		} else {
			containerAttach(args[0])
		}
	},
}

var cmdContainerList = &cobra.Command{
	Use:     "list",
	Short:   "List docker containers : sailgo container list [applicationName]",
	Aliases: []string{"ls", "ps"},
	Run: func(cmd *cobra.Command, args []string) {
		containerList(internal.GetListApplications(args))
	},
}

var cmdContainerInspect = &cobra.Command{
	Use:   "inspect",
	Short: "Inspect a docker container : sailgo container inspect <applicationName> <containerId>",
	Long: `Inspect a docker container : sailgo container inspect <applicationName> <containerId>
	\"example : sailgo container inspect myApp"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Invalid usage. sailgo container inspect <applicationName> <containerId>. Please see sailgo container inspect --help")
		} else {
			fmt.Println(internal.GetWantJSON(fmt.Sprintf("/applications/%s/containers/%s", args[0], args[1])))
		}
	},
}

func containerAttach(containerID string) {
	t := strings.Split(containerID, "/")
	if len(t) != 2 {
		fmt.Println("Invalid usage. sailgo service inspect <applicationName>/<serviceId>. Please see sailgo service inspect --help")
	} else {
		internal.StreamWant("GET", http.StatusOK, fmt.Sprintf("/applications/%s/containers/%s/attach", t[0], t[1]), nil)
	}
}

func containerList(apps []string) {
	w := tabwriter.NewWriter(os.Stdout, 20, 1, 3, ' ', 0)
	titles := []string{"APPLICATION", "SERVICE", "CONTAINER", "STATE", "DEPLOYED"}
	fmt.Fprintln(w, strings.Join(titles, "\t"))

	containers := []string{}
	var container map[string]interface{}
	for _, app := range apps {
		b := internal.ReqWant("GET", http.StatusOK, fmt.Sprintf("/applications/%s/containers", app), nil)
		internal.Check(json.Unmarshal(b, &containers))
		for _, containerID := range containers {
			b := internal.ReqWant("GET", http.StatusOK, fmt.Sprintf("/applications/%s/containers/%s", app, containerID), nil)
			internal.Check(json.Unmarshal(b, &container))
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t\n", app, container["service"], container["name"], strings.ToUpper(container["state"].(string)), container["deployment_date"])
			w.Flush()
		}
	}
}
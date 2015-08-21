package service

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"stash.ovh.net/sailabove/sailgo/internal"
)

var cmdServiceStop = &cobra.Command{
	Use:   "stop",
	Short: "Stop a docker service : sailgo service stop <applicationName>/<serviceId>",
	Long: `Stop a docker service : sailgo service stop <applicationName>/<serviceId>
	\"example : sailgo service stop myApp/myService"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Invalid usage. sailgo service stop <applicationName>/<serviceId>. Please see sailgo service stop --help")
		} else {
			serviceStop(args[0])
		}
	},
}

func serviceStop(serviceID string) {
	t := strings.Split(serviceID, "/")
	if len(t) != 2 {
		fmt.Println("Invalid usage. sailgo service stop <applicationName>/<serviceId>. Please see sailgo service stop --help")
	} else {
		var empty map[string]interface{}
		em, _ := json.Marshal(empty)
		fmt.Println(internal.PostBodyWantJSON(fmt.Sprintf("/applications/%s/services/%s/stop", t[0], t[1]), em))
	}
}
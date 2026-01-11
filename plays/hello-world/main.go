package main

import (
	"fmt"
	"github.com/relex/aini"
	u "github.com/sunshine69/golang-tools/utils"
	"os"
	"path/filepath"
)

var (
	HostsPattern  string
	inventoryPath string
)

func init() {
	inventoryPath = "inventory/hosts"
	HostsPattern = os.Args[1]
}
func main() {
	inventory := u.Must(aini.ParseFile(inventoryPath))
	inventory.HostsToLower()
	inventory.GroupsToLower()

	inventoryDir := filepath.Dir(inventoryPath)
	if err := inventory.AddVarsLowerCased(inventoryDir); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load inventory variables %s: %v\n", inventoryDir, err)
		os.Exit(4)
	}
	matchedHostsMap := u.Must(inventory.MatchHosts(HostsPattern))
	execHost := matchedHostsMap[HostsPattern]

	execHost.Vars["deploy_var"] = "new value override"

	// From now on you can use inventory vars in your exec command, template command etc whatever you need
	println(u.JsonDump(execHost.Vars, ""))
	o := u.Must(u.RunSystemCommandV2(u.GoTemplateString(`echo env: {{.env}}
echo deploy_var: {{.deploy_var}}
	`, execHost.Vars), true))
	println(o)

}

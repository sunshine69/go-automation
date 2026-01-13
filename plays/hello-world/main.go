package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/relex/aini"
	ag "github.com/sunshine69/automation-go/lib"
	u "github.com/sunshine69/golang-tools/utils"
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
	println("*** DUMP ALL VARS ***", u.JsonDump(execHost.Vars, ""))

	// Flattern all vars to resovle into values
	vars := u.Must(ag.FlattenAllVars(StringMapToAnyMap(execHost.Vars)))
	println("*** DUMP ALL FLATTERN VARS ***", u.JsonDump(vars, ""))

	o := u.Must(u.RunSystemCommandV2(u.GoTemplateString(`echo env: {{.env}}
echo deploy_var: {{.deploy_var}}
	`, vars), true))
	println("*** OUTPUT COMMAND ***", o)

	println("complex_var: ", vars["complex_var"].(string))
}

// StringMapToAnyMap converts map[string]string to map[string]any
func StringMapToAnyMap(m map[string]string) map[string]any {
	result := make(map[string]any, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

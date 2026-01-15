package main

import (
	"maps"
	"os"
	"path/filepath"
	"strings"

	"github.com/relex/aini"
	ag "github.com/sunshine69/automation-go/lib"
	u "github.com/sunshine69/golang-tools/utils"
)

// Inventory parser block
var (
	HostsPattern    string
	InventoryPath   string
	MatchedHostsMap map[string]*aini.Host
	HostList        []string
	Inventory       *aini.InventoryData
	Vars            map[string]any = make(map[string]any)
)

func init() {
	println("Args Length: ", len(os.Args))
	HostsPattern = os.Args[1]
	if len(os.Args) > 2 {
		InventoryPath = os.Args[2]
	} else {
		InventoryPath = "inventory/hosts.ini"
	}
	LoadInventory(InventoryPath)
}

func LoadInventory(inventoryPath string) {
	Inventory = u.Must(aini.ParseFile(inventoryPath))
	inventoryDir := filepath.Dir(inventoryPath)
	u.CheckErr(Inventory.AddVars(inventoryDir), "AddVars")
	MatchedHostsMap = u.Must(Inventory.MatchHosts(HostsPattern))
	HostList = u.MapKeysToSlice(MatchedHostsMap)

	if len(os.Args) > 3 {
		// Loads command line vars
		for _, item := range os.Args[3:] {
			_tmp := strings.Split(item, "=")
			key, val := strings.TrimSpace(_tmp[0]), strings.TrimSpace(_tmp[1])
			Vars[key] = val
		}
	}
}

// end

func playHost(host *aini.Host) {
	// Inventory.HostsToLower()
	// Inventory.GroupsToLower()
	println("Start play host: " + host.Name)
	vars := u.Must(ag.FlattenAllVars(u.StringMapToAnyMap(host.Vars)))
	maps.Copy(vars, Vars)
	// From now on you can use inventory vars in your exec command, template command etc whatever you need
	println("*** DUMP ALL VARS ***", u.JsonDump(vars, ""))
}

func main() {
	for _, h := range HostList { // can spawn go routine to run in parallel or exec in sequence
		playHost(MatchedHostsMap[h])
	}
}

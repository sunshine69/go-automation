package main

import (
	"os"
	"path/filepath"

	ag "github.com/sunshine69/automation-go/lib"
	u "github.com/sunshine69/golang-tools/utils"
)

// Inventory parser block. This can be copied to a new play
var (
	HostsPattern    string
	InventoryPath   string
	HostList        []string
	Inventory       *ag.Inventory
	CommandlineVars map[string]any = make(map[string]any)
)

// Load inventory and return command line vars in Vars. Also populate global vars.
// Per host will get its own vars later
func LoadInventory() {
	println("Args Length: ", len(os.Args))
	if _, ok := CommandlineVars["inventory_dir"]; ok {
		return // Not reload it again
	}
	HostsPattern = os.Args[1]
	if len(os.Args) > 2 {
		InventoryPath = os.Args[2]
	} else {
		InventoryPath = "inventory/hosts.ini"
	}
	println("[INFO] InventoryPath: " + InventoryPath)
	Inventory = ag.ParseInventoryDirAll(InventoryPath)
	Inventory.ParseAllInventoryVars()
	HostList = Inventory.MatchHost(HostsPattern)

	// Populate some default inventory vars. The specific host before use will update this Vars with ansible vars and flattern them
	CommandlineVars["inventory_dir"] = filepath.Dir(InventoryPath)
	CommandlineVars["playbook_dir"] = u.Must(os.Getwd())

	if len(os.Args) > 3 {
		// Loads command line vars
		Inventory.SetFact("", os.Args[3:]...)
	}
}

func init() {
	LoadInventory()
}

// End inventory parser block
func playHost(host string) {
	// Inventory.HostsToLower()
	// Inventory.GroupsToLower()
	println("Start play host: " + host)
	// From now on you can use inventory vars in your exec command, template command etc whatever you need
	println("*** DUMP ALL VARS ***", u.JsonDump(Inventory.Hosts[host].Vars, ""))
}

func main() {
	for _, h := range HostList { // can spawn go routine to run in parallel or exec in sequence
		playHost(h)
	}
}

package main

import (
	"flag"

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
func LoadInventory(InventoryPath, hostPtn string, extraArg ...string) {
	HostsPattern = hostPtn

	println("[INFO] InventoryPath: " + InventoryPath)
	Inventory = ag.ParseInventoryDirAll(InventoryPath)
	Inventory.ParseAllInventoryVars()
	HostList = Inventory.MatchHost(HostsPattern)
	Inventory.SetFact(HostsPattern, extraArg...)
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
	invDir := flag.String("i", "", "Inventory dir")
	hostPtn := flag.String("H", "", "Host pattern to match")
	var extraVars u.ArrayFlags
	flag.Var(&extraVars, "e", "Extra vars to pass to inventory data, like -e action=create_user -e var1=value1")

	flag.Parse()
	LoadInventory(*invDir, *hostPtn, extraVars...)

	for _, h := range HostList { // can spawn go routine to run in parallel or exec in sequence
		playHost(h)
	}
}

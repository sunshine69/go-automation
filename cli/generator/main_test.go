package main

import (
	"testing"

	// ag "github.com/sunshine69/automation-go/lib"

	u "github.com/sunshine69/golang-tools/utils"
)

func TestExpandLayers(t *testing.T) {
	layers := map[string][]string{
		"ops": {"build", "deploy"},
		"pkg": {"encrypt", "sonic"},
		"env": {"dev", "uat", "prod"},
	}

	o := ExpandLayers(layers)
	println(u.JsonDump(o, ""))
}

func TestGenerateFromConfig(t *testing.T) {
	inventory := ParseInventoryGenerator("../../inventory/hosts.yaml")

	// for name := range inventory.Hosts {
	// 	fmt.Println("host:", name)
	// }
	// println("[DEBUG] inventory" + u.JsonDump(inventory, ""))
	h := u.Must(inventory.MatchHosts("*uat*"))
	// println("[DEBUG] Hostmatched" + u.JsonDump(h, ""))
	for hn, h := range h {
		println("Host vars: "+hn, u.JsonDump(h.Vars, ""))
	}
}

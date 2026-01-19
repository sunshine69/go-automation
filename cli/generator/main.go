package main

import (
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"strings"

	"github.com/relex/aini"
	ag "github.com/sunshine69/automation-go/lib"
	u "github.com/sunshine69/golang-tools/utils"
	"gopkg.in/yaml.v3"
)

type Inventory struct {
	Groups map[string]*Group
}

type Group struct {
	Hosts map[string]*Host
	Vars  map[string]any
}

type Host struct {
	Vars map[string]any
}

type GeneratorConfig struct {
	Plugin string              `yaml:"plugin"`
	Hosts  HostConfig          `yaml:"hosts"`
	Layers map[string][]string `yaml:"layers"`
}

type HostConfig struct {
	Name    string            `yaml:"name"`
	Vars    map[string]string `yaml:"vars"`
	Parents []GroupConfig     `yaml:"parents"`
}

type GroupConfig struct {
	Name    string            `yaml:"name"`
	Vars    map[string]string `yaml:"vars"`
	Parents []GroupConfig     `yaml:"parents"`
}

// Layer expansion (cartesian product)
// This turns:
// ops: [update]
// pkg: [letsencrypt]
// env: [dev, uat, prod]
// into:
// []map[string]any{
// 	{"ops":"update","pkg":"letsencrypt","env":"dev"},
// 	{"ops":"update","pkg":"letsencrypt","env":"uat"},
// 	{"ops":"update","pkg":"letsencrypt","env":"prod"},
// }

func ExpandLayers(layers map[string][]string) []map[string]any {
	keys := make([]string, 0, len(layers))
	for k := range layers {
		keys = append(keys, k)
	}

	var res []map[string]any

	var walk func(int, map[string]any)
	walk = func(i int, cur map[string]any) {
		if i == len(keys) {
			// println("[DEBUG]", u.JsonDump(cur, ""), "[END]")
			// m := map[string]any{}
			// for k, v := range cur {
			// 	m[k] = v
			// }
			m := maps.Clone(cur)
			res = append(res, m)
			return
		}

		key := keys[i]
		for _, val := range layers[key] {
			cur[key] = val
			walk(i+1, cur)
		}
	}

	walk(0, map[string]any{})
	return res
}

type IniInventory struct {
	Groups        map[string][]string          // group -> hosts
	GroupVars     map[string]map[string]string // group -> vars
	GroupChildren map[string]map[string]bool   // parent -> child groups
}

func ensureGroupChildren(m map[string]map[string]bool, parent string) {
	if _, ok := m[parent]; !ok {
		m[parent] = map[string]bool{}
	}
}

func ensureGroup(m map[string][]string, name string) {
	if _, ok := m[name]; !ok {
		m[name] = []string{}
	}
}

func ensureGroupVars(m map[string]map[string]string, name string) {
	if _, ok := m[name]; !ok {
		m[name] = map[string]string{}
	}
}

func GenerateIniFromConfig(cfg *GeneratorConfig) string {
	inv := &IniInventory{
		Groups:        map[string][]string{},
		GroupVars:     map[string]map[string]string{},
		GroupChildren: map[string]map[string]bool{},
	}

	objects := ExpandLayers(cfg.Layers)

	for _, ctx := range objects {
		host := ag.TemplateString(cfg.Hosts.Name, ctx)

		ensureGroup(inv.Groups, "all")
		inv.Groups["all"] = append(inv.Groups["all"], host)

		for _, p := range cfg.Hosts.Parents {
			processGroupINI(inv, p, ctx, host, true)
		}

	}

	if len(cfg.Hosts.Vars) > 0 {
		ensureGroupVars(inv.GroupVars, "all")
		for k, v := range cfg.Hosts.Vars {
			inv.GroupVars["all"][k] = v
		}
	}
	// fmt.Printf("DEBUG GroupChildren = %s\n", u.JsonDump(inv.GroupChildren, ""))
	return renderIni(inv)
}

func renderIni(inv *IniInventory) string {
	var b strings.Builder

	for g, hosts := range inv.Groups {
		b.WriteString("[" + g + "]\n")
		for _, h := range hosts {
			b.WriteString(h + "\n")
		}
		b.WriteString("\n")
	}

	for parent, children := range inv.GroupChildren {
		b.WriteString("[" + parent + ":children]\n")
		for c := range maps.Keys(children) {
			b.WriteString(c + "\n")
		}
		// render [parent:children]
	}
	for g, vars := range inv.GroupVars {
		b.WriteString("[" + g + ":vars]\n")
		for k, v := range vars {
			b.WriteString(fmt.Sprintf("%s=%s\n", k, v))
		}
		b.WriteString("\n")
	}

	return b.String()
}

func processGroupINI(
	inv *IniInventory,
	cfg GroupConfig,
	ctx map[string]any,
	host string,
	withHost bool,
) {
	groupName := ag.TemplateString(cfg.Name, ctx)
	if groupName == "" {
		return
	}

	// ensure group exists
	ensureGroup(inv.Groups, groupName)

	// only leaf groups get hosts
	if withHost {
		inv.Groups[groupName] = append(inv.Groups[groupName], host)
	}

	// group vars
	for k, tmpl := range cfg.Vars {
		v := ag.TemplateString(tmpl, ctx)
		ensureGroupVars(inv.GroupVars, groupName)
		inv.GroupVars[groupName][k] = v
	}

	// parents → children relationship
	for _, p := range cfg.Parents {
		parentName := ag.TemplateString(p.Name, ctx)
		if parentName == "" {
			continue
		}

		// 🔑 THIS IS THE CRITICAL LINE
		ensureGroupChildren(inv.GroupChildren, parentName)
		inv.GroupChildren[parentName][groupName] = true

		// ensure parent exists
		ensureGroup(inv.Groups, parentName)

		// recurse upward WITHOUT hosts
		processGroupINI(inv, p, ctx, host, false)
	}
}

func ParseInventoryGenerator(inventoryFile string) *aini.InventoryData {
	datab := u.Must(os.ReadFile(inventoryFile))
	invConfig := GeneratorConfig{}
	u.CheckErr(yaml.Unmarshal(datab, &invConfig), "")
	iniText := GenerateIniFromConfig(&invConfig)
	inventory := u.Must(aini.ParseString(iniText))
	inventory.AddVars(filepath.Dir(inventoryFile))
	return inventory
}

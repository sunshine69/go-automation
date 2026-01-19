package main

import (
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/relex/aini"
	ag "github.com/sunshine69/automation-go/lib"
	u "github.com/sunshine69/golang-tools/utils"
	"gopkg.in/yaml.v3"
)

// GeneratorInventory represents the structure of the generator plugin inventory
type GeneratorInventory struct {
	Plugin string                 `yaml:"plugin"`
	Hosts  map[string]interface{} `yaml:"hosts"`
	Layers map[string][]string    `yaml:"layers"`
}

// ParseGeneratorInventory parses a generator format inventory file and returns an aini InventoryData
func ParseGeneratorInventory(filename string) (*aini.InventoryData, error) {
	// Read the YAML file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Parse the generator inventory structure
	var genInv GeneratorInventory
	if err := yaml.Unmarshal(data, &genInv); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	// Generate standard inventory format
	standardInventory, err := generateStandardInventory(&genInv)
	if err != nil {
		return nil, fmt.Errorf("failed to generate inventory: %w", err)
	}

	// Convert to INI format string (aini expects INI format, not YAML)
	iniContent := convertToINI(standardInventory)

	// Parse using aini with a strings.Reader
	reader := strings.NewReader(iniContent)
	inventoryData, err := aini.Parse(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to parse with aini: %w", err)
	}

	// Load group_vars and host_vars from the inventory directory using aini's built-in function
	inventoryDir := filepath.Dir(filename)
	inventoryData.Reconcile()

	if err := inventoryData.AddVars(inventoryDir); err != nil {
		return nil, fmt.Errorf("failed to load vars from %s: %w", inventoryDir, err)
	}

	return inventoryData, nil
}

// convertToINI converts the inventory structure to INI format
func convertToINI(inventory map[string]interface{}) string {
	var builder strings.Builder

	allGroup := inventory["all"].(map[string]interface{})
	children := allGroup["children"].(map[string]interface{})

	// Process each group
	processedGroups := make(map[string]bool)

	for groupName, groupDataInterface := range children {
		writeGroupToINI(&builder, groupName, groupDataInterface.(map[string]interface{}), processedGroups)
	}

	return builder.String()
}

// writeGroupToINI writes a group and its contents to INI format
func writeGroupToINI(builder *strings.Builder, groupName string, groupData map[string]interface{}, processed map[string]bool) {
	if processed[groupName] {
		return
	}
	processed[groupName] = true

	// Write hosts section
	if hosts, ok := groupData["hosts"].(map[string]interface{}); ok && len(hosts) > 0 {
		builder.WriteString(fmt.Sprintf("[%s]\n", groupName))
		for hostName := range hosts {
			builder.WriteString(fmt.Sprintf("%s\n", hostName))
		}
		builder.WriteString("\n")
	}

	// Write children section
	if children, ok := groupData["children"].(map[string]interface{}); ok && len(children) > 0 {
		builder.WriteString(fmt.Sprintf("[%s:children]\n", groupName))
		for childName := range children {
			builder.WriteString(fmt.Sprintf("%s\n", childName))
		}
		builder.WriteString("\n")
	}

	// Write vars section
	if vars, ok := groupData["vars"].(map[string]interface{}); ok && len(vars) > 0 {
		builder.WriteString(fmt.Sprintf("[%s:vars]\n", groupName))
		for varName, varValue := range vars {
			builder.WriteString(fmt.Sprintf("%s=%v\n", varName, varValue))
		}
		builder.WriteString("\n")
	}
}

// convertGroupsToYAML converts groups to proper YAML structure for aini
func convertGroupsToYAML(groups map[string]*GroupData) map[string]interface{} {
	result := make(map[string]interface{})

	for groupName, groupData := range groups {
		groupMap := make(map[string]interface{})

		// Add hosts
		if len(groupData.Hosts) > 0 {
			hostsMap := make(map[string]interface{})
			for _, host := range groupData.Hosts {
				hostsMap[host] = make(map[string]interface{}) // Empty map instead of nil
			}
			groupMap["hosts"] = hostsMap
		}

		// Add children
		if len(groupData.Children) > 0 {
			childrenMap := make(map[string]interface{})
			for _, child := range groupData.Children {
				childrenMap[child] = make(map[string]interface{}) // Empty map instead of nil
			}
			groupMap["children"] = childrenMap
		}

		// Add vars
		if len(groupData.Vars) > 0 {
			groupMap["vars"] = groupData.Vars
		}

		result[groupName] = groupMap
	}

	return result
}

// generateStandardInventory converts generator format to standard Ansible inventory
func generateStandardInventory(genInv *GeneratorInventory) (map[string]interface{}, error) {
	// Extract layer names and values dynamically
	layerNames := make([]string, 0, len(genInv.Layers))
	layerValues := make(map[string][]string)

	for layerName, values := range genInv.Layers {
		layerNames = append(layerNames, layerName)
		layerValues[layerName] = values
	}

	// Parse the hosts structure to understand the template
	hostsConfig, err := parseHostsConfig(genInv.Hosts)
	if err != nil {
		return nil, err
	}

	// Track all groups to avoid duplicates
	groups := make(map[string]*GroupData)

	// Generate all combinations (cartesian product) of layer values
	combinations := generateCombinations(layerNames, layerValues)

	for _, combo := range combinations {
		// Generate host name from template
		hostName := expandTemplate(hostsConfig.HostTemplate, combo)

		// Create all parent groups for this combination
		createGroupsFromStructure(groups, hostsConfig.ParentStructure, combo, hostName)
	}

	// Convert groups to children format using the helper function
	children := convertGroupsToYAML(groups)

	all := make(map[string]interface{})
	all["children"] = children

	return map[string]interface{}{
		"all": all,
	}, nil
}

// GroupData holds information about a group
type GroupData struct {
	Hosts    []string
	Children []string
	Vars     map[string]interface{}
}

// HostsConfig holds the parsed hosts configuration
type HostsConfig struct {
	HostTemplate    string
	Vars            map[string]string
	ParentStructure []ParentNode
}

// ParentNode represents a parent group in the hierarchy
type ParentNode struct {
	NameTemplate string
	Vars         map[string]string
	Parents      []ParentNode
}

// parseHostsConfig extracts the host template and parent structure
func parseHostsConfig(hostsData map[string]interface{}) (*HostsConfig, error) {
	config := &HostsConfig{}

	// Get host name template
	if name, ok := hostsData["name"].(string); ok {
		config.HostTemplate = name
	}

	// Parse parent structure
	if parents, ok := hostsData["parents"]; ok {
		switch p := parents.(type) {
		case []interface{}:
			config.ParentStructure = parseParentNodes(p)
		case []map[interface{}]interface{}:
			// Convert to []interface{}
			interfaceSlice := make([]interface{}, len(p))
			for i, v := range p {
				interfaceSlice[i] = v
			}
			config.ParentStructure = parseParentNodes(interfaceSlice)
		}
	}

	return config, nil
}

// parseParentNodes recursively parses parent node structure
func parseParentNodes(parents []interface{}) []ParentNode {
	nodes := make([]ParentNode, 0, len(parents))

	for _, parent := range parents {
		// Handle both map[interface{}]interface{} and map[string]interface{}
		var parentMap map[string]interface{}

		switch p := parent.(type) {
		case map[interface{}]interface{}:
			// Convert to map[string]interface{}
			parentMap = make(map[string]interface{})
			for k, v := range p {
				if keyStr, ok := k.(string); ok {
					parentMap[keyStr] = v
				}
			}
		case map[string]interface{}:
			parentMap = p
		default:
			continue
		}

		node := ParentNode{
			Vars: make(map[string]string),
		}

		// Get name template
		if name, ok := parentMap["name"].(string); ok {
			node.NameTemplate = name
		}

		// Get vars
		if vars, ok := parentMap["vars"]; ok {
			switch v := vars.(type) {
			case map[interface{}]interface{}:
				for k, val := range v {
					if keyStr, ok := k.(string); ok {
						if valStr, ok := val.(string); ok {
							node.Vars[keyStr] = valStr
						}
					}
				}
			case map[string]interface{}:
				for k, val := range v {
					if valStr, ok := val.(string); ok {
						node.Vars[k] = valStr
					}
				}
			}
		}

		// Get nested parents
		if nestedParents, ok := parentMap["parents"]; ok {
			if nestedSlice, ok := nestedParents.([]interface{}); ok {
				node.Parents = parseParentNodes(nestedSlice)
			}
		}

		nodes = append(nodes, node)
	}

	return nodes
}

// generateCombinations creates all possible combinations of layer values
func generateCombinations(layerNames []string, layerValues map[string][]string) []map[string]string {
	if len(layerNames) == 0 {
		return []map[string]string{{}}
	}

	// Start with the first layer
	firstLayer := layerNames[0]
	firstValues := layerValues[firstLayer]

	// Recursively generate combinations for remaining layers
	remainingLayers := layerNames[1:]
	subCombinations := generateCombinations(remainingLayers, layerValues)

	// Combine first layer with sub-combinations
	var result []map[string]string
	for _, value := range firstValues {
		for _, subCombo := range subCombinations {
			combo := make(map[string]string)
			combo[firstLayer] = value
			for k, v := range subCombo {
				combo[k] = v
			}
			result = append(result, combo)
		}
	}

	return result
}

// createGroupsFromStructure creates groups based on the parent structure
func createGroupsFromStructure(groups map[string]*GroupData, parents []ParentNode, vars map[string]string, hostName string) {
	// Process each parent and create its structure
	for _, parent := range parents {
		groupName := expandTemplate(parent.NameTemplate, vars)

		// Ensure group exists
		if _, exists := groups[groupName]; !exists {
			groups[groupName] = &GroupData{
				Hosts:    []string{},
				Children: []string{},
				Vars:     make(map[string]interface{}),
			}
		}

		// Add vars to group
		for varKey, varTemplate := range parent.Vars {
			varValue := expandTemplate(varTemplate, vars)
			groups[groupName].Vars[varKey] = varValue
		}

		// Process parent groups
		if len(parent.Parents) > 0 {
			// Recursively create parent groups (but don't add the host to them)
			for _, parentNode := range parent.Parents {
				// Create the parent group structure without adding the host
				createGroupStructureOnly(groups, parentNode, vars)

				parentGroupName := expandTemplate(parentNode.NameTemplate, vars)

				// Add parent groups as children of current group
				childExists := false
				for _, c := range groups[groupName].Children {
					if c == parentGroupName {
						childExists = true
						break
					}
				}
				if !childExists {
					groups[groupName].Children = append(groups[groupName].Children, parentGroupName)
				}
			}
		}

		// Add the host to THIS parent group (not just the first one)
		// Each top-level parent should contain the host
		hostExists := false
		for _, h := range groups[groupName].Hosts {
			if h == hostName {
				hostExists = true
				break
			}
		}
		if !hostExists {
			groups[groupName].Hosts = append(groups[groupName].Hosts, hostName)
		}
	}
}

// createGroupStructureOnly creates group structures and relationships without adding hosts
func createGroupStructureOnly(groups map[string]*GroupData, parent ParentNode, vars map[string]string) {
	groupName := expandTemplate(parent.NameTemplate, vars)

	// Ensure group exists
	if _, exists := groups[groupName]; !exists {
		groups[groupName] = &GroupData{
			Hosts:    []string{},
			Children: []string{},
			Vars:     make(map[string]interface{}),
		}
	}

	// Add vars to group
	for varKey, varTemplate := range parent.Vars {
		varValue := expandTemplate(varTemplate, vars)
		groups[groupName].Vars[varKey] = varValue
	}

	// Process nested parents
	if len(parent.Parents) > 0 {
		for _, parentNode := range parent.Parents {
			// Recursively create parent structure
			createGroupStructureOnly(groups, parentNode, vars)

			parentGroupName := expandTemplate(parentNode.NameTemplate, vars)

			// Add parent as child
			childExists := false
			for _, c := range groups[groupName].Children {
				if c == parentGroupName {
					childExists = true
					break
				}
			}
			if !childExists {
				groups[groupName].Children = append(groups[groupName].Children, parentGroupName)
			}
		}
	}
}

// GetHostVars retrieves all variables for a specific host, including:
// - Direct host vars
// - Host vars from host_vars directory
// - Group vars from all groups the host belongs to
// - Group vars from group_vars directory
// Variables are resolved in order of precedence (host vars override group vars)
func GetHostVars(inventory *aini.InventoryData, hostName string) (map[string]string, error) {
	// Find the host
	var targetHost *aini.Host
	for _, host := range inventory.Hosts {
		if host.Name == hostName {
			targetHost = host
			break
		}
	}

	if targetHost == nil {
		return nil, fmt.Errorf("host %s not found", hostName)
	}

	// Collect all variables with proper precedence
	// Order: group vars (least specific to most specific) -> host vars
	allVars := make(map[string]string)

	// Get all groups this host belongs to
	hostGroups := findHostGroups(inventory, hostName)

	// Apply group vars (from least specific to most specific)
	// First apply vars from parent groups, then child groups
	for _, groupName := range hostGroups {
		if group, ok := inventory.Groups[groupName]; ok {
			for k, v := range group.Vars {
				allVars[k] = v
			}
		}
	}

	// Apply host vars (highest priority)
	for k, v := range targetHost.Vars {
		allVars[k] = v
	}

	return allVars, nil
}

// findHostGroups returns all groups that a host belongs to
func findHostGroups(inventory *aini.InventoryData, hostName string) []string {
	var groups []string

	for groupName, group := range inventory.Groups {
		for _, host := range group.Hosts {
			if host.Name == hostName {
				groups = append(groups, groupName)
				break
			}
		}
	}

	return groups
}
func expandTemplate(template string, vars map[string]string) string {
	result := template

	// Find all {{ variable }} patterns
	re := regexp.MustCompile(`\{\{\s*(\w+)\s*\}\}`)
	matches := re.FindAllStringSubmatch(template, -1)

	for _, match := range matches {
		if len(match) >= 2 {
			varName := match[1]
			if value, ok := vars[varName]; ok {
				placeholder := match[0]
				result = strings.ReplaceAll(result, placeholder, value)
			}
		}
	}

	return result
}

// Example usage
func main() {
	inputFile := os.Args[1]
	inventory, err := ParseGeneratorInventory(inputFile)
	if err != nil {
		fmt.Printf("Error parsing inventory using generator. Will try to pass normal: %v\n", err)
		inventory = u.Must(aini.ParseFile(inputFile))
	}

	// Example: List all hosts
	// fmt.Println("Hosts:")
	// for _, host := range inventory.Hosts {
	// 	fmt.Printf("  - %s\n", host.Name)
	// }

	// Example: List all groups
	fmt.Println("\nGroups:")
	for groupName := range inventory.Groups {
		fmt.Printf("  - %s\n", groupName)
		if group, ok := inventory.Groups[groupName]; ok {
			if len(group.Hosts) > 0 {
				fmt.Printf("    Hosts: %d\n", len(group.Hosts))
			}
			if len(group.Children) > 0 {
				fmt.Printf("    Children: %d\n", len(group.Children))
			}
			if len(group.Vars) > 0 {
				fmt.Printf("    Vars: %v\n", group.Vars)
			}
		}
	}

	// Example: Get hosts in a specific group
	fmt.Println("\nExample - Hosts in group:")
	myinvent := map[string]map[string]string{}
	for gn, group := range inventory.Groups {
		println("Group name: " + gn)
		for _, host := range group.Hosts {
			fmt.Printf("  - %s\n", host.Name)
			if myinvent[host.Name] == nil {
				myinvent[host.Name] = make(map[string]string)
			}
			myinvent[host.Name] = u.Must(GetHostVars(inventory, host.Name))
		}
	}

	// inventory.HostsToLower()
	// inventory.GroupsToLower()

	inventoryDir := filepath.Dir(inputFile)
	println("\nStarted with InventoryDir: " + inventoryDir)
	if err := inventory.AddVars(inventoryDir); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load inventory variables %s: %v\n", inventoryDir, err)
		os.Exit(4)
	}
	hostsPattern := os.Args[2]
	matchedHostsMap := u.Must(inventory.MatchHosts(hostsPattern))
	HostList := u.MapKeysToSlice(matchedHostsMap)
	println("[INFO] Hosts matched the pattern: " + u.JsonDump(HostList, ""))
	execHost := matchedHostsMap[HostList[0]]
	Vars := execHost.Vars
	// maps.Copy(Vars, execHost.InventoryVars)
	maps.Copy(Vars, myinvent[HostList[0]]) // TODO vars declared in the hosts config does not get in
	vars := u.Must(ag.FlattenAllVars(u.StringMapToAnyMap(Vars)))
	println("Inventory Matched with Host: "+hostsPattern+" \n", u.JsonDump(vars, ""))
}

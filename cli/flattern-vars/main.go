package main

import (
	"fmt"

	ag "github.com/sunshine69/automation-go/lib"
	u "github.com/sunshine69/golang-tools/utils"
)

// Example TemplateString implementation (you'll use your own)
// func TemplateString(s string, data map[string]any) string {
// 	// This is a placeholder - use your actual implementation
// 	re := regexp.MustCompile(`\{\{\s*(\w+)\s*\}\}`)
// 	return re.ReplaceAllStringFunc(s, func(match string) string {
// 		key := re.FindStringSubmatch(match)[1]
// 		if val, ok := data[key]; ok {
// 			return fmt.Sprintf("%v", val)
// 		}
// 		return match
// 	})
// }

// Example usage
func main() {
	data := map[string]any{
		"var1": "var1 plus-{{ var2 }}",
		"var2": "var2 plus-{{ var3 }}",
		"var3": "has final value",
	}

	visited := make(map[string]bool)
	result, err := ag.FlattenVar("var1", data, visited)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("var1 flattened: %s\n", result)
	// Output: var1 flattened: var1 plus var2 plus has final value

	data1, err := ag.FlattenAllVars(data)
	if err != nil {
		println(err.Error())
		return
	}
	println(u.JsonDump(data1, ""))
}

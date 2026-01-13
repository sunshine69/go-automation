package main

import (
	"fmt"

	ag "github.com/sunshine69/automation-go/lib"
)

// func ExampleFlatternVars() {
// 	data := map[string]any{
// 		"var1": "var1 plus {{ var2 }}",
// 		"var2": "Var2 plus {{ var3 }}",
// 		"var3": "This si var3",
// 	}
// 	o := FlatternVars(data["var1"].(string), data)
// 	fmt.Println(o)

// 	// XXXOutput: var1 plus Var2 plus This si var3
// 	o := ag.TemplateString(`var1 {{ var2 }}`, map[string]any{"var2": "Value var2 "})
// 	fmt.Println(o)
// 	// Output: var1 Value var2
// }

package main

import (
	"testing"

	// ag "github.com/sunshine69/automation-go/lib"
	u "github.com/sunshine69/golang-tools/utils"
	"gopkg.in/yaml.v3"
)

func TestAdhoc(t *testing.T) {
	data := `a: 123
b: 'string value'
`
	o := map[string]any{}
	yaml.Unmarshal([]byte(data), &o)
	println(u.JsonDump(o, ""))
}

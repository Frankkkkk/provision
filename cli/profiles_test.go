package cli

import (
	"testing"
)

func TestProfileCli(t *testing.T) {

	var profileCreateBadJSONString = "{asdgasdg"

	var profileCreateBadJSON2String = "[asdgasdg]"
	var profileCreateInputString string = `{
  "Name": "john",
  "Params": {
    "FRED": "GREG"
  }
}
`
	var profileUpdateBadJSONString = "asdgasdg"

	var profileUpdateInputString string = `{
  "Params": {
    "JESSIE": "JAMES"
  }
}
`
	var profilesParamsNextString string = `{
  "jj": 3
}
`

	cliTest(true, false, "profiles").run(t)
	cliTest(false, false, "profiles", "list").run(t)
	cliTest(true, true, "profiles", "create").run(t)
	cliTest(true, true, "profiles", "create", "john", "john2").run(t)
	cliTest(false, true, "profiles", "create", profileCreateBadJSONString).run(t)
	cliTest(false, true, "profiles", "create", profileCreateBadJSON2String).run(t)
	cliTest(false, false, "profiles", "create", profileCreateInputString).run(t)
	cliTest(false, true, "profiles", "create", profileCreateInputString).run(t)
	cliTest(false, false, "profiles", "list").run(t)
	cliTest(false, false, "profiles", "list", "Name=fred").run(t)
	cliTest(false, false, "profiles", "list", "Name=john").run(t)
	cliTest(true, true, "profiles", "show").run(t)
	cliTest(true, true, "profiles", "show", "john", "john2").run(t)
	cliTest(false, true, "profiles", "show", "john2").run(t)
	cliTest(false, false, "profiles", "show", "john").run(t)
	cliTest(true, true, "profiles", "exists").run(t)
	cliTest(true, true, "profiles", "exists", "john", "john2").run(t)
	cliTest(false, false, "profiles", "exists", "john").run(t)
	cliTest(false, true, "profiles", "exists", "john2").run(t)
	cliTest(true, true, "profiles", "exists", "john", "john2").run(t)
	cliTest(true, true, "profiles", "update").run(t)
	cliTest(true, true, "profiles", "update", "john", "john2", "john3").run(t)
	cliTest(false, true, "profiles", "update", "john", profileUpdateBadJSONString).run(t)
	cliTest(false, false, "profiles", "update", "john", profileUpdateInputString).run(t)
	cliTest(false, true, "profiles", "update", "john2", profileUpdateInputString).run(t)
	cliTest(false, false, "profiles", "show", "john").run(t)
	cliTest(false, false, "profiles", "show", "john").run(t)
	cliTest(true, true, "profiles", "destroy").run(t)
	cliTest(true, true, "profiles", "destroy", "john", "june").run(t)
	cliTest(false, false, "profiles", "destroy", "john").run(t)
	cliTest(false, true, "profiles", "destroy", "john").run(t)
	cliTest(false, false, "profiles", "list").run(t)
	cliTest(false, false, "profiles", "create", "-").Stdin(profileCreateInputString + "\n").run(t)
	cliTest(false, false, "profiles", "list").run(t)
	cliTest(false, false, "profiles", "update", "john", "-").Stdin(profileUpdateInputString + "\n").run(t)
	cliTest(false, false, "profiles", "show", "john").run(t)
	cliTest(true, true, "profiles", "get").run(t)
	cliTest(false, true, "profiles", "get", "john2", "param", "john2").run(t)
	cliTest(false, false, "profiles", "get", "john", "param", "john2").run(t)
	cliTest(true, true, "profiles", "set").run(t)
	cliTest(false, true, "profiles", "set", "john2", "param", "john2", "to", "cow").run(t)
	cliTest(false, false, "profiles", "set", "john", "param", "john2", "to", "cow").run(t)
	cliTest(false, false, "profiles", "get", "john", "param", "john2").run(t)
	cliTest(false, false, "profiles", "set", "john", "param", "john2", "to", "3").run(t)
	cliTest(false, false, "profiles", "set", "john", "param", "john3", "to", "4").run(t)
	cliTest(false, false, "profiles", "get", "john", "param", "john2").run(t)
	cliTest(false, false, "profiles", "get", "john", "param", "john3").run(t)
	cliTest(false, false, "profiles", "set", "john", "param", "john2", "to", "null").run(t)
	cliTest(false, false, "profiles", "get", "john", "param", "john2").run(t)
	cliTest(false, false, "profiles", "get", "john", "param", "john3").run(t)
	cliTest(true, true, "profiles", "params").run(t)
	cliTest(false, true, "profiles", "params", "john2").run(t)
	cliTest(false, false, "profiles", "params", "john").run(t)
	cliTest(false, true, "profiles", "params", "john2", profilesParamsNextString).run(t)
	cliTest(false, false, "profiles", "params", "john", profilesParamsNextString).run(t)
	cliTest(false, false, "profiles", "params", "john").run(t)
	cliTest(false, false, "profiles", "show", "john").run(t)
	cliTest(false, false, "profiles", "destroy", "john").run(t)
	cliTest(false, false, "profiles", "list").run(t)
}

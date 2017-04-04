package cli

import (
	"fmt"

	"github.com/rackn/rocket-skates/client/interfaces"
	"github.com/spf13/cobra"
)

type InterfaceOps struct{}

func (be InterfaceOps) List() (interface{}, error) {
	d, e := session.Interfaces.ListInterfaces(interfaces.NewListInterfacesParams(), basicAuth)
	if e != nil {
		return nil, e
	}
	return d.Payload, nil
}

func (be InterfaceOps) Get(id string) (interface{}, error) {
	d, e := session.Interfaces.GetInterface(interfaces.NewGetInterfaceParams().WithName(id), basicAuth)
	if e != nil {
		return nil, e
	}
	return d.Payload, nil
}

func init() {
	tree := addInterfaceCommands()
	App.AddCommand(tree)
}

func addInterfaceCommands() (res *cobra.Command) {
	singularName := "interface"
	name := "interfaces"
	d("Making command tree for %v\n", name)
	res = &cobra.Command{
		Use:   name,
		Short: fmt.Sprintf("Access CLI commands relating to %v", name),
	}
	commands := commonOps(singularName, name, &InterfaceOps{})
	res.AddCommand(commands...)
	return res
}

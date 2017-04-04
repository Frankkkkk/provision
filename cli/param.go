package cli

import (
	"fmt"

	"github.com/rackn/rocket-skates/client/params"
	"github.com/rackn/rocket-skates/models"
	"github.com/spf13/cobra"
)

type ParamOps struct{}

func (be ParamOps) GetType() interface{} {
	return &models.Param{}
}

func (be ParamOps) GetId(obj interface{}) (string, error) {
	param, ok := obj.(*models.Param)
	if !ok {
		return "", fmt.Errorf("Invalid type passed to param create")
	}
	return *param.Name, nil
}

func (be ParamOps) List() (interface{}, error) {
	d, e := session.Params.ListParams(params.NewListParamsParams(), basicAuth)
	if e != nil {
		return nil, e
	}
	return d.Payload, nil
}

func (be ParamOps) Get(id string) (interface{}, error) {
	d, e := session.Params.GetParam(params.NewGetParamParams().WithName(id), basicAuth)
	if e != nil {
		return nil, e
	}
	return d.Payload, nil
}

func (be ParamOps) Create(obj interface{}) (interface{}, error) {
	param, ok := obj.(*models.Param)
	if !ok {
		return nil, fmt.Errorf("Invalid type passed to param create")
	}
	d, e := session.Params.CreateParam(params.NewCreateParamParams().WithBody(param), basicAuth)
	if e != nil {
		return nil, e
	}
	return d.Payload, nil
}

func (be ParamOps) Patch(id string, obj interface{}) (interface{}, error) {
	data, ok := obj.(models.Patch)
	if !ok {
		return nil, fmt.Errorf("Invalid type passed to param patch")
	}
	d, e := session.Params.PatchParam(params.NewPatchParamParams().WithName(id).WithBody(data), basicAuth)
	if e != nil {
		return nil, e
	}
	return d.Payload, nil
}

func (be ParamOps) Delete(id string) (interface{}, error) {
	d, e := session.Params.DeleteParam(params.NewDeleteParamParams().WithName(id), basicAuth)
	if e != nil {
		return nil, e
	}
	return d.Payload, nil
}

func init() {
	tree := addParamCommands()
	App.AddCommand(tree)
}

func addParamCommands() (res *cobra.Command) {
	singularName := "param"
	name := "params"
	d("Making command tree for %v\n", name)
	res = &cobra.Command{
		Use:   name,
		Short: fmt.Sprintf("Access CLI commands relating to %v", name),
	}
	commands := commonOps(singularName, name, &ParamOps{})
	res.AddCommand(commands...)
	return res
}

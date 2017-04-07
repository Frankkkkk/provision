package cli

import (
	"fmt"

	"github.com/digitalrebar/provision/client/leases"
	"github.com/digitalrebar/provision/models"
	"github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
)

type LeaseOps struct{}

func convertStringToAddress(id string) (strfmt.IPv4, error) {
	var s strfmt.IPv4
	err := s.Scan(id)
	if err != nil {
		return "", fmt.Errorf("%v is not a valid IPv4: %v", id, err)
	}
	return s, nil
}

func (be LeaseOps) GetType() interface{} {
	return &models.Lease{}
}

func (be LeaseOps) GetId(obj interface{}) (string, error) {
	lease, ok := obj.(*models.Lease)
	if !ok {
		return "", fmt.Errorf("Invalid type passed to lease create")
	}
	return lease.Addr.String(), nil
}

func (be LeaseOps) List() (interface{}, error) {
	d, e := session.Leases.ListLeases(leases.NewListLeasesParams(), basicAuth)
	if e != nil {
		return nil, e
	}
	return d.Payload, nil
}

func (be LeaseOps) Get(id string) (interface{}, error) {
	s, e := convertStringToAddress(id)
	if e != nil {
		return nil, e
	}
	d, e := session.Leases.GetLease(leases.NewGetLeaseParams().WithAddress(s), basicAuth)
	if e != nil {
		return nil, e
	}
	return d.Payload, nil
}

func (be LeaseOps) Create(obj interface{}) (interface{}, error) {
	lease, ok := obj.(*models.Lease)
	if !ok {
		return nil, fmt.Errorf("Invalid type passed to lease create")
	}
	d, e := session.Leases.CreateLease(leases.NewCreateLeaseParams().WithBody(lease), basicAuth)
	if e != nil {
		return nil, e
	}
	return d.Payload, nil
}

func (be LeaseOps) Patch(id string, obj interface{}) (interface{}, error) {
	data, ok := obj.(models.Patch)
	if !ok {
		return nil, fmt.Errorf("Invalid type passed to lease patch")
	}
	s, e := convertStringToAddress(id)
	if e != nil {
		return nil, e
	}
	d, e := session.Leases.PatchLease(leases.NewPatchLeaseParams().WithAddress(s).WithBody(data), basicAuth)
	if e != nil {
		return nil, e
	}
	return d.Payload, nil
}

func (be LeaseOps) Delete(id string) (interface{}, error) {
	s, e := convertStringToAddress(id)
	if e != nil {
		return nil, e
	}
	d, e := session.Leases.DeleteLease(leases.NewDeleteLeaseParams().WithAddress(s), basicAuth)
	if e != nil {
		return nil, e
	}
	return d.Payload, nil
}

func init() {
	tree := addLeaseCommands()
	App.AddCommand(tree)
}

func addLeaseCommands() (res *cobra.Command) {
	singularName := "lease"
	name := "leases"
	d("Making command tree for %v\n", name)
	res = &cobra.Command{
		Use:   name,
		Short: fmt.Sprintf("Access CLI commands relating to %v", name),
	}

	commands := commonOps(singularName, name, &LeaseOps{})
	res.AddCommand(commands...)
	return res
}

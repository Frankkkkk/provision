package models

import (
	"encoding/json"
	"net"

	"github.com/pborman/uuid"
)

// Plugin Provider describes the available functions that could be
// instantiated by a plugin.
// swagger:model
type PluginProvider struct {
	Meta

	Name    string
	Version string

	HasPublish       bool
	AvailableActions []*AvailableAction

	RequiredParams []string
	OptionalParams []string

	// Ensure that these are in the system.
	// This should be deprecated and should be converted to content.
	Parameters []*Param

	// Content Bundle Yaml string - can be optional or empty
	Content string
}

func (p *PluginProvider) Prefix() string { return "plugin_providers" }
func (p *PluginProvider) Key() string    { return p.Name }

func (p *PluginProvider) SliceOf() interface{} {
	s := []*PluginProvider{}
	return &s
}

func (p *PluginProvider) ToModels(obj interface{}) []Model {
	items := obj.(*[]*PluginProvider)
	res := make([]Model, len(*items))
	for i, item := range *items {
		res[i] = Model(item)
	}
	return res
}

func (p *PluginProvider) Fill() {
	if p.Meta == nil {
		p.Meta = Meta{}
	}
	if p.RequiredParams == nil {
		p.RequiredParams = []string{}
	}
	if p.OptionalParams == nil {
		p.OptionalParams = []string{}
	}
	if p.AvailableActions == nil {
		p.AvailableActions = []*AvailableAction{}
	}
	if p.Parameters == nil {
		p.Parameters = []*Param{}
	}
	for _, a := range p.AvailableActions {
		a.Fill()
	}
	for _, param := range p.Parameters {
		param.Fill()
	}
}

// swagger:model
type PluginProviderUploadInfo struct {
	Path string `json:"path"`
	Size int64  `json:"size"`
}

// Plugins can provide actions for machines
// Assumes that there are parameters on the
// call in addition to the machine.
//
// swagger:model
type AvailableAction struct {
	Provider       string
	Command        string
	RequiredParams []string
	OptionalParams []string
}

func (a *AvailableAction) Fill() {
	if a.RequiredParams == nil {
		a.RequiredParams = []string{}
	}
	if a.OptionalParams == nil {
		a.OptionalParams = []string{}
	}
}

//
// Params is built from the caller, plus
// the machine, plus profiles, plus global.
//
// This is used by the frontend to talk to
// the plugin.
//
type MachineAction struct {
	Name    string
	Uuid    uuid.UUID
	Address net.IP
	BootEnv string
	Command string
	Params  map[string]interface{}
}

func (m *MachineAction) Fill() {
	if m.Params == nil {
		m.Params = map[string]interface{}{}
	}
}

// Id of request, and JSON blob
type PluginClientRequest struct {
	Id     int
	Action string
	Data   []byte
}

// If code == 0,2xx, then success and call should json decode.
// If code != 0,2xx, then error and data is Error.
type PluginClientReply struct {
	Id   int
	Code int
	Data []byte
}

func (r *PluginClientReply) Error() *Error {
	var err Error
	jerr := json.Unmarshal(r.Data, &err)
	if jerr != nil {
		err = Error{Code: 400, Messages: []string{jerr.Error()}, Model: "plugin", Type: "plugin"}
	}
	return &err
}

func (r *PluginClientReply) HasError() bool {
	return r.Code != 0 && (r.Code < 200 || r.Code > 299)
}

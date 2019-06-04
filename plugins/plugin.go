package plugins

import "plugin"

/*
*	Parameter structure, used for defining parameter data
 */
type Parameter struct {
	Name      string `yaml:"name"`
	Mandatory bool   `yaml:"mandatory"`
	HasValue  bool   `yaml:"hasvalue"`
}

/*
*	Action structure, used for defining action data
 */
type PluginSelector struct {
	PluginFile string `yaml:"plugin"`
}

/*
*	ActionOps interface, defining abstract Action Operations
 */
type PluginOps interface {
	Load() (plugin.Plugin, error)
}

func (s *PluginSelector) Load() (*plugin.Plugin, error) {
	return plugin.Open(s.PluginFile)
}

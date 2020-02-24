package main

import (
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	interfaces "github.com/innoq/go-plugins-examples/hashicorp/intefaces"
)

// Here is a real implementation of Greeter
type Hello struct {
	logger hclog.Logger
}

func (h *Hello) Greet() string {
	h.logger.Debug("message from HelloPlugin.Greet")
	return "Hello!"
}

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:       "HelloPlugin",
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})

	hello := &Hello{
		logger: logger,
	}

	var handshakeConfig = plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "BASIC_PLUGIN",
		MagicCookieValue: "hello",
	}

	// pluginMap is the map of plugins we can dispense.
	var pluginMap = map[string]plugin.Plugin{
		"hello": &interfaces.HelloPlugin{Impl: hello},
	}

	logger.Debug("message from plugin", "foo", "bar")

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}

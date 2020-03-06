package interfaces

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

// Hello is the interface that we're exposing as a plugin.
type Hello interface {
	Greet() string
}

// Here is an implementation that talks over RPC
type HelloRPC struct{ client *rpc.Client }

func (h *HelloRPC) Greet() string {
	var resp string
	err := h.client.Call("Plugin.Greet", new(interface{}), &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		panic(err)
	}

	return resp
}

// Here is the RPC server that HelloRPC talks to, conforming to
// the requirements of net/rpc
type HelloRPCServer struct {
	// This is the real implementation
	Impl Hello
}

func (s *HelloRPCServer) Greet(args interface{}, resp *string) error {
	*resp = s.Impl.Greet()
	return nil
}

// This is the implementation of plugin.Plugin so we can serve/consume this
//
// This has two methods: Server must return an RPC server for this plugin
// type. We construct a GreeterRPCServer for this.
//
// Client must return an implementation of our interface that communicates
// over an RPC client. We return GreeterRPC for this.
//
// Ignore MuxBroker. That is used to create more multiplexed streams on our
// plugin connection and is a more advanced use case.
type HelloPlugin struct {
	// Impl Injection
	Impl Hello
}

func (p *HelloPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &HelloRPCServer{Impl: p.Impl}, nil
}

func (HelloPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &HelloRPC{client: c}, nil
}

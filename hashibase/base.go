package hashibase

import (
	"fmt"
	"github.com/hashicorp/go-plugin"
	"net/rpc"
)

type Greeter interface {
	Greet() string
}

type GreeterRpc struct {
	client *rpc.Client
}

func (g *GreeterRpc) Greet() string {
	reply := ""
	err := g.client.Call("Greet", new(any), &reply)
	if err != nil {
		fmt.Printf("error while calling Greet call on the client: %v\n", err)
		return ""
	}
	return reply
}

type GreeterServer struct {
	Impl Greeter
}

func (g *GreeterServer) Greet(args any, resp *string) error {
	*resp = g.Impl.Greet()
	return nil
}

type GreeterPlugin struct {
	Impl Greeter
}

func (p *GreeterPlugin) Server(*plugin.MuxBroker) (any, error) {
	return &GreeterServer{p.Impl}, nil
}

func (p *GreeterPlugin) Client(broker *plugin.MuxBroker, c *rpc.Client) (any, error) {
	return &GreeterRpc{c}, nil
}

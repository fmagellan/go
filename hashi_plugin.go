package main

import (
	"fmt"
	"github.com/hashicorp/go-plugin"
	"github.com/fmagellan/go/hashibase"
)

func main() {
	pluginMap := map[string]plugin.Plugin{
		"Greeter": &hashibase.sha
	}
}

type GreeterHello struct {
}

func (g *GreeterHello) Greet() string {
	fmt.Printf("got Greet called in GreeterHello\n")
	return "Greetings from GreeterHello"
}

var hsConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "Cookie-key-1",
	MagicCookieValue: "Cookie-value-1",
}

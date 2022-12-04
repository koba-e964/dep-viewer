//go:build wireinject
// +build wireinject

package rust

import "github.com/google/wire"

func InitializeConcreteDiscoverer() *ConcreteDiscoverer {
	wire.Build(NewConcreteDiscoverer)
	return &ConcreteDiscoverer{}
}

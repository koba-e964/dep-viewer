// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package rust

// Injectors from wire.go:

func InitializeConcreteDiscoverer() *ConcreteDiscoverer {
	concreteDiscoverer := NewConcreteDiscoverer()
	return concreteDiscoverer
}

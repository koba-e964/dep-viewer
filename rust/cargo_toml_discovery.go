//go:generate mockgen -source=${GOFILE} -destination=./mock/${GOFILE}

package rust

type Discoverer interface {
	DiscoverCargoToml(path string) []string
}

type ConcreteDiscoverer struct{}

func NewConcreteDiscoverer() *ConcreteDiscoverer {
	return nil
}

func (*ConcreteDiscoverer) DiscoverCargoToml(path string) []string {
	return nil
}

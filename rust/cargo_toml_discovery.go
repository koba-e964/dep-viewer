//go:generate mockgen -source=${GOFILE} -destination=./mock/${GOFILE}

package rust

import (
	"io/fs"
	"path/filepath"

	"github.com/pkg/errors"
)

type Discoverer interface {
	DiscoverCargoToml(path string) []string
}

type ConcreteDiscoverer struct{}

func NewConcreteDiscoverer() *ConcreteDiscoverer {
	return &ConcreteDiscoverer{}
}

func (*ConcreteDiscoverer) DiscoverCargoToml(path string) ([]string, error) {
	cargoTomls := []string{}

	err := filepath.WalkDir(path, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return errors.Wrap(err, "error found in filepath.WalkDir")
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Base(path) == "Cargo.toml" {
			cargoTomls = append(cargoTomls, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return cargoTomls, nil
}

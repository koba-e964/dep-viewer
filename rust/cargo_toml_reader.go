package rust

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// ReadAllPathDependencies reads all path dependencies (i.e. dependencies in the same repository)
func ReadAllPathDependencies(tomlDoc string) []string {
	var val struct {
		Dependencies map[string]interface{} `toml:"dependencies"`
	}
	toml.Decode(tomlDoc, &val)
	fmt.Println(val)
	deps := []string{}
	for key, depEntry := range val.Dependencies {
		if value, ok := depEntry.(map[string]interface{}); ok {
			if _, ok := value["path"]; ok {
				deps = append(deps, key)
			}
		}
	}
	return deps
}

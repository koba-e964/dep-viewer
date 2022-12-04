package main

import (
	"fmt"
	"log"
	"os"

	"github.com/koba-e964/dep-viewer/rust"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Usage: go run main.go DIRNAME") // TODO: change it to the actual executable's name
	}
	dirName := os.Args[1]
	disc := rust.InitializeConcreteDiscoverer()
	cargoTomls, err := disc.DiscoverCargoToml(dirName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cargoTomls)
}

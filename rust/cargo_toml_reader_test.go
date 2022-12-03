package rust

import (
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	// https://github.com/RustPython/RustPython/blob/f23ea5eb09cfac83fcb58dd2ae5c1fb1774a4f1e/compiler/codegen/Cargo.toml
	tomlDoc := `[package]
name = "rustpython-codegen"
version = "0.1.2"
description = "Compiler for python code into bytecode for the rustpython VM."
authors = ["RustPython Team"]
repository = "https://github.com/RustPython/RustPython"
license = "MIT"
edition = "2021"

[dependencies]
rustpython-ast = { path = "../ast", features = ["unparse"] }
rustpython-compiler-core = { path = "../core", version = "0.1.1" }

ahash = "0.7.6"
bitflags = "1.3.2"
indexmap = "1.8.1"
itertools = "0.10.3"
log = "0.4.16"
num-complex = { version = "0.4.0", features = ["serde"] }
num-traits = "0.2.14"
thiserror = "1.0"

[dev-dependencies]
rustpython-parser = { path = "../parser" }

insta = "1.14.0"
`
	deps := ReadAllPathDependencies(tomlDoc)
	expectedDeps := []string{"rustpython-ast", "rustpython-compiler-core"}
	if !reflect.DeepEqual(deps, expectedDeps) {
		t.Errorf("expected %v, got %v", expectedDeps, deps)
	}
}

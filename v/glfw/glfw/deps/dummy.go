//go:build required
// +build required

// Package dummy prevents go tooling from stripping the c dependencies.
package dummy

import (
	// Prevent go tooling from stripping out the c source files.
	_ "github.com/2dprototype/tender/v/glfw/glfw/deps/glad"
	_ "github.com/2dprototype/tender/v/glfw/glfw/deps/mingw"
	_ "github.com/2dprototype/tender/v/glfw/glfw/deps/wayland"
)

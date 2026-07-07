//go:build glfw

package stdlib

func init() {
	BuiltinModules["glfw"] = glfwModule
}

//go:build gl

package stdlib

func init() {
	BuiltinModules["gl"] = glModule
	BuiltinModules["graphics"] = graphicsModule
}

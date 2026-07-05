//go:build glut

package stdlib

func init() {
	BuiltinModules["glut"] = glutModule
}

//go:build windows

package stdlib

func init() {
	BuiltinModules["dll"] = dllModule
	BuiltinModules["audio"] = audioModule
	BuiltinModules["wui"] = wuiModule
}

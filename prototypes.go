package tender

// // TypePrototypes holds runtime-defined properties/methods for each type.
// var TypePrototypes = map[string]map[string]Object{
	// "string":          make(map[string]Object),
	// "int":             make(map[string]Object),
	// "float":           make(map[string]Object),
	// "array":           make(map[string]Object),
	// "immutable-array": make(map[string]Object),
	// "map":             make(map[string]Object),
	// "bool":            make(map[string]Object),
// }

// func builtinPrototype(args ...Object) (Object, error) {
	// if len(args) != 3 {
		// return nil, ErrWrongNumArguments
	// }
	
	// typeName, ok := args[0].(*String)
	// propName, ok2 := args[1].(*String)
	// if !ok || !ok2 {
		// return nil, fmt.Errorf("first two arguments must be strings")
	// }
	
	// if _, exists := TypePrototypes[typeName.Value]; !exists {
		// TypePrototypes[typeName.Value] = make(map[string]Object)
	// }
	
	// TypePrototypes[typeName.Value][propName.Value] = args[2]
	// return TrueValue, nil
// }
package stdlib

import (
	"os"
	"path/filepath"
	"github.com/2dprototype/tender"
)

var pathModule = map[string]tender.Object{
	"join":        &tender.UserFunction{Name: "join", Value: pathJoin},
	"base":        &tender.UserFunction{Name: "base", Value: FuncASRS(filepath.Base)},
	"ext":         &tender.UserFunction{Name: "ext", Value: FuncASRS(filepath.Ext)},
	"clean":       &tender.UserFunction{Name: "clean", Value: FuncASRS(filepath.Clean)},
	"dir":         &tender.UserFunction{Name: "dir", Value: FuncASRS(filepath.Dir)},
	"isabs":       &tender.UserFunction{Name: "isabs", Value: FuncASRB(filepath.IsAbs)},
	// "islocal":       &tender.UserFunction{Name: "islocal", Value: FuncASRB(filepath.IsLocal)},
	"abs":         &tender.UserFunction{Name: "abs", Value: FuncASRSE(filepath.Abs)},
	"to_slash":    &tender.UserFunction{Name: "to_slash", Value: FuncASRS(filepath.ToSlash)},
	"from_slash":  &tender.UserFunction{Name: "from_slash", Value: FuncASRS(filepath.FromSlash)},
	"vol":         &tender.UserFunction{Name: "vol", Value: FuncASRS(filepath.VolumeName)},
	
	"walklist":   &tender.UserFunction{Name: "walklist", Value: pathWalkList},
	"splitlist":  &tender.UserFunction{Name: "splitlist", Value: FuncASRSs(filepath.SplitList)},
	"resolve": &tender.UserFunction{
		Name: "resolve", 
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrWrongNumArguments
			}
			
			path, ok := tender.ToString(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgumentType{
					Name: "path", 
					Expected: "string",
				}
			}
			
			resolvedPath := tender.ResolvePath(path)
			return &tender.String{Value: resolvedPath}, nil
		},
	},

	"resolve_from": &tender.UserFunction{
		Name: "resolve_from",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrWrongNumArguments
			}
			
			path, ok1 := tender.ToString(args[0])
			dir, ok2 := tender.ToString(args[1])
			if !ok1 || !ok2 {
				return nil, tender.ErrInvalidArgument
			}
			
			resolvedPath := tender.ResolvePathFromDir(path, dir)
			return &tender.String{Value: resolvedPath}, nil
		},
	},
}

func pathJoin(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) < 2 {
		return nil, tender.ErrWrongNumArguments
	}

	elements := make([]string, len(args))
	for i, arg := range args {
		s, _ := tender.ToString(arg)
		elements[i] = s
	}

	joined := filepath.Join(elements...)
	return &tender.String{Value: joined}, nil
}


func pathWalkList(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}

	root, _ := tender.ToString(args[0])

	var result []string

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Handle the error, or you can ignore it.
			return nil
		}
		result = append(result, path)
		return nil
	})

	if err != nil {
		return nil, nil
	}

	// Convert the result to a Tender list
	var elements []tender.Object
	for _, path := range result {
		elements = append(elements, &tender.String{Value: path})
	}

	return &tender.Array{Value: elements}, nil
}
package stdlib

import (
	"os/exec"
	
	"github.com/2dprototype/tender"
)

func makeOSExecCommand(cmd *exec.Cmd) *tender.ImmutableMap {
	return &tender.ImmutableMap{
		Value: map[string]tender.Object{
			"stderr": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					writer, ok := args[0].(*IOWriter)
					if !ok {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "io.writer",
							Found:    args[0].TypeName(),
						}
					}
					cmd.Stderr = writer.Value
					return nil, nil
				},
			},	
			"stdout": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					writer, ok := args[0].(*IOWriter)
					if !ok {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "io.writer",
							Found:    args[0].TypeName(),
						}
					}
					cmd.Stdout = writer.Value
					return nil, nil
				},
			},	
			"stdin": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					reader, ok := args[0].(*IOReader)
					if !ok {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "io.reader",
							Found:    args[0].TypeName(),
						}
					}
					cmd.Stdin = reader.Value
					return nil, nil
				},
			},
			"string": &tender.NativeFunction{
				Name:  "string",
				Value: FuncARS(cmd.String),
			},	
			"environ": &tender.NativeFunction{
				Name:  "environ",
				Value: FuncARSs(cmd.Environ),
			},		
			// combined_output() => bytes/error
			"combined_output": &tender.NativeFunction{
				Name:  "combined_output",
				Value: FuncARYE(cmd.CombinedOutput),
			},
			// output() => bytes/error
			"output": &tender.NativeFunction{
				Name:  "output",
				Value: FuncARYE(cmd.Output),
			}, //
			// run() => error
			"run": &tender.NativeFunction{
				Name:  "run",
				Value: FuncARE(cmd.Run),
			}, //
			// start() => error
			"start": &tender.NativeFunction{
				Name:  "start",
				Value: FuncARE(cmd.Start),
			}, //
			// wait() => error
			"wait": &tender.NativeFunction{
				Name:  "wait",
				Value: FuncARE(cmd.Wait),
			}, //
			// set_path(path string)
			"set_path": &tender.NativeFunction{
				Name: "set_path",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					s1, ok := tender.ToString(args[0])
					if !ok {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
					}
					cmd.Path = s1
					return tender.NullValue, nil
				},
			},
			// set_dir(dir string)
			"set_dir": &tender.NativeFunction{
				Name: "set_dir",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					s1, ok := tender.ToString(args[0])
					if !ok {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
					}
					cmd.Dir = s1
					return tender.NullValue, nil
				},
			},
			// set_env(env array(string))
			"set_env": &tender.NativeFunction{
				Name: "set_env",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}

					var env []string
					var err error
					switch arg0 := args[0].(type) {
					case *tender.Array:
						env, err = stringArray(arg0.Value, "first")
						if err != nil {
							return nil, err
						}
					case *tender.ImmutableArray:
						env, err = stringArray(arg0.Value, "first")
						if err != nil {
							return nil, err
						}
					default:
						return nil, tender.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "array",
							Found:    arg0.TypeName(),
						}
					}
					cmd.Env = env
					return tender.NullValue, nil
				},
			},
			// process() => imap(process)
			"process": &tender.NativeFunction{
				Name: "process",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 0 {
						return nil, tender.ErrWrongNumArguments
					}
					return makeOSProcess(cmd.Process), nil
				},
			},
		},
	}
}

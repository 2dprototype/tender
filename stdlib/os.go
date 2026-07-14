package stdlib

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/2dprototype/tender"
)

var osModule = map[string]tender.Object{
	// "stdout":              &IOWriter{Value: os.Stdout},
	// "stderr":              &IOWriter{Value: os.Stderr},
	// "stdin":               &IOReader{Value: os.Stdin},
	"stdout": &tender.NativeFunction{
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrWrongNumArguments
			}
			return &IOWriter{Value: os.Stdout}, nil
		},
	},	
	"stderr": &tender.NativeFunction{
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrWrongNumArguments
			}
			return &IOWriter{Value: os.Stderr}, nil
		},
	},
	"stdin": &tender.NativeFunction{
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrWrongNumArguments
			}
			return &IOReader{Value: os.Stdin}, nil
		},
	},
	"platform":            &tender.String{Value: runtime.GOOS},
	"arch":                &tender.String{Value: runtime.GOARCH},
	"o_rdonly":            &tender.Int{Value: int64(os.O_RDONLY)},
	"o_wronly":            &tender.Int{Value: int64(os.O_WRONLY)},
	"o_rdwr":              &tender.Int{Value: int64(os.O_RDWR)},
	"o_append":            &tender.Int{Value: int64(os.O_APPEND)},
	"o_create":            &tender.Int{Value: int64(os.O_CREATE)},
	"o_excl":              &tender.Int{Value: int64(os.O_EXCL)},
	"o_sync":              &tender.Int{Value: int64(os.O_SYNC)},
	"o_trunc":             &tender.Int{Value: int64(os.O_TRUNC)},
	"mode_dir":            &tender.Int{Value: int64(os.ModeDir)},
	"mode_append":         &tender.Int{Value: int64(os.ModeAppend)},
	"mode_exclusive":      &tender.Int{Value: int64(os.ModeExclusive)},
	"mode_temporary":      &tender.Int{Value: int64(os.ModeTemporary)},
	"mode_symlink":        &tender.Int{Value: int64(os.ModeSymlink)},
	"mode_device":         &tender.Int{Value: int64(os.ModeDevice)},
	"mode_named_pipe":     &tender.Int{Value: int64(os.ModeNamedPipe)},
	"mode_socket":         &tender.Int{Value: int64(os.ModeSocket)},
	"mode_setuid":         &tender.Int{Value: int64(os.ModeSetuid)},
	"mode_setgui":         &tender.Int{Value: int64(os.ModeSetgid)},
	"mode_char_device":    &tender.Int{Value: int64(os.ModeCharDevice)},
	"mode_sticky":         &tender.Int{Value: int64(os.ModeSticky)},
	"mode_type":           &tender.Int{Value: int64(os.ModeType)},
	"mode_perm":           &tender.Int{Value: int64(os.ModePerm)},
	"path_separator":      &tender.Char{Value: os.PathSeparator},
	"path_list_separator": &tender.Char{Value: os.PathListSeparator},
	"dev_null":            &tender.String{Value: os.DevNull},
	"seek_set":            &tender.Int{Value: int64(io.SeekStart)},
	"seek_cur":            &tender.Int{Value: int64(io.SeekCurrent)},
	"seek_end":            &tender.Int{Value: int64(io.SeekEnd)},
	"args": &tender.NativeFunction{
		Name:      "args",
		Value:     osArgs,
		NeedVMObj: true,
	}, // args() => array(string)
	"chdir": &tender.NativeFunction{
		Name:  "chdir",
		Value: FuncASRE(os.Chdir),
	}, // chdir(dir string) => error
	"chmod": osFuncASFmRE("chmod", os.Chmod), // chmod(name string, mode int) => error
	"chtimes": &tender.NativeFunction{
		Name:  "chtimes",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrWrongNumArguments
			}
			name, _ := tender.ToString(args[0])
			atime, _ := tender.ToTime(args[1])
			mtime, _ := tender.ToTime(args[2])
			err := os.Chtimes(name, atime, mtime)
			if err != nil {
				return wrapError(err), nil
			}
			return nil, nil
		},
	},
	"chown": &tender.NativeFunction{
		Name:  "chown",
		Value: FuncASIIRE(os.Chown),
	},
	"clearenv": &tender.NativeFunction{
		Name:  "clearenv",
		Value: FuncAR(os.Clearenv),
	}, // clearenv()
	"copy": &tender.NativeFunction{
		Name:  "copy",
		Value: osCopy,
	}, // copy(src string, dest string) => error
	"environ": &tender.NativeFunction{
		Name:  "environ",
		Value: FuncARSs(os.Environ),
	}, // environ() => array(string)
	"exit": &tender.NativeFunction{
		Name:  "exit",
		Value: FuncAIR(os.Exit),
	}, // exit(code int)
	"executable": &tender.NativeFunction{
		Name:  "executable",
		Value: FuncARSE(os.Executable),
	}, 
	"expand_env": &tender.NativeFunction{
		Name:  "expand_env",
		Value: osExpandEnv,
	}, // expand_env(s string) => string
	"getegid": &tender.NativeFunction{
		Name:  "getegid",
		Value: FuncARI(os.Getegid),
	}, // getegid() => int
	"getenv": &tender.NativeFunction{
		Name:  "getenv",
		Value: FuncASRS(os.Getenv),
	}, // getenv(s string) => string
	"geteuid": &tender.NativeFunction{
		Name:  "geteuid",
		Value: FuncARI(os.Geteuid),
	}, // geteuid() => int
	"getgid": &tender.NativeFunction{
		Name:  "getgid",
		Value: FuncARI(os.Getgid),
	}, // getgid() => int
	"getgroups": &tender.NativeFunction{
		Name:  "getgroups",
		Value: FuncARIsE(os.Getgroups),
	}, // getgroups() => array(string)/error
	"getpagesize": &tender.NativeFunction{
		Name:  "getpagesize",
		Value: FuncARI(os.Getpagesize),
	}, // getpagesize() => int
	"getpid": &tender.NativeFunction{
		Name:  "getpid",
		Value: FuncARI(os.Getpid),
	}, // getpid() => int
	"getppid": &tender.NativeFunction{
		Name:  "getppid",
		Value: FuncARI(os.Getppid),
	}, // getppid() => int
	"getuid": &tender.NativeFunction{
		Name:  "getuid",
		Value: FuncARI(os.Getuid),
	}, // getuid() => int
	"getwd": &tender.NativeFunction{
		Name:  "getwd",
		Value: FuncARSE(os.Getwd),
	}, // getwd() => string/error
	"hostname": &tender.NativeFunction{
		Name:  "hostname",
		Value: FuncARSE(os.Hostname),
	}, // hostname() => string/error
	"lchown": &tender.NativeFunction{
		Name:  "lchown",
		Value: FuncASIIRE(os.Lchown),
	}, // lchown(name string, uid int, gid int) => error
	"link": &tender.NativeFunction{
		Name:  "link",
		Value: FuncASSRE(os.Link),
	}, // link(oldname string, newname string) => error
	"lookup_env": &tender.NativeFunction{
		Name:  "lookup_env",
		Value: osLookupEnv,
	}, // lookup_env(key string) => string/false
	"mkdir":     osFuncASFmRE("mkdir", os.Mkdir),        // mkdir(name string, perm int) => error
	"mkdir_all": osFuncASFmRE("mkdir_all", os.MkdirAll), // mkdir_all(name string, perm int) => error
	"readlink": &tender.NativeFunction{
		Name:  "readlink",
		Value: FuncASRSE(os.Readlink),
	}, // readlink(name string) => string/error
	"remove": &tender.NativeFunction{
		Name:  "remove",
		Value: FuncASRE(os.Remove),
	}, // remove(name string) => error
	"remove_all": &tender.NativeFunction{
		Name:  "remove_all",
		Value: FuncASRE(os.RemoveAll),
	}, // remove_all(name string) => error
	"rename": &tender.NativeFunction{
		Name:  "rename",
		Value: FuncASSRE(os.Rename),
	}, // rename(oldpath string, newpath string) => error
	"setenv": &tender.NativeFunction{
		Name:  "setenv",
		Value: FuncASSRE(os.Setenv),
	}, // setenv(key string, value string) => error
	"symlink": &tender.NativeFunction{
		Name:  "symlink",
		Value: FuncASSRE(os.Symlink),
	}, // symlink(oldname string newname string) => error
	"temp_dir": &tender.NativeFunction{
		Name:  "temp_dir",
		Value: FuncARS(os.TempDir),
	}, // temp_dir() => string
	"truncate": &tender.NativeFunction{
		Name:  "truncate",
		Value: FuncASI64RE(os.Truncate),
	}, // truncate(name string, size int) => error
	"unsetenv": &tender.NativeFunction{
		Name:  "unsetenv",
		Value: FuncASRE(os.Unsetenv),
	}, // unsetenv(key string) => error
	"create": &tender.NativeFunction{
		Name:  "create",
		Value: osCreate,
	}, // create(name string) => imap(file)/error
	"open": &tender.NativeFunction{
		Name:  "open",
		Value: osOpen,
	}, // open(name string) => imap(file)/error
	"open_file": &tender.NativeFunction{
		Name:  "open_file",
		Value: osOpenFile,
	}, // open_file(name string, flag int, perm int) => imap(file)/error
	"find_process": &tender.NativeFunction{
		Name:  "find_process",
		Value: osFindProcess,
	}, // find_process(pid int) => imap(process)/error
	"start_process": &tender.NativeFunction{
		Name:  "start_process",
		Value: osStartProcess,
	}, // start_process(name string, argv array(string), dir string, env array(string)) => imap(process)/error
	"exec_look_path": &tender.NativeFunction{
		Name:  "exec_look_path",
		Value: FuncASRSE(exec.LookPath),
	}, // exec_look_path(file) => string/error
	"exec": &tender.NativeFunction{
		Name:  "exec",
		Value: osExec,
	}, // exec(name, args...) => command
	"stat": &tender.NativeFunction{
		Name:  "stat",
		Value: osStat,
	}, // stat(name) => imap(fileinfo)/error
	"read_file": &tender.NativeFunction{
		Name:  "read_file",
		Value: osReadFile,
	}, // readfile(name) => array(byte)/error
	"read_dir": &tender.NativeFunction{
		Name:  "read_dir",
		Value: osReadDir,
	},
}

func osReadDir(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}

	dirname, ok := tender.ToString(args[0])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return wrapError(err), nil
	}

	var resultArray tender.Array
	for _, file := range files {
		fileInfo := &tender.ImmutableMap{
			Value: map[string]tender.Object{
				"name":  &tender.String{Value: file.Name()},
				"size":  &tender.Int{Value: file.Size()},
				"mode":  &tender.Int{Value: int64(file.Mode())},
				"mtime": &tender.Time{Value: file.ModTime()},
			},
		}

		if file.IsDir() {
			fileInfo.Value["is_dir"] = tender.TrueValue
		} else {
			fileInfo.Value["is_dir"] = tender.FalseValue
		}

		resultArray.Value = append(resultArray.Value, fileInfo)
	}

	return &resultArray, nil
}

func osReadFile(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}
	fname, ok := tender.ToString(args[0])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	bytes, err := ioutil.ReadFile(fname)
	if err != nil {
		return wrapError(err), nil
	}
	if len(bytes) > tender.MaxBytesLen {
		return nil, tender.ErrBytesLimit
	}
	return &tender.Bytes{Value: bytes}, nil
}

func osStat(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}
	fname, ok := tender.ToString(args[0])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	stat, err := os.Stat(fname)
	if err != nil {
		return wrapError(err), nil
	}
	fstat := &tender.ImmutableMap{
		Value: map[string]tender.Object{
			"name":  &tender.String{Value: stat.Name()},
			"mtime": &tender.Time{Value: stat.ModTime()},
			"size":  &tender.Int{Value: stat.Size()},
			"mode":  &tender.Int{Value: int64(stat.Mode())},
		},
	}
	if stat.IsDir() {
		fstat.Value["directory"] = tender.TrueValue
	} else {
		fstat.Value["directory"] = tender.FalseValue
	}
	return fstat, nil
}

func osCreate(args ...tender.Object) (tender.Object, error) {
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
	res, err := os.Create(s1)
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSFile(res), nil
}

func osCopy(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 2 {
		return nil, tender.ErrWrongNumArguments
	}

	src, ok := tender.ToString(args[0])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}

	dest, ok := tender.ToString(args[1])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
	}

	err = copyFileOrDir(src, dest)
	if err != nil {
		return wrapError(err), nil
	}

	return tender.NullValue, nil
}

func copyFileOrDir(src, dest string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	if srcInfo.IsDir() {
		return copyDir(src, dest)
	}

	return copyFile(src, dest)
}

func copyDir(src, dest string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dest, srcInfo.Mode())
	if err != nil {
		return err
	}

	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		err = copyFileOrDir(srcPath, destPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func copyFile(src, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}


func osOpen(args ...tender.Object) (tender.Object, error) {
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
	res, err := os.Open(s1)
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSFile(res), nil
}

func osOpenFile(args ...tender.Object) (tender.Object, error) {
	if len(args) != 3 {
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
	i2, ok := tender.ToInt(args[1])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
	}
	i3, ok := tender.ToInt(args[2])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
	}
	res, err := os.OpenFile(s1, i2, os.FileMode(i3))
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSFile(res), nil
}

func osArgs(args ...tender.Object) (tender.Object, error) {
	vm := args[0].(*tender.VMObj).Value
	args = args[1:] // the first arg is VMObj inserted by VM
	// fmt.Println(vm.CallCompiledFunc)
	if len(args) != 0 {
		return nil, tender.ErrWrongNumArguments
	}
	arr := &tender.Array{}
	for _, osArg := range vm.Args {
		if len(osArg) > tender.MaxStringLen {
			return nil, tender.ErrStringLimit
		}
		arr.Value = append(arr.Value, &tender.String{Value: osArg})
	}
	return arr, nil
}

func osFuncASFmRE(
	name string,
	fn func(string, os.FileMode) error,
) *tender.NativeFunction {
	return &tender.NativeFunction{
		Name: name,
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
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
			i2, ok := tender.ToInt64(args[1])
			if !ok {
				return nil, tender.ErrInvalidArgumentType{
					Name:     "second",
					Expected: "int(compatible)",
					Found:    args[1].TypeName(),
				}
			}
			return wrapError(fn(s1, os.FileMode(i2))), nil
		},
	}
}

func osLookupEnv(args ...tender.Object) (tender.Object, error) {
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
	res, ok := os.LookupEnv(s1)
	if !ok {
		return tender.FalseValue, nil
	}
	if len(res) > tender.MaxStringLen {
		return nil, tender.ErrStringLimit
	}
	return &tender.String{Value: res}, nil
}

func osExpandEnv(args ...tender.Object) (tender.Object, error) {
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
	var vlen int
	var failed bool
	s := os.Expand(s1, func(k string) string {
		if failed {
			return ""
		}
		v := os.Getenv(k)

		// this does not count the other texts that are not being replaced
		// but the code checks the final length at the end
		vlen += len(v)
		if vlen > tender.MaxStringLen {
			failed = true
			return ""
		}
		return v
	})
	if failed || len(s) > tender.MaxStringLen {
		return nil, tender.ErrStringLimit
	}
	return &tender.String{Value: s}, nil
}

func osExec(args ...tender.Object) (tender.Object, error) {
	if len(args) == 0 {
		return nil, tender.ErrWrongNumArguments
	}
	name, ok := tender.ToString(args[0])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	var execArgs []string
	for idx, arg := range args[1:] {
		execArg, ok := tender.ToString(arg)
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     fmt.Sprintf("args[%d]", idx),
				Expected: "string(compatible)",
				Found:    args[1+idx].TypeName(),
			}
		}
		execArgs = append(execArgs, execArg)
	}
	return makeOSExecCommand(exec.Command(name, execArgs...)), nil
}

func osFindProcess(args ...tender.Object) (tender.Object, error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}
	i1, ok := tender.ToInt(args[0])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	proc, err := os.FindProcess(i1)
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSProcess(proc), nil
}

func osStartProcess(args ...tender.Object) (tender.Object, error) {
	if len(args) != 4 {
		return nil, tender.ErrWrongNumArguments
	}
	name, ok := tender.ToString(args[0])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	var argv []string
	var err error
	switch arg1 := args[1].(type) {
	case *tender.Array:
		argv, err = stringArray(arg1.Value, "second")
		if err != nil {
			return nil, err
		}
	case *tender.ImmutableArray:
		argv, err = stringArray(arg1.Value, "second")
		if err != nil {
			return nil, err
		}
	default:
		return nil, tender.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "array",
			Found:    arg1.TypeName(),
		}
	}

	dir, ok := tender.ToString(args[2])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "string(compatible)",
			Found:    args[2].TypeName(),
		}
	}

	var env []string
	switch arg3 := args[3].(type) {
	case *tender.Array:
		env, err = stringArray(arg3.Value, "fourth")
		if err != nil {
			return nil, err
		}
	case *tender.ImmutableArray:
		env, err = stringArray(arg3.Value, "fourth")
		if err != nil {
			return nil, err
		}
	default:
		return nil, tender.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "array",
			Found:    arg3.TypeName(),
		}
	}

	proc, err := os.StartProcess(name, argv, &os.ProcAttr{
		Dir: dir,
		Env: env,
	})
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSProcess(proc), nil
}

func stringArray(arr []tender.Object, argName string) ([]string, error) {
	var sarr []string
	for idx, elem := range arr {
		str, ok := elem.(*tender.String)
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     fmt.Sprintf("%s[%d]", argName, idx),
				Expected: "string",
				Found:    elem.TypeName(),
			}
		}
		sarr = append(sarr, str.Value)
	}
	return sarr, nil
}

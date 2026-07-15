package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	// "strconv"
	"encoding/json"
	"encoding/gob"
	
	_ "embed"

	"github.com/2dprototype/tender"
	"github.com/2dprototype/tender/parser"
	"github.com/2dprototype/tender/stdlib"
	"github.com/2dprototype/tender/v/colorable"
	"runtime"
)

func init() {
	runtime.LockOSThread()
	gob.Register(&parser.SourceFileSet{})
	gob.Register(&parser.SourceFile{})
	gob.Register(&tender.Array{})
	gob.Register(&tender.Bool{})
	gob.Register(&tender.Bytes{})
	gob.Register(&tender.Char{})
	gob.Register(&tender.Function{})
	gob.Register(&tender.Error{})
	gob.Register(&tender.Int{})
	gob.Register(&tender.Float{})
	gob.Register(&tender.BigInt{})
	gob.Register(&tender.BigFloat{})
	gob.Register(&tender.Complex{})
	gob.Register(&tender.ImmutableArray{})
	gob.Register(&tender.ImmutableMap{})
	gob.Register(&tender.Map{})
	gob.Register(&tender.String{})
	gob.Register(&tender.Time{})
	gob.Register(&tender.Null{})
	gob.Register(&tender.NativeFunction{})
	gob.Register(&tender.Tuple{})
	gob.Register(&tender.Struct{})
	gob.Register(&tender.StructType{})
	gob.Register(&tender.BoundMethod{})
	gob.Register(&stdlib.IOWriter{})
	gob.Register(&stdlib.IOReader{})
	// gob.Register(&colorable.writer{})
	// gob.Register((*os.File)(nil))
}

const (
	sourceFileExt = ".td"
	replPrompt    = ">> "
)

var (
	compileOutput  string
	parseOutput    string
	showHelp       bool
	showVersion    bool
	resolvePath    bool
	// version       = "v1.0.0"
)

//go:embed version.txt
var version string

var isAnsiSupportedTerminal = false

func init() {
	if isTerminal(os.Stdout) {
        isAnsiSupportedTerminal = true
    }
	flag.BoolVar(&showHelp, "help", false, "Show help")
	flag.StringVar(&compileOutput, "o", "", "Compile output file")
	flag.StringVar(&compileOutput, "c", "", "Compile output file")
	flag.StringVar(&parseOutput, "parse", "", "Parse output file")
	flag.BoolVar(&showVersion, "version", false, "Show version")
	flag.BoolVar(&showVersion, "v", false, "Show version")
	flag.BoolVar(&resolvePath, "resolve", true, "Resolve relative import paths")
	flag.Parse()
}

func main() {
	if showHelp {
		doHelp()
		os.Exit(2)
	} else if showVersion {
		fmt.Println(version)
		return
	}

	modules := stdlib.GetModuleMap(stdlib.AllModuleNames()...)
	inputFile := flag.Arg(0)
	if inputFile == "" {
		// REPL
		RunREPL(modules, os.Stdin, os.Stdout)
		return
	}
	
	inputData, err := ioutil.ReadFile(inputFile)
	if err != nil {
		printError(string(err.Error()))
		os.Exit(1)
	}

	inputFile, err = filepath.Abs(inputFile)
	if err != nil {
		printError(string(err.Error()))
		os.Exit(1)
	}
	
	// Configure the global path resolver
	if resolvePath && inputFile != "" {
		importDir := filepath.Dir(inputFile)
		tender.SetPathResolution(true, importDir)
	} else {
		tender.SetPathResolution(false, "")
	}

	if len(inputData) > 1 && string(inputData[:2]) == "#!" {
		copy(inputData, "//")
	}
	
	if parseOutput != "" {
		err := ParseOnly(inputData, inputFile, parseOutput)
		if err != nil {
			printError(string(err.Error()))
			os.Exit(1)
		}
		return
	}

	if compileOutput != "" {
		err := CompileOnly(modules, inputData, inputFile, compileOutput)
		if err != nil {
			printError(string(err.Error()))
			os.Exit(1)
		}
	} else if filepath.Ext(inputFile) == sourceFileExt {
		err := CompileAndRun(modules, inputData, inputFile)
		if err != nil {
			printError(string(err.Error()))
			os.Exit(1)
		}
	} else {
		if err := RunCompiled(modules, inputData); err != nil {
			printError(string(err.Error()))
			os.Exit(1)
		}
	}
}


func ParseOnly(src []byte, inputFile, outputFile string) (err error) {
	if filepath.Ext(outputFile) != ".json" {
		outputFile = outputFile + ".json"
	}
	out, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			_ = out.Close()
		} else {
			err = out.Close()
		}
	}()
	
	fileSet := parser.NewFileSet()
	srcFile := fileSet.AddFile(filepath.Base(inputFile), -1, len(src))
	p := parser.NewParser(srcFile, src, nil)
	file, err := p.ParseFile()
	if err != nil {
		return err
	}
	
    jsonData, err := json.MarshalIndent(file, "", "\t")
    if err != nil {
        return err
    }
    _, err = out.Write(jsonData)
	
	if err != nil { 
		return err
	}

	fmt.Println(outputFile)
	return nil
}

// CompileOnly compiles the source code and writes the compiled binary into
// outputFile.
func CompileOnly(modules *tender.ModuleMap, data []byte, inputFile, outputFile string) (err error) {
	bytecode, err := compileSrc(modules, data, inputFile)
	if err != nil {
		return
	}
	
	isTDC := true
	if filepath.Ext(outputFile) == ".tdo" {
		isTDC = false
	}
	if filepath.Ext(outputFile) == "." || filepath.Ext(outputFile) == "" {
		outputFile = inputFile[:len(inputFile)-len(filepath.Ext(inputFile))] + ".tdc"
		isTDC = true
	}

	out, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			_ = out.Close()
		} else {
			err = out.Close()
		}
	}()
	
	if isTDC {
		err = bytecode.EncodeTDC(out, version)
	} else {
		err = bytecode.Encode(out)
	}
	if err != nil {
		return
	}
	fmt.Println(outputFile)
	return
}

// CompileAndRun compiles the source code and executes it.
func CompileAndRun(modules *tender.ModuleMap, data []byte, inputFile string) (err error) {
	bytecode, err := compileSrc(modules, data, inputFile)
	if err != nil {
		return
	}

	machine := tender.NewVM(bytecode, nil, -1)
	err = machine.Run()
	return
}

// RunCompiled reads the compiled binary from file and executes it.
func RunCompiled(modules *tender.ModuleMap, data []byte) (err error) {
	bytecode := &tender.Bytecode{}
	err = bytecode.Decode(bytes.NewReader(data), modules)
	if err != nil {
		return
	}

	machine := tender.NewVM(bytecode, nil, -1)
	err = machine.Run()
	return
}

// RunREPL starts REPL.
func RunREPL(modules *tender.ModuleMap, in io.Reader, out io.Writer) {
	stdin := bufio.NewScanner(in)
	fileSet := parser.NewFileSet()
	globals := make([]tender.Object, tender.GlobalsSize)
	symbolTable := tender.NewSymbolTable()
	for idx, fn := range tender.GetAllBuiltinFunctions() {
		symbolTable.DefineBuiltin(idx, fn.Name)
	}

	// embed println function
	symbol := symbolTable.Define("__repl_println__")
	globals[symbol.Index] = &tender.NativeFunction{
		Name: "__repl_println__",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (ret tender.Object, err error) { 
			vm := args[0].(*tender.VMObj).Value
			args = args[1:] 
			if isAnsiSupportedTerminal {
				str := ""
				for i, arg := range args {
					str += tender.ToStringPrettyColored(vm, arg)
					if i < len(args) - 1 {
						str += " "
					}
				}
				fmt.Fprintln(colorable.NewColorableStdout(), str)
				return
			}
			str := ""
			for i, arg := range args {
				str += tender.ToStringPretty(vm, arg)
				if i < len(args) - 1 {
					str += " "
				}
			}
			fmt.Println(str)
			return
		},
	}
	
	fmt.Println("tender " + version + " (REPL)")
	fmt.Println(`Type ".exit" to end the program`)

	var constants []tender.Object
	for {
		_, _ = fmt.Fprint(out, replPrompt)
		scanned := stdin.Scan()
		if !scanned {
			return
		}

		line := stdin.Text()
		if line == ".exit" {
			os.Exit(0)
		} else if line == ".stdlib" {
			for k, _ := range stdlib.BuiltinModules {
				fmt.Println("stdlib." + k)
			}
			continue
		}
		srcFile := fileSet.AddFile("repl", -1, len(line))
		p := parser.NewParser(srcFile, []byte(line), nil)
		file, err := p.ParseFile()
		if err != nil {
			printError(string(err.Error()))
			continue
		}

		file = addPrints(file)
		c := tender.NewCompiler(srcFile, symbolTable, constants, modules, nil)
		c.EnableFileImport(true)
		if err := c.Compile(file); err != nil {
			printError(string(err.Error()))
			// _, _ = fmt.Fprintln(out, err.Error())
			continue
		}

		bytecode := c.Bytecode()
		machine := tender.NewVM(bytecode, globals, -1)
		if err := machine.Run(); err != nil {
			printError(string(err.Error()))
			continue
		}
		constants = bytecode.Constants
	}
}

func compileSrc(modules *tender.ModuleMap, src []byte, inputFile string) (*tender.Bytecode, error) {	
	fileSet := parser.NewFileSet()
	srcFile := fileSet.AddFile(filepath.Base(inputFile), -1, len(src))

	p := parser.NewParser(srcFile, src, nil)
	file, err := p.ParseFile()
	if err != nil {
		return nil, err
	}

	c := tender.NewCompiler(srcFile, nil, nil, modules, nil)
	c.EnableFileImport(true)
	if resolvePath {
		c.SetImportDir(filepath.Dir(inputFile))
	}

	if err := c.Compile(file); err != nil {
		return nil, err
	}

	bytecode := c.Bytecode()
	bytecode.RemoveDuplicates()
	return bytecode, nil
}

func doHelp() {
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("    tender [flags] {input-file}")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println()
	fmt.Println("    -c, -o    compile output file")
	fmt.Println("              Specify the name of the output file when compiling.")
	fmt.Println("    -version  show version")
	fmt.Println("              Display the current version of the Tender tool.")
	fmt.Println("    -v        show version")
	fmt.Println("              Alias for -version. Display the current version of the Tender tool.")
	fmt.Println("    -parse    parse file")
	fmt.Println("              Parse the input file and display the parsed structure.")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println()
	fmt.Println("    tender")
	fmt.Println()
	fmt.Println("              Start the Tender REPL (Read-Eval-Print Loop) environment.")
	fmt.Println()
	fmt.Println("    tender myapp.td")
	fmt.Println()
	fmt.Println("              Compile and run the source file (myapp.td).")
	fmt.Println("              The source file must have a .td extension.")
	fmt.Println()
	fmt.Println("    tender -o myapp myapp.td")
	fmt.Println()
	fmt.Println("              Compile the source file (myapp.td) into a bytecode file (myapp).")
	fmt.Println()
	fmt.Println("    tender myapp")
	fmt.Println()
	fmt.Println("              Run the compiled bytecode file (myapp).")
	fmt.Println()
}

func addPrints(file *parser.File) *parser.File {
	var stmts []parser.Stmt
	for _, s := range file.Stmts {
		switch s := s.(type) {
		case *parser.ExprStmt:
			stmts = append(stmts, &parser.ExprStmt{
				Expr: &parser.CallExpr{
					Func: &parser.Ident{Name: "__repl_println__"},
					Args: []parser.Expr{s.Expr},
				},
			})
		case *parser.AssignStmt:
			stmts = append(stmts, s)

			stmts = append(stmts, &parser.ExprStmt{
				Expr: &parser.CallExpr{
					Func: &parser.Ident{
						Name: "__repl_println__",
					},
					Args: s.LHS,
				},
			})
		default:
			stmts = append(stmts, s)
		}
	}
	return &parser.File{
		InputFile: file.InputFile,
		Stmts:     stmts,
	}
}

func basename(s string) string {
	s = filepath.Base(s)
	n := strings.LastIndexByte(s, '.')
	if n > 0 {
		return s[:n]
	}
	return s
}


func printError(e string) {
	if isAnsiSupportedTerminal {
		fmt.Fprintln(colorable.NewColorableStdout(), "\033[0;31m" + e + "\033[0m")
	} else {
		fmt.Println(e)
	}
}

func isTerminal(f *os.File) bool {
    // Get file mode
    mode, err := f.Stat()
    if err != nil {
        return false
    }
    // Check if it's a character device
    return mode.Mode()&os.ModeCharDevice != 0
}

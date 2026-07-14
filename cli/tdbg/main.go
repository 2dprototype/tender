package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/2dprototype/tender"
	"github.com/2dprototype/tender/parser"
	"github.com/2dprototype/tender/stdlib"
	"github.com/2dprototype/tender/v/colorable"
)

func init() {
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
}

var (
	showHelp    bool
	showVersion bool
)

type Debugger struct {
	vm         *tender.VM
	bytecode   *tender.Bytecode
	modules    *tender.ModuleMap
	inputFile  string
	isCompiled bool

	breakpoints   map[string]map[int]int // filename -> line -> ip
	ipBreakpoints map[int]*BreakpointInfo

	currentIP   int
	currentFile string
	currentLine int
	currentFn   *tender.Function
	stackFrames []FrameInfo
	lastCommand string
	running     bool

	out io.Writer
	in  *bufio.Reader

	isTerminal bool
}

type BreakpointInfo struct {
	File string
	Line int
	IP   int
	Fn   *tender.Function
}

type FrameInfo struct {
	Fn       *tender.Function
	IP       int
	Locals   map[string]tender.Object
	FreeVars map[string]tender.Object
	File     string
	Line     int
}

type DebugCommand struct {
	Name    string
	Aliases []string
	Help    string
	MinArgs int
	MaxArgs int
	Handler func(*Debugger, []string) error
}

var debugCommands []*DebugCommand

func initCommands() {
	debugCommands = []*DebugCommand{
		{
			Name:    "continue",
			Aliases: []string{"c", "cont"},
			Help:    "Continue execution until next breakpoint",
			Handler: cmdContinue,
		},
		{
			Name:    "step",
			Aliases: []string{"s"},
			Help:    "Step into next instruction (function calls)",
			Handler: cmdStep,
		},
		{
			Name:    "next",
			Aliases: []string{"n"},
			Help:    "Step over to next line (skip function calls)",
			Handler: cmdNext,
		},
		{
			Name:    "stepout",
			Aliases: []string{"so"},
			Help:    "Step out of current function",
			Handler: cmdStepOut,
		},
		{
			Name:    "break",
			Aliases: []string{"b"},
			Help:    "Set breakpoint: break <file:line> or break <line>",
			Handler: cmdBreak,
		},
		{
			Name:    "clear",
			Aliases: []string{"cl", "delete"},
			Help:    "Clear breakpoint: clear <file:line> or clear <line> or clear all",
			Handler: cmdClear,
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Help:    "List breakpoints",
			Handler: cmdList,
		},
		{
			Name:    "print",
			Aliases: []string{"p"},
			Help:    "Print variable or expression: print <name>",
			Handler: cmdPrint,
		},
		{
			Name:    "locals",
			Aliases: []string{"lv", "vars"},
			Help:    "Show local variables in current frame",
			Handler: cmdLocals,
		},
		{
			Name:    "globals",
			Aliases: []string{"gv"},
			Help:    "Show global variables",
			Handler: cmdGlobals,
		},
		{
			Name:    "stack",
			Aliases: []string{"bt", "backtrace"},
			Help:    "Show call stack",
			Handler: cmdStack,
		},
		{
			Name:    "info",
			Aliases: []string{"i"},
			Help:    "Show current state: info [ip|file|line]",
			Handler: cmdInfo,
		},
		{
			Name:    "eval",
			Aliases: []string{"e"},
			Help:    "Evaluate expression: eval <expression>",
			Handler: cmdEval,
		},
		{
			Name:    "help",
			Aliases: []string{"h", "?"},
			Help:    "Show help",
			Handler: cmdHelp,
		},
		{
			Name:    "quit",
			Aliases: []string{"q", "exit"},
			Help:    "Quit debugger",
			Handler: cmdQuit,
		},
	}
}

func main() {
	flag.BoolVar(&showHelp, "help", false, "Show help")
	flag.BoolVar(&showVersion, "version", false, "Show version")
	flag.BoolVar(&showVersion, "v", false, "Show version")
	flag.Parse()

	if showHelp {
		printHelp()
		os.Exit(0)
	}
	if showVersion {
		printVersion()
		os.Exit(0)
	}

	if len(flag.Args()) < 1 {
		printHelp()
		os.Exit(1)
	}

	initCommands()

	inputFile := flag.Arg(0)
	modules := stdlib.GetModuleMap(stdlib.AllModuleNames()...)

	debugger, err := NewDebugger(inputFile, modules)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if err := debugger.Run(); err != nil {
		fmt.Printf("Debugger error: %v\n", err)
		os.Exit(1)
	}
}

func NewDebugger(inputFile string, modules *tender.ModuleMap) (*Debugger, error) {
	d := &Debugger{
		modules:       modules,
		inputFile:     inputFile,
		breakpoints:   make(map[string]map[int]int),
		ipBreakpoints: make(map[int]*BreakpointInfo),
		out:           colorable.NewColorableStdout(),
		in:            bufio.NewReader(os.Stdin),
		isTerminal:    isTerminal(os.Stdout),
		currentIP:     -1,
		currentLine:   0,
	}

	ext := filepath.Ext(inputFile)
	if ext == ".tdo" {
		d.isCompiled = true
		data, err := os.ReadFile(inputFile)
		if err != nil {
			return nil, fmt.Errorf("failed to read bytecode file: %v", err)
		}
		bytecode := &tender.Bytecode{}
		if err := bytecode.Decode(bytes.NewReader(data), modules); err != nil {
			return nil, fmt.Errorf("failed to decode bytecode: %v", err)
		}
		d.bytecode = bytecode
	} else {
		d.isCompiled = false
		data, err := os.ReadFile(inputFile)
		if err != nil {
			return nil, fmt.Errorf("failed to read source file: %v", err)
		}

		fileSet := parser.NewFileSet()
		srcFile := fileSet.AddFile(filepath.Base(inputFile), -1, len(data))
		p := parser.NewParser(srcFile, data, nil)
		file, err := p.ParseFile()
		if err != nil {
			return nil, fmt.Errorf("parse error: %v", err)
		}

		c := tender.NewCompiler(srcFile, nil, nil, modules, nil)
		c.EnableFileImport(true)
		c.SetImportDir(filepath.Dir(inputFile))
		if err := c.Compile(file); err != nil {
			return nil, fmt.Errorf("compile error: %v", err)
		}

		bytecode := c.Bytecode()
		bytecode.RemoveDuplicates()
		d.bytecode = bytecode
	}

	d.vm = tender.NewVM(d.bytecode, nil, -1)
	return d, nil
}

func (d *Debugger) Run() error {
	d.printHeader()
	d.running = true

	d.setBreakpointOnEntry()

	if err := d.runUntilBreakpoint(); err != nil {
		return err
	}

	for d.running {
		if err := d.printState(); err != nil {
			return err
		}
		if err := d.processCommand(); err != nil {
			if err == ErrQuit {
				return nil
			}
			return err
		}
	}

	return nil
}

var ErrQuit = fmt.Errorf("quit")

func (d *Debugger) runUntilBreakpoint() error {
	return d.singleStepLoop()
}

func (d *Debugger) singleStepLoop() error {
	for {
		d.updateCurrentState()

		if bp := d.getBreakpointAtIP(d.currentIP, d.currentFn); bp != nil {
			d.currentFile = bp.File
			d.currentLine = bp.Line
			return nil
		}

		if err := d.executeSingleStep(); err != nil {
			return err
		}

		if d.vm.IsStackEmpty() {
			d.print("Program finished")
			d.running = false
			return nil
		}
	}
}

func (d *Debugger) executeSingleStep() error {
	return d.vm.Run()
}

func (d *Debugger) updateCurrentState() {
	if d.vm == nil {
		return
	}

	d.currentIP = d.vm.GetIP()
	d.currentFn = d.vm.GetCurrentFunction()

	if d.currentFn != nil && d.currentIP >= 0 {
		pos := d.currentFn.SourcePos(d.currentIP)
		if pos != parser.NoPos && d.bytecode.FileSet != nil {
			if f := d.bytecode.FileSet.File(pos); f != nil {
				d.currentFile = f.Name
				// Use Position() method which returns SourceFilePos with Line field
				filePos := f.Position(pos)
				if filePos.IsValid() {
					d.currentLine = filePos.Line
				}
			}
		}
	}
}

func countLines(data []byte) int {
	if len(data) == 0 {
		return 0
	}
	lines := 0
	for _, b := range data {
		if b == '\n' {
			lines++
		}
	}
	return lines
}

func (d *Debugger) getBreakpointAtIP(ip int, fn *tender.Function) *BreakpointInfo {
	if bp, ok := d.ipBreakpoints[ip]; ok {
		if bp.Fn == fn || bp.Fn == nil {
			return bp
		}
	}
	return nil
}

func (d *Debugger) setBreakpointOnEntry() {
	if d.bytecode.MainFunction != nil {
		d.addIPBreakpoint(0, d.bytecode.MainFunction, "(main)", 0)
	}
}

func (d *Debugger) addIPBreakpoint(ip int, fn *tender.Function, file string, line int) {
	info := &BreakpointInfo{
		IP:   ip,
		Fn:   fn,
		File: file,
		Line: line,
	}
	d.ipBreakpoints[ip] = info

	if file != "" {
		if _, ok := d.breakpoints[file]; !ok {
			d.breakpoints[file] = make(map[int]int)
		}
		d.breakpoints[file][line] = ip
	}
}

func (d *Debugger) removeIPBreakpoint(ip int) {
	if bp, ok := d.ipBreakpoints[ip]; ok {
		if bp.File != "" {
			if fileBps, ok := d.breakpoints[bp.File]; ok {
				delete(fileBps, bp.Line)
			}
		}
		delete(d.ipBreakpoints, ip)
	}
}

func (d *Debugger) printHeader() {
	header := `
╔═══════════════════════════════════════════════════════════════╗
║                      Tender Debugger                        ║
╚═══════════════════════════════════════════════════════════════╝
`
	d.print(header)
	d.printColored("cyan", "Debugging: "+d.inputFile)
	if d.isCompiled {
		d.printColored("yellow", "  [Compiled bytecode]")
	} else {
		d.printColored("yellow", "  [Source file]")
	}
	d.print("Type 'help' for commands, 'quit' to exit")
	d.print("")
}

func (d *Debugger) printState() error {
	d.printColored("green", "=== Debugger State ===")

	if d.currentFile != "" {
		d.print(fmt.Sprintf("  File: %s", filepath.Base(d.currentFile)))
	}
	if d.currentLine > 0 {
		d.print(fmt.Sprintf("  Line: %d", d.currentLine))
	}
	if d.currentIP >= 0 {
		d.print(fmt.Sprintf("  IP:   %d", d.currentIP))
	}
	if d.currentFn != nil {
		d.print(fmt.Sprintf("  Function: <function> (params: %d, locals: %d)",
			d.currentFn.NumParameters,
			d.currentFn.NumLocals))
	}

	if d.currentFile != "" && d.currentLine > 0 {
		line, err := d.getSourceLine(d.currentFile, d.currentLine)
		if err == nil && line != "" {
			d.print("")
			d.printColored("yellow", fmt.Sprintf("  > %s", strings.TrimSpace(line)))
		}
	}

	d.print("")

	if len(d.ipBreakpoints) > 0 {
		d.printColored("magenta", "Breakpoints:")
		for _, bp := range d.ipBreakpoints {
			loc := "entry"
			if bp.File != "" && bp.Line > 0 {
				loc = fmt.Sprintf("%s:%d", filepath.Base(bp.File), bp.Line)
			}
			d.print(fmt.Sprintf("  [%d] %s", bp.IP, loc))
		}
	}

	return nil
}

func (d *Debugger) getSourceLine(file string, line int) (string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(data), "\n")
	if line <= len(lines) {
		return strings.TrimRight(lines[line-1], "\r\n"), nil
	}
	return "", fmt.Errorf("line %d not found", line)
}

func (d *Debugger) processCommand() error {
	fmt.Print("(tdb) ")
	input, err := d.in.ReadString('\n')
	if err != nil {
		return err
	}
	input = strings.TrimSpace(input)
	if input == "" {
		return nil
	}

	parts := strings.Fields(input)
	if len(parts) == 0 {
		return nil
	}
	cmdName := parts[0]
	args := parts[1:]

	for _, cmd := range debugCommands {
		if strings.EqualFold(cmd.Name, cmdName) {
			if len(args) < cmd.MinArgs {
				return fmt.Errorf("not enough arguments for %s", cmd.Name)
			}
			if cmd.MaxArgs >= 0 && len(args) > cmd.MaxArgs {
				return fmt.Errorf("too many arguments for %s", cmd.Name)
			}
			return cmd.Handler(d, args)
		}
		for _, alias := range cmd.Aliases {
			if strings.EqualFold(alias, cmdName) {
				if len(args) < cmd.MinArgs {
					return fmt.Errorf("not enough arguments for %s", cmd.Name)
				}
				if cmd.MaxArgs >= 0 && len(args) > cmd.MaxArgs {
					return fmt.Errorf("too many arguments for %s", cmd.Name)
				}
				return cmd.Handler(d, args)
			}
		}
	}

	return fmt.Errorf("unknown command: %s", cmdName)
}

// Command Handlers

func cmdContinue(d *Debugger, args []string) error {
	d.print("Continuing...")
	return d.runUntilBreakpoint()
}

func cmdStep(d *Debugger, args []string) error {
	if err := d.executeSingleStep(); err != nil {
		return err
	}
	d.updateCurrentState()
	return nil
}

func cmdNext(d *Debugger, args []string) error {
	return cmdStep(d, args)
}

func cmdStepOut(d *Debugger, args []string) error {
	return cmdStep(d, args)
}

func cmdBreak(d *Debugger, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("specify breakpoint location: break <file:line> or break <line>")
	}

	loc := args[0]
	var file string
	var lineStr string
	var line int
	var err error

	if strings.Contains(loc, ":") {
		parts := strings.SplitN(loc, ":", 2)
		file = parts[0]
		lineStr = parts[1]
	} else {
		lineStr = loc
		file = d.currentFile
	}

	line, err = strconv.Atoi(lineStr)
	if err != nil {
		return fmt.Errorf("invalid line number: %s", lineStr)
	}
	if line <= 0 {
		return fmt.Errorf("line number must be positive")
	}

	if !filepath.IsAbs(file) && d.currentFile != "" {
		dir := filepath.Dir(d.currentFile)
		testPath := filepath.Join(dir, file)
		if _, err := os.Stat(testPath); err == nil {
			file = testPath
		}
	}

	if d.currentFn != nil {
		d.addIPBreakpoint(d.currentIP, d.currentFn, file, line)
		d.print(fmt.Sprintf("Breakpoint set at %s:%d (IP: %d)", file, line, d.currentIP))
	} else {
		d.setBreakpointOnEntry()
		d.print("Breakpoint set at entry")
	}

	return nil
}

func cmdClear(d *Debugger, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("specify breakpoint to clear: clear <ip> or clear all")
	}

	arg := args[0]
	if strings.EqualFold(arg, "all") {
		d.ipBreakpoints = make(map[int]*BreakpointInfo)
		d.breakpoints = make(map[string]map[int]int)
		d.print("All breakpoints cleared")
		return nil
	}

	if ip, err := strconv.Atoi(arg); err == nil {
		if _, ok := d.ipBreakpoints[ip]; ok {
			d.removeIPBreakpoint(ip)
			d.print(fmt.Sprintf("Breakpoint at IP %d cleared", ip))
			return nil
		}
		return fmt.Errorf("breakpoint at IP %d not found", ip)
	}

	loc := arg
	var file string
	var lineStr string
	var line int

	if strings.Contains(loc, ":") {
		parts := strings.SplitN(loc, ":", 2)
		file = parts[0]
		lineStr = parts[1]
	} else {
		lineStr = loc
		file = d.currentFile
	}

	line, err := strconv.Atoi(lineStr)
	if err != nil {
		return fmt.Errorf("invalid location: %s", arg)
	}

	if fileBps, ok := d.breakpoints[file]; ok {
		if ip, ok := fileBps[line]; ok {
			d.removeIPBreakpoint(ip)
			d.print(fmt.Sprintf("Breakpoint at %s:%d cleared", file, line))
			return nil
		}
	}

	return fmt.Errorf("breakpoint at %s:%d not found", file, line)
}

func cmdList(d *Debugger, args []string) error {
	if len(d.ipBreakpoints) == 0 {
		d.print("No breakpoints set")
		return nil
	}

	d.printColored("magenta", "Breakpoints:")
	for ip, bp := range d.ipBreakpoints {
		loc := "entry"
		if bp.File != "" && bp.Line > 0 {
			loc = fmt.Sprintf("%s:%d", filepath.Base(bp.File), bp.Line)
		}
		d.print(fmt.Sprintf("  %d: %s", ip, loc))
	}
	return nil
}

func cmdPrint(d *Debugger, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("specify variable to print")
	}

	name := args[0]
	val := d.findVariable(name)
	if val != nil {
		d.print(fmt.Sprintf("%s = %s", name, val.String()))
	} else {
		d.print(fmt.Sprintf("'%s' not found", name))
	}
	return nil
}

func cmdLocals(d *Debugger, args []string) error {
	locals := d.getCurrentLocals()
	if len(locals) == 0 {
		d.print("No local variables")
		return nil
	}

	d.printColored("cyan", "Local variables:")
	for name, val := range locals {
		d.print(fmt.Sprintf("  %s = %s", name, val.String()))
	}
	return nil
}

func cmdGlobals(d *Debugger, args []string) error {
	globals := d.getGlobals()
	if len(globals) == 0 {
		d.print("No global variables")
		return nil
	}

	d.printColored("green", "Global variables:")
	for name, val := range globals {
		if name != "" {
			d.print(fmt.Sprintf("  %s = %s", name, val.String()))
		}
	}
	return nil
}

func cmdStack(d *Debugger, args []string) error {
	frames := d.getStackFrames()
	if len(frames) == 0 {
		d.print("No stack frames")
		return nil
	}

	d.printColored("yellow", "Call stack:")
	for i, frame := range frames {
		marker := ""
		if i == 0 {
			marker = " (current)"
		}
		loc := "?"
		if frame.File != "" && frame.Line > 0 {
			loc = fmt.Sprintf("%s:%d", filepath.Base(frame.File), frame.Line)
		}
		d.print(fmt.Sprintf("  #%d <function> at %s%s", i, loc, marker))
	}
	return nil
}

func cmdInfo(d *Debugger, args []string) error {
	if len(args) == 0 {
		return d.printState()
	}

	switch strings.ToLower(args[0]) {
	case "ip":
		d.print(fmt.Sprintf("Current IP: %d", d.currentIP))
	case "file":
		d.print(fmt.Sprintf("Current file: %s", d.currentFile))
	case "line":
		d.print(fmt.Sprintf("Current line: %d", d.currentLine))
	default:
		return fmt.Errorf("unknown info: %s", args[0])
	}
	return nil
}

func cmdEval(d *Debugger, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("specify expression to evaluate")
	}
	expr := strings.Join(args, " ")
	d.print(fmt.Sprintf("Evaluating: %s", expr))
	d.print("Expression evaluation not fully implemented yet")
	return nil
}

func cmdHelp(d *Debugger, args []string) error {
	if len(args) == 0 {
		d.printColored("cyan", "Available commands:")
		for _, cmd := range debugCommands {
			aliases := ""
			if len(cmd.Aliases) > 0 {
				aliases = fmt.Sprintf(" (aliases: %s)", strings.Join(cmd.Aliases, ", "))
			}
			d.print(fmt.Sprintf("  %-10s %s%s", cmd.Name, cmd.Help, aliases))
		}
		return nil
	}

	for _, cmd := range debugCommands {
		if strings.EqualFold(cmd.Name, args[0]) {
			d.print(fmt.Sprintf("%s: %s", cmd.Name, cmd.Help))
			return nil
		}
	}
	return fmt.Errorf("unknown command: %s", args[0])
}

func cmdQuit(d *Debugger, args []string) error {
	d.print("Quitting debugger...")
	d.running = false
	return ErrQuit
}

// Helper methods

func (d *Debugger) findVariable(name string) tender.Object {
	if locals := d.getCurrentLocals(); locals != nil {
		if val, ok := locals[name]; ok {
			return val
		}
	}

	if globals := d.getGlobals(); globals != nil {
		if val, ok := globals[name]; ok {
			return val
		}
	}

	return nil
}

func (d *Debugger) getCurrentLocals() map[string]tender.Object {
	return make(map[string]tender.Object)
}

func (d *Debugger) getGlobals() map[string]tender.Object {
	return make(map[string]tender.Object)
}

func (d *Debugger) getStackFrames() []FrameInfo {
	return []FrameInfo{}
}

func (d *Debugger) print(msg string) {
	fmt.Fprintln(d.out, msg)
}

func (d *Debugger) printColored(color, msg string) {
	if d.isTerminal {
		colors := map[string]string{
			"red":     "\033[0;31m",
			"green":   "\033[0;32m",
			"yellow":  "\033[0;33m",
			"blue":    "\033[0;34m",
			"magenta": "\033[0;35m",
			"cyan":    "\033[0;36m",
			"white":   "\033[0;37m",
		}
		if code, ok := colors[color]; ok {
			fmt.Fprintf(d.out, "%s%s\033[0m\n", code, msg)
			return
		}
	}
	fmt.Fprintln(d.out, msg)
}

func isTerminal(f *os.File) bool {
	mode, err := f.Stat()
	if err != nil {
		return false
	}
	return mode.Mode()&os.ModeCharDevice != 0
}

func printHelp() {
	fmt.Println("Tender Debugger")
	fmt.Println()
	fmt.Println("Usage: tdb [flags] <file>")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  -help     Show this help")
	fmt.Println("  -version  Show version")
	fmt.Println()
	fmt.Println("Debugger commands:")
	fmt.Println("  continue (c, cont)    Continue execution")
	fmt.Println("  step (s)              Step into next instruction")
	fmt.Println("  next (n)              Step over to next line")
	fmt.Println("  stepout (so)          Step out of current function")
	fmt.Println("  break (b)             Set breakpoint")
	fmt.Println("  clear (cl, delete)    Clear breakpoint")
	fmt.Println("  list (l)              List breakpoints")
	fmt.Println("  print (p)             Print variable")
	fmt.Println("  locals (lv, vars)     Show local variables")
	fmt.Println("  globals (gv)          Show global variables")
	fmt.Println("  stack (bt, backtrace) Show call stack")
	fmt.Println("  info (i)              Show current state")
	fmt.Println("  eval (e)              Evaluate expression")
	fmt.Println("  help (h, ?)           Show this help")
	fmt.Println("  quit (q, exit)        Quit debugger")
}

func printVersion() {
	fmt.Println("Tender Debugger v1.0.0")
}

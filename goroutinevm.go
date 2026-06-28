package tender

import (
	"fmt"
	"runtime/debug"
	"sync/atomic"
	"time"
)

func init() {
	addBuiltinFunction("go", builtinGovm, true)
	addBuiltinFunction("abort", builtinAbort, true)
	addBuiltinFunction("makechan", builtinMakechan, false)
}

type ret struct {
	val Object
	err error
}

type goroutineVM struct {
	*VM      // if not nil, run CompiledFunction in VM
	ret      // return value
	waitChan chan ret
	done     int64
}

// Starts a independent concurrent goroutine which runs fn(arg1, arg2, ...)
	//
	// If fn is CompiledFunction, the current running VM will be cloned to create
	// a new VM in which the CompiledFunction will be running.
	//
	// The fn can also be any object that has Call() method, such as BuiltinFunction,
	// in which case no cloned VM will be created.
	//
	// Returns a goroutineVM object that has wait, result, abort methods.
	//
	// The goroutineVM will not exit unless:
	//  1. All its descendant goroutineVMs exit
	//  2. It calls abort()
	//  3. Its goroutineVM object abort() is called on behalf of its parent VM
	// The latter 2 cases will trigger aborting procedure of all the descendant goroutineVMs,
// which will further result in #1 above.
func builtinGovm(args ...Object) (Object, error) {
	vm := args[0].(*VMObj).Value
	args = args[1:] // the first arg is VMObj inserted by VM
	if len(args) == 0 {
		return nil, ErrWrongNumArguments
	}
	
	fn := args[0]
	if !fn.CanCall() {
		return nil, ErrInvalidArgumentType{
			Name:     "first",
			Expected: "callable function",
			Found:    fn.TypeName(),
		}
	}
	
	gvm := &goroutineVM{
		waitChan: make(chan ret, 1),
	}
	
	var callers []frame
	cfn, compiled := fn.(*CompiledFunction)
	if compiled {
	gvm.VM = vm.ShallowClone()
	} else {
		callers = vm.callers()
	}
	
	if err := vm.addChild(gvm.VM); err != nil {
		return nil, err
	}
	go func() {
		var val Object
		var err error
		defer func() {
			if perr := recover(); perr != nil {
				if callers == nil {
					panic("callers not saved")
				}
				err = fmt.Errorf("\nRuntime Panic: %v%s\n%s", perr, vm.callStack(callers), debug.Stack())
			}
			if err != nil {
				vm.addError(err)
			}
			gvm.waitChan <- ret{val, err}
			vm.delChild(gvm.VM)
		gvm.VM = nil
		}()
		
		if cfn != nil {
		val, err = gvm.RunCompiled(cfn, args[1:]...)
		} else {
			var nargs []Object
			if bltnfn, ok := fn.(*BuiltinFunction); ok {
				if bltnfn.NeedVMObj {
					// pass VM as the first para to builtin functions
					nargs = append(nargs, vm.selfObject())
				}
			}
			nargs = append(nargs, args[1:]...)
			val, err = fn.Call(nargs...)
		}
		}()
		
		obj := map[string]Object{
			"result": &BuiltinFunction{Value: gvm.getRet},
			"wait":   &BuiltinFunction{Value: gvm.waitTimeout},
			"abort":  &BuiltinFunction{Value: gvm.abort},
		}
		return &Map{Value: obj}, nil
}

// Triggers the termination process of the current VM and all its descendant VMs.
func builtinAbort(args ...Object) (Object, error) {
	vm := args[0].(*VMObj).Value
	args = args[1:] // the first arg is VMObj inserted by VM
	if len(args) != 0 {
		return nil, ErrWrongNumArguments
	}
	vm.Abort() // aborts self and all descendant VMs
	return nil, nil
}

// Returns true if the goroutineVM is done
func (gvm *goroutineVM) wait(seconds int64) bool {
	if atomic.LoadInt64(&gvm.done) == 1 {
		return true
	}
	
	if seconds < 0 {
		seconds = 3153600000 // 100 years
	}
	
	select {
		case gvm.ret = <-gvm.waitChan:
		atomic.StoreInt64(&gvm.done, 1)
		case <-time.After(time.Duration(seconds) * time.Second):
		return false
	}
	
	return true
}

// Waits for the goroutineVM to complete up to timeout seconds.
	// Returns true if the goroutineVM exited(successfully or not) within the timeout.
// Waits forever if the optional timeout not specified, or timeout < 0.
func (gvm *goroutineVM) waitTimeout(args ...Object) (Object, error) {
	if len(args) > 1 {
		return nil, ErrWrongNumArguments
	}
	timeOut := -1
	if len(args) == 1 {
		t, ok := ToInt(args[0])
		if !ok {
			return nil, ErrInvalidArgumentType{
				Name:     "first",
				Expected: "int(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		timeOut = t
	}
	
	if gvm.wait(int64(timeOut)) {
		return TrueValue, nil
	}
	return FalseValue, nil
}

// Triggers the termination process of the goroutineVM and all its descendant VMs.
func (gvm *goroutineVM) abort(args ...Object) (Object, error) {
	if len(args) != 0 {
		return nil, ErrWrongNumArguments
	}
	if gvm.VM != nil {
		gvm.Abort()
	}
	return nil, nil
}

// Waits the goroutineVM to complete, return Error object if any runtime error occurred
// during the execution, otherwise return the result value of fn(arg1, arg2, ...)
func (gvm *goroutineVM) getRet(args ...Object) (Object, error) {
	if len(args) != 0 {
		return nil, ErrWrongNumArguments
	}
	
	gvm.wait(-1)
	if gvm.ret.err != nil {
		return &Error{Value: &String{Value: gvm.ret.err.Error()}}, nil
	}
	
	return gvm.ret.val, nil
}

type objchan chan Object

// Makes a channel to send/receive object
// Returns a chan object that has send, recv, close methods.
func builtinMakechan(args ...Object) (Object, error) {
	var size int
	switch len(args) {
		case 0:
		case 1:
		n, ok := ToInt(args[0])
		if !ok {
			return nil, ErrInvalidArgumentType{
				Name:     "first",
				Expected: "int(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		size = n
		default:
		return nil, ErrWrongNumArguments
	}
	
	oc := make(objchan, size)
	obj := map[string]Object{
		"send":  &BuiltinFunction{Value: oc.send, NeedVMObj: true},
		"recv":  &BuiltinFunction{Value: oc.recv, NeedVMObj: true},
		"close": &BuiltinFunction{Value: oc.close},
	}
	return &Map{Value: obj}, nil
}

// Sends an obj to the channel, will block if channel is full and the VM has not been aborted.
// Sends to a closed channel causes panic.
func (oc objchan) send(args ...Object) (Object, error) {
	vm := args[0].(*VMObj).Value
	args = args[1:] // the first arg is VMObj inserted by VM
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	select {
		case <-vm.AbortChan:
		return nil, ErrVMAborted
		case oc <- args[0]:
	}
	return nil, nil
}

// Receives an obj from the channel, will block if channel is empty and the VM has not been aborted.
// Receives from a closed channel returns null value.
func (oc objchan) recv(args ...Object) (Object, error) {
	vm := args[0].(*VMObj).Value
	args = args[1:] // the first arg is VMObj inserted by VM
	if len(args) != 0 {
		return nil, ErrWrongNumArguments
	}
	select {
		case <-vm.AbortChan:
		return nil, ErrVMAborted
		case obj, ok := <-oc:
		if ok {
			return obj, nil
		}
	}
	return nil, nil
}

// Closes the channel.
func (oc objchan) close(args ...Object) (Object, error) {
	if len(args) != 0 {
		return nil, ErrWrongNumArguments
	}
	close(oc)
	return nil, nil
}


// WrapFuncCall synchronously executes a callable object from Go space.
// It creates a shallow clone of the VM to maintain stack isolation and contains panic recovery.
func WrapFuncCall(vm *VM, args ...Object) (retVal Object, err error) {
	if len(args) == 0 {
		return nil, ErrWrongNumArguments
	}
	
	fn := args[0]
	if !fn.CanCall() {
		return nil, ErrInvalidArgumentType{
			Name:     "first",
			Expected: "callable function",
			Found:    fn.TypeName(),
		}
	}
	
	cfn, compiled := fn.(*CompiledFunction)
	clone := vm.ShallowClone()
	
	if err := vm.addChild(clone); err != nil {
		return nil, err
	}
	// Explicitly clear references on function unwind to prevent memory leak
	defer func() {
		vm.delChild(clone)
		if perr := recover(); perr != nil {
			err = fmt.Errorf("\nRuntime Panic within Native Bridge: %v\n%s", perr, debug.Stack())
			vm.addError(err)
		}
	}()
	
	// Execute Compiled Functions
	if compiled {
		return clone.RunCompiled(cfn, args[1:]...)
	}
	
	// Execute Builtin Functions
	var nargs []Object
	if bltnfn, ok := fn.(*BuiltinFunction); ok {
		if bltnfn.NeedVMObj {
			nargs = append(nargs, clone.selfObject())
		}
	}
	nargs = append(nargs, args[1:]...)
	
	return fn.Call(nargs...)
}

// // WrapFuncCall synchronously executes a callable object from Go space.
// // It creates a shallow clone of the VM to maintain stack isolation.
// func WrapFuncCall(vm *VM, args ...Object) (Object, error) {
	// if len(args) == 0 {
		// return nil, ErrWrongNumArguments
	// }
	
	// fn := args[0]
	// if !fn.CanCall() {
		// return nil, ErrInvalidArgumentType{
			// Name:     "first",
			// Expected: "callable function",
			// Found:    fn.TypeName(),
		// }
	// }
	
	// cfn, compiled := fn.(*CompiledFunction)
	
	// // Create a shallow clone for execution. 
	// // This gives us fresh frames and stack, but shares globals.
	// clone := vm.ShallowClone()
	
	// // Link the child VM to the parent so Abort() propagates correctly
	// if err := vm.addChild(clone); err != nil {
		// return nil, err
	// }
	
	// // Ensure we cleanup the child reference when execution finishes
	// defer vm.delChild(clone)
	
	// // Execute Compiled Functions
	// if compiled {
		// return clone.RunCompiled(cfn, args[1:]...)
	// }
	
	// // Execute Builtin Functions
	// var nargs []Object
	// if bltnfn, ok := fn.(*BuiltinFunction); ok {
		// if bltnfn.NeedVMObj {
			// // pass the cloned VM as the first param to builtin functions
			// nargs = append(nargs, clone.selfObject())
		// }
	// }
	// nargs = append(nargs, args[1:]...)
	
	// return fn.Call(nargs...)
// }
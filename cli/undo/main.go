package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/2dprototype/tender"
	"github.com/2dprototype/tender/parser"
	"github.com/2dprototype/tender/stdlib"
	"github.com/2dprototype/tender/token"
)

func init() {
	gob.Register(&parser.SourceFileSet{})
	gob.Register(&parser.SourceFile{})
	gob.Register(&tender.Array{})
	gob.Register(&tender.Bool{})
	gob.Register(&tender.Bytes{})
	gob.Register(&tender.Char{})
	gob.Register(&tender.CompiledFunction{})
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
	gob.Register(&tender.UserFunction{})
	gob.Register(&tender.BuiltinFunction{})
	gob.Register(&tender.Tuple{})
	gob.Register(&tender.Struct{})
	gob.Register(&tender.StructType{})
	gob.Register(&tender.BoundMethod{})
	gob.Register(&stdlib.IOWriter{})
	gob.Register(&stdlib.IOReader{})
}

var globalNames = make(map[int]string)

func getGlobalName(idx int) string {
	if name, ok := globalNames[idx]; ok {
		return name
	}
	return fmt.Sprintf("g_%d", idx)
}

func getBuiltinName(idx int) string {
	builtins := tender.GetAllBuiltinFunctions()
	if idx >= 0 && idx < len(builtins) {
		return builtins[idx].Name
	}
	return fmt.Sprintf("builtin_%d", idx)
}

func getBinaryOpStr(op int) string {
	tokens := map[int]string{
		int(token.Add):       "+",
		int(token.Sub):       "-",
		int(token.Mul):       "*",
		int(token.Quo):       "/",
		int(token.Rem):       "%",
		int(token.And):       "&",
		int(token.Or):        "|",
		int(token.Xor):       "^",
		int(token.Shl):       "<<",
		int(token.Shr):       ">>",
		int(token.AndNot):    "&^",
		int(token.Greater):   ">",
		int(token.GreaterEq): ">=",
		int(token.Spaceship): "<=>",
	}
	if str, ok := tokens[op]; ok {
		return str
	}
	return fmt.Sprintf("op_%d", op)
}

func isKeyword(s string) bool {
	keywords := map[string]bool{
		"break": true, "continue": true, "else": true, "for": true,
		"fn": true, "error": true, "immutable": true, "if": true,
		"return": true, "export": true, "embed": true, "true": true,
		"false": true, "in": true, "null": true, "import": true,
		"as": true, "from": true, "var": true, "const": true,
		"sysout": true, "type": true, "struct": true,
	}
	return keywords[s]
}

func preScanGlobals(bytecode *tender.Bytecode) {
	globalNames = make(map[int]string)

	scanFn := func(fn *tender.CompiledFunction) {
		if fn == nil {
			return
		}
		insts := fn.Instructions
		i := 0
		for i < len(insts) {
			op := insts[i]
			numOperands := parser.OpcodeOperands[op]
			operands, read := parser.ReadOperands(numOperands, insts[i+1:])
			i += 1 + read

			if op == parser.OpConstant && i < len(insts) {
				nextOp := insts[i]
				if nextOp == parser.OpSetGlobal {
					nextOperands, _ := parser.ReadOperands(parser.OpcodeOperands[parser.OpSetGlobal], insts[i+1:])
					constIdx := operands[0]
					gIdx := nextOperands[0]

					cObj := bytecode.Constants[constIdx]
					if val, ok := cObj.(*tender.ImmutableMap); ok {
						if nameObj, ok := val.Value["__module_name__"]; ok {
							if strObj, ok := nameObj.(*tender.String); ok {
								globalNames[gIdx] = strObj.Value
							}
						}
					}
				}
			}
		}
	}

	if bytecode.MainFunction != nil {
		scanFn(bytecode.MainFunction)
	}
	for _, c := range bytecode.Constants {
		if fn, ok := c.(*tender.CompiledFunction); ok {
			scanFn(fn)
		}
	}
}

type LoopInfo struct {
	HeaderIP int
	EndIP    int
	ExitIP   int
	PostIP   int
	IsIn     bool
	ItVar    int
}

type FunctionDecompiler struct {
	fn             *tender.CompiledFunction
	bytecode       *tender.Bytecode
	constants      []string
	localNames     map[int]string
	freeNames      []string
	loops          map[int]*LoopInfo
	definedGlobals map[int]bool // Tracks if a global has been declared
}

func NewFunctionDecompiler(fn *tender.CompiledFunction, bc *tender.Bytecode, freeNames []string) *FunctionDecompiler {
	fd := &FunctionDecompiler{
		fn:             fn,
		bytecode:       bc,
		constants:      make([]string, len(bc.Constants)),
		localNames:     make(map[int]string),
		freeNames:      freeNames,
		definedGlobals: make(map[int]bool),
	}

	for idx, c := range bc.Constants {
		fd.constants[idx] = fd.formatConstant(idx, c)
	}

	for i := 0; i < fn.NumParameters; i++ {
		fd.localNames[i] = fmt.Sprintf("arg_%d", i)
	}

	return fd
}

func fdFormatString(s string) string {
	if len(s) > 100 {
		hasNonPrintable := false
		for _, r := range s {
			if r > 127 || (!unicode.IsPrint(r) && !unicode.IsSpace(r)) {
				hasNonPrintable = true
				break
			}
		}
		if hasNonPrintable {
			return fmt.Sprintf("/* binary data length: %d */ \"<truncated binary data>\"", len(s))
		}
	}
	return fmt.Sprintf("%q", s)
}

func (fd *FunctionDecompiler) formatConstant(idx int, c tender.Object) string {
	if c == nil {
		return "null"
	}
	switch val := c.(type) {
	case *tender.String:
		return fdFormatString(val.Value)
	case *tender.Int:
		return fmt.Sprintf("%d", val.Value)
	case *tender.Float:
		return fmt.Sprintf("%f", val.Value)
	case *tender.BigInt:
		return fmt.Sprintf("bigint(%s)", val.Value.String())
	case *tender.BigFloat:
		return fmt.Sprintf("bigfloat(%s)", val.Value.String())
	case *tender.Complex:
		return fmt.Sprintf("complex(%g, %g)", real(val.Value), imag(val.Value))
	case *tender.Bool:
		if val.IsFalsy() {
			return "false"
		}
		return "true"
	case *tender.Bytes:
		return fmt.Sprintf("bytes(%v)", val.Value)
	case *tender.Char:
		return fmt.Sprintf("'%c'", val.Value)
	case *tender.Null:
		return "null"
	case *tender.ImmutableMap:
		if nameObj, ok := val.Value["__module_name__"]; ok {
			if strObj, ok := nameObj.(*tender.String); ok {
				return fmt.Sprintf("import_module:%s", strObj.Value)
			}
		}
		return fmt.Sprintf("immutable_map_const_%d", idx)
	case *tender.CompiledFunction:
		return fmt.Sprintf("fn_const_%d", idx)
	case *tender.StructType:
		var fields []string
		for _, f := range val.Fields {
			tagStr := ""
			if f.Tag != "" {
				tagStr = " " + fmt.Sprintf("%q", f.Tag)
			}
			fields = append(fields, fmt.Sprintf("%s %s%s", f.Name, f.Type, tagStr))
		}
		return fmt.Sprintf("struct { %s }", strings.Join(fields, "; "))
	default:
		return c.String()
	}
}

func (fd *FunctionDecompiler) getLocal(idx int) string {
	if name, ok := fd.localNames[idx]; ok {
		return name
	}
	fd.localNames[idx] = fmt.Sprintf("l_%d", idx)
	return fd.localNames[idx]
}

func (fd *FunctionDecompiler) getFree(idx int) string {
	if idx >= 0 && idx < len(fd.freeNames) {
		return fd.freeNames[idx]
	}
	return fmt.Sprintf("free_%d", idx)
}

func (fd *FunctionDecompiler) analyzeLoops() map[int]*LoopInfo {
	loops := make(map[int]*LoopInfo)
	insts := fd.fn.Instructions
	i := 0
	for i < len(insts) {
		op := insts[i]
		numOperands := parser.OpcodeOperands[op]
		operands, read := parser.ReadOperands(numOperands, insts[i+1:])
		instIP := i
		i += 1 + read

		if op == parser.OpJump {
			target := operands[0]
			if target <= instIP {
				info := &LoopInfo{
					HeaderIP: target,
					EndIP:    instIP,
				}

				if target < len(insts) {
					op2 := insts[target]
					if op2 == parser.OpGetLocal || op2 == parser.OpGetGlobal {
						opLen := 2
						if op2 == parser.OpGetGlobal {
							opLen = 3
						}
						nextIP := target + opLen
						if nextIP+1 < len(insts) && insts[nextIP] == parser.OpIteratorNext {
							jumpFalsyIP := nextIP + 1
							if jumpFalsyIP < len(insts) && insts[jumpFalsyIP] == parser.OpJumpFalsy {
								opsJF, _ := parser.ReadOperands(parser.OpcodeOperands[parser.OpJumpFalsy], insts[jumpFalsyIP+1:])
								info.ExitIP = opsJF[0]
								info.IsIn = true
								opsGet, _ := parser.ReadOperands(parser.OpcodeOperands[op2], insts[target+1:])
								info.ItVar = opsGet[0]
							}
						}
					}
				}

				if !info.IsIn {
					k := target
					for k < instIP {
						opK := insts[k]
						numOpsK := parser.OpcodeOperands[opK]
						opsK, readK := parser.ReadOperands(numOpsK, insts[k+1:])
						if opK == parser.OpJumpFalsy {
							info.ExitIP = opsK[0]
							break
						}
						k += 1 + readK
					}
					if info.ExitIP == 0 {
						info.ExitIP = instIP + 3
					}
				}

				postIP := instIP
				k := target
				for k < instIP {
					opK := insts[k]
					numOpsK := parser.OpcodeOperands[opK]
					opsK, readK := parser.ReadOperands(numOpsK, insts[k+1:])
					if opK == parser.OpJump {
						tgt := opsK[0]
						if tgt > target && tgt < instIP {
							if tgt < postIP {
								postIP = tgt
							}
						}
					}
					k += 1 + readK
				}
				info.PostIP = postIP

				loops[target] = info
			}
		}
	}
	return loops
}

func isPostStatement(stmt string) bool {
	trimmed := strings.TrimSpace(stmt)
	if strings.HasSuffix(trimmed, "++") || strings.HasSuffix(trimmed, "--") {
		return true
	}
	if strings.Contains(trimmed, " = ") || strings.Contains(trimmed, " := ") ||
		strings.Contains(trimmed, " += ") || strings.Contains(trimmed, " -= ") ||
		strings.Contains(trimmed, " *= ") || strings.Contains(trimmed, " /= ") {
		return true
	}
	return false
}

func (fd *FunctionDecompiler) Decompile() []string {
	fd.loops = fd.analyzeLoops()
	code, _ := fd.decompileBlock(0, len(fd.fn.Instructions), []string{}, nil)
	return code
}

func (fd *FunctionDecompiler) decompileBlock(startIP, endIP int, stack []string, activeLoops []*LoopInfo) ([]string, []string) {
	var code []string
	insts := fd.fn.Instructions
	i := startIP
	seenReturn := false

	for i < endIP && !seenReturn {
		instIP := i

		var activeLoop *LoopInfo
		var isForIn bool

		if loop, ok := fd.loops[i]; ok {
			activeLoop = loop
		} else {
			for _, loop := range fd.loops {
				if loop.IsIn && i < loop.HeaderIP && loop.HeaderIP-i <= 15 {
					hasIterInit := false
					for k := i; k < loop.HeaderIP; k++ {
						if insts[k] == parser.OpIteratorInit {
							hasIterInit = true
							break
						}
					}
					if hasIterInit {
						activeLoop = loop
						isForIn = true
						break
					}
				}
			}
		}

		if activeLoop != nil {
			delete(fd.loops, activeLoop.HeaderIP)

			if isForIn {
				iterInitIP := i
				for k := i; k < activeLoop.HeaderIP; k++ {
					if insts[k] == parser.OpIteratorInit {
						iterInitIP = k
						break
					}
				}

				iterableStack := append([]string{}, stack...)
				iterCode, finalIterStack := fd.decompileBlock(i, iterInitIP, iterableStack, activeLoops)
				code = append(code, iterCode...)

				iterable := "nil"
				if len(finalIterStack) > 0 {
					iterable = finalIterStack[len(finalIterStack)-1]
				}

				bodyStartIP := activeLoop.HeaderIP
				opH := insts[activeLoop.HeaderIP]
				skip := 1
				if opH == parser.OpGetLocal {
					skip = 2
				} else if opH == parser.OpGetGlobal {
					skip = 3
				}
				bodyStartIP = activeLoop.HeaderIP + skip + 1 + 3

				keyVar := "_"
				valVar := "_"
				if bodyStartIP < len(insts) {
					k := bodyStartIP
					op1 := insts[k]
					if op1 == parser.OpGetLocal || op1 == parser.OpGetGlobal {
						opLen1 := 2
						if op1 == parser.OpGetGlobal {
							opLen1 = 3
						}
						nextIP := k + opLen1
						if nextIP < len(insts) && insts[nextIP] == parser.OpIteratorKey {
							keyAssignIP := nextIP + 1
							if keyAssignIP < len(insts) {
								op2 := insts[keyAssignIP]
								if op2 == parser.OpDefineLocal || op2 == parser.OpSetGlobal {
									opLen2 := 2
									if op2 == parser.OpSetGlobal {
										opLen2 = 3
									}
									ops2, _ := parser.ReadOperands([]int{opLen2 - 1}, insts[keyAssignIP+1:])
									if op2 == parser.OpDefineLocal {
										keyVar = fd.getLocal(ops2[0])
									} else {
										keyVar = getGlobalName(ops2[0])
									}
									bodyStartIP = keyAssignIP + opLen2
								}
							}
						}
					}
				}

				if bodyStartIP < len(insts) {
					k := bodyStartIP
					op1 := insts[k]
					if op1 == parser.OpGetLocal || op1 == parser.OpGetGlobal {
						opLen1 := 2
						if op1 == parser.OpGetGlobal {
							opLen1 = 3
						}
						nextIP := k + opLen1
						if nextIP < len(insts) && insts[nextIP] == parser.OpIteratorValue {
							valAssignIP := nextIP + 1
							if valAssignIP < len(insts) {
								op2 := insts[valAssignIP]
								if op2 == parser.OpDefineLocal || op2 == parser.OpSetGlobal {
									opLen2 := 2
									if op2 == parser.OpSetGlobal {
										opLen2 = 3
									}
									ops2, _ := parser.ReadOperands([]int{opLen2 - 1}, insts[valAssignIP+1:])
									if op2 == parser.OpDefineLocal {
										valVar = fd.getLocal(ops2[0])
									} else {
										valVar = getGlobalName(ops2[0])
									}
									bodyStartIP = valAssignIP + opLen2
								}
							}
						}
					}
				}

				bodyCode, _ := fd.decompileBlock(bodyStartIP, activeLoop.EndIP, []string{}, append(activeLoops, activeLoop))

				header := ""
				if valVar == "_" {
					header = fmt.Sprintf("for %s in %s {", keyVar, iterable)
				} else {
					header = fmt.Sprintf("for %s, %s in %s {", keyVar, valVar, iterable)
				}
				code = append(code, header)
				code = append(code, bodyCode...)
				code = append(code, "}")
			} else {
				initExpr := ""
				if len(code) > 0 {
					lastLine := code[len(code)-1]
					if isPostStatement(lastLine) {
						initExpr = lastLine
						code = code[:len(code)-1]
					}
				}

				condJumpIP := activeLoop.ExitIP - 3
				k := activeLoop.HeaderIP
				for k < activeLoop.EndIP {
					opK := insts[k]
					numOpsK := parser.OpcodeOperands[opK]
					opsK, readK := parser.ReadOperands(numOpsK, insts[k+1:])
					if opK == parser.OpJumpFalsy && opsK[0] == activeLoop.ExitIP {
						condJumpIP = k
						break
					}
					k += 1 + readK
				}

				condExpr := ""
				var bodyStartIP int
				if condJumpIP < activeLoop.EndIP && insts[condJumpIP] == parser.OpJumpFalsy {
					_, finalCondStack := fd.decompileBlock(activeLoop.HeaderIP, condJumpIP, []string{}, activeLoops)
					if len(finalCondStack) > 0 {
						condExpr = finalCondStack[len(finalCondStack)-1]
					}
					bodyStartIP = condJumpIP + 3
				} else {
					bodyStartIP = activeLoop.HeaderIP
				}

				bodyCode, _ := fd.decompileBlock(bodyStartIP, activeLoop.PostIP, []string{}, append(activeLoops, activeLoop))

				postExpr := ""
				if activeLoop.PostIP < activeLoop.EndIP {
					postCode, _ := fd.decompileBlock(activeLoop.PostIP, activeLoop.EndIP, []string{}, activeLoops)
					if len(postCode) > 0 {
						postExpr = postCode[len(postCode)-1]
					}
				}

				if postExpr == "" && len(bodyCode) > 0 {
					lastLine := bodyCode[len(bodyCode)-1]
					if isPostStatement(lastLine) {
						postExpr = lastLine
						bodyCode = bodyCode[:len(bodyCode)-1]
					}
				}

				header := "for"
				if initExpr != "" || postExpr != "" {
					header += " " + initExpr + "; " + condExpr + "; " + postExpr
				} else if condExpr != "" && condExpr != "true" {
					header += " " + condExpr
				}
				header += " {"

				code = append(code, header)
				code = append(code, bodyCode...)
				code = append(code, "}")
			}

			fd.loops[activeLoop.HeaderIP] = activeLoop
			i = activeLoop.EndIP + 3
			continue
		}

		op := insts[i]
		numOperands := parser.OpcodeOperands[op]
		operands, read := parser.ReadOperands(numOperands, insts[i+1:])
		i += 1 + read

		if op == parser.OpJumpFalsy {
			condExpr := "true"
			if len(stack) > 0 {
				condExpr = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}

			// If condition is still "true", try to derive it from context
			if condExpr == "true" && instIP > 0 {
				// Walk backwards to find the condition expression
				for j := instIP - 1; j >= 0 && j > instIP-20; j-- {
					if insts[j] == parser.OpBinaryOp || insts[j] == parser.OpEqual || insts[j] == parser.OpNotEqual {
						// The condition is on the stack - try to reconstruct it
						if len(stack) > 0 {
							condExpr = stack[len(stack)-1]
							stack = stack[:len(stack)-1]
							break
						}
					}
				}
			}

			elseStartIP := operands[0]
			var ifEndIP int
			var hasElse bool

			if elseStartIP-3 >= instIP && insts[elseStartIP-3] == parser.OpJump {
				opsJump, _ := parser.ReadOperands(parser.OpcodeOperands[parser.OpJump], insts[elseStartIP-2:])
				if opsJump[0] > elseStartIP {
					ifEndIP = opsJump[0]
					hasElse = true
				} else {
					ifEndIP = elseStartIP
				}
			} else {
				ifEndIP = elseStartIP
			}

			var ifBody []string
			if hasElse {
				ifBody, _ = fd.decompileBlock(instIP+3, elseStartIP-3, []string{}, activeLoops)
			} else {
				ifBody, _ = fd.decompileBlock(instIP+3, elseStartIP, []string{}, activeLoops)
			}

			var elseBody []string
			if hasElse {
				elseBody, _ = fd.decompileBlock(elseStartIP, ifEndIP, []string{}, activeLoops)
			}

			// Skip empty if statements
			if len(ifBody) > 0 || (hasElse && len(elseBody) > 0) {
				code = append(code, fmt.Sprintf("if %s {", condExpr))
				code = append(code, ifBody...)

				if hasElse {
					if len(elseBody) > 0 && strings.HasPrefix(strings.TrimSpace(elseBody[0]), "if ") {
						trimmedIf := strings.TrimSpace(elseBody[0])
						code = append(code, "} else "+trimmedIf)
						code = append(code, elseBody[1:]...)
					} else if len(elseBody) > 0 {
						code = append(code, "} else {")
						code = append(code, elseBody...)
						code = append(code, "}")
					} else {
						code = append(code, "} else {}")
					}
				} else {
					code = append(code, "}")
				}
			}

			i = ifEndIP
			continue
		}

		if op == parser.OpAndJump || op == parser.OpOrJump || op == parser.OpNotNullJump || op == parser.OpNullJump {
			targetIP := operands[0]
			lhs := "nil"
			if len(stack) > 0 {
				lhs = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}

			if op == parser.OpAndJump {
				_, finalRhsStack := fd.decompileBlock(instIP+3, targetIP, []string{}, activeLoops)
				rhs := "nil"
				if len(finalRhsStack) > 0 {
					rhs = finalRhsStack[len(finalRhsStack)-1]
				}
				stack = append(stack, fmt.Sprintf("(%s && %s)", lhs, rhs))
			} else if op == parser.OpOrJump {
				_, finalRhsStack := fd.decompileBlock(instIP+3, targetIP, []string{}, activeLoops)
				rhs := "nil"
				if len(finalRhsStack) > 0 {
					rhs = finalRhsStack[len(finalRhsStack)-1]
				}
				stack = append(stack, fmt.Sprintf("(%s || %s)", lhs, rhs))
			} else if op == parser.OpNotNullJump {
				_, finalRhsStack := fd.decompileBlock(instIP+3, targetIP, []string{}, activeLoops)
				rhs := "nil"
				if len(finalRhsStack) > 0 {
					rhs = finalRhsStack[len(finalRhsStack)-1]
				}
				stack = append(stack, fmt.Sprintf("(%s ?? %s)", lhs, rhs))
			} else if op == parser.OpNullJump {
				if targetIP > 0 {
					lastOp := insts[targetIP-1]
					if lastOp == parser.OpIndex {
						_, finalRhsStack := fd.decompileBlock(instIP+3, targetIP-1, []string{}, activeLoops)
						idx := "nil"
						if len(finalRhsStack) > 0 {
							idx = finalRhsStack[len(finalRhsStack)-1]
						}
						if strings.HasPrefix(idx, "\"") && strings.HasSuffix(idx, "\"") {
							selName := idx[1 : len(idx)-1]
							if isValidIdent(selName) {
								stack = append(stack, fmt.Sprintf("%s?.%s", lhs, selName))
								i = targetIP
								continue
							}
						}
						stack = append(stack, fmt.Sprintf("%s?.[%s]", lhs, idx))
					} else if targetIP >= 3 && insts[targetIP-3] == parser.OpCall {
						opsCall, _ := parser.ReadOperands(parser.OpcodeOperands[parser.OpCall], insts[targetIP-2:])
						numArgs := opsCall[0]
						isEllipsis := opsCall[1] == 1

						_, finalRhsStack := fd.decompileBlock(instIP+3, targetIP-3, []string{}, activeLoops)
						var args []string
						for k := 0; k < numArgs; k++ {
							if len(finalRhsStack) > 0 {
								args = append([]string{finalRhsStack[len(finalRhsStack)-1]}, args...)
								finalRhsStack = finalRhsStack[:len(finalRhsStack)-1]
							}
						}
						ellipStr := ""
						if isEllipsis {
							ellipStr = "..."
						}
						stack = append(stack, fmt.Sprintf("%s?.(%s%s)", lhs, strings.Join(args, ", "), ellipStr))
					}
				}
			}
			i = targetIP
			continue
		}

		switch op {
		case parser.OpConstant:
			constIdx := operands[0]
			cObj := fd.bytecode.Constants[constIdx]
			if compiledFn, ok := cObj.(*tender.CompiledFunction); ok {
				subFD := NewFunctionDecompiler(compiledFn, fd.bytecode, nil)
				lines := subFD.Decompile()
				body := "fn() {\n"
				for _, line := range lines {
					body += "    " + line + "\n"
				}
				body += "}"
				stack = append(stack, body)
			} else {
				stack = append(stack, fd.constants[constIdx])
			}

		case parser.OpPop:
			if len(stack) > 0 {
				val := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if strings.Contains(val, "(") || strings.Contains(val, "=") || strings.Contains(val, ":=") {
					code = append(code, val)
				}
			}

		case parser.OpTrue:
			stack = append(stack, "true")
		case parser.OpFalse:
			stack = append(stack, "false")
		case parser.OpNull:
			stack = append(stack, "null")

		case parser.OpGetGlobal:
			stack = append(stack, getGlobalName(operands[0]))

		case parser.OpSetGlobal:
			if len(stack) > 0 {
				val := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				gName := getGlobalName(operands[0])

				if strings.HasPrefix(val, "import_module:") {
					modName := strings.TrimPrefix(val, "import_module:")
					if gName == modName || gName == filepath.Base(modName) || strings.HasPrefix(gName, "__mod_") {
						code = append(code, fmt.Sprintf("import %q", modName))
						fd.definedGlobals[operands[0]] = true
						continue
					} else {
						val = fmt.Sprintf("import(%q)", modName)
					}
				}

				// Check if the value was flagged by OpImmutable
				isConst := false
				if strings.HasPrefix(val, "immutable(") && strings.HasSuffix(val, ")") {
					val = val[10 : len(val)-1]
					isConst = true
				}

				opStr := "="
				prefix := ""
				if !fd.definedGlobals[operands[0]] {
					if isConst {
						prefix = "const "
						opStr = "="
					} else {
						opStr = ":="
					}
					fd.definedGlobals[operands[0]] = true
				} else if isConst {
					// If already defined, it's an immutable expression, not a declaration
					val = fmt.Sprintf("immutable(%s)", val)
				}

				code = append(code, fmt.Sprintf("%s%s %s %s", prefix, gName, opStr, val))
			}

		case parser.OpSetSelGlobal:
			numSel := operands[1]
			var sels []string
			for k := 0; k < numSel; k++ {
				if len(stack) > 0 {
					sels = append([]string{stack[len(stack)-1]}, sels...)
					stack = stack[:len(stack)-1]
				}
			}
			if len(stack) > 0 {
				val := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				target := getGlobalName(operands[0])
				for _, sel := range sels {
					if strings.HasPrefix(sel, "\"") && strings.HasSuffix(sel, "\"") {
						selName := sel[1 : len(sel)-1]
						if isValidIdent(selName) {
							target += fmt.Sprintf(".%s", selName)
							continue
						}
					}
					target += fmt.Sprintf("[%s]", sel)
				}
				code = append(code, fmt.Sprintf("%s = %s", target, val))
			}

		case parser.OpGetLocal:
			stack = append(stack, fd.getLocal(operands[0]))

		case parser.OpDefineLocal:
			if len(stack) > 0 {
				val := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if strings.HasPrefix(val, "import_module:") {
					val = fmt.Sprintf("import(%q)", strings.TrimPrefix(val, "import_module:"))
				}

				// Check if the value was flagged by OpImmutable
				isConst := false
				if strings.HasPrefix(val, "immutable(") && strings.HasSuffix(val, ")") {
					val = val[10 : len(val)-1]
					isConst = true
				}

				if isConst {
					code = append(code, fmt.Sprintf("const %s = %s", fd.getLocal(operands[0]), val))
				} else {
					code = append(code, fmt.Sprintf("%s := %s", fd.getLocal(operands[0]), val))
				}
			}

		case parser.OpSetLocal:
			if len(stack) > 0 {
				val := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if strings.HasPrefix(val, "import_module:") {
					val = fmt.Sprintf("import(%q)", strings.TrimPrefix(val, "import_module:"))
				}
				code = append(code, fmt.Sprintf("%s = %s", fd.getLocal(operands[0]), val))
			}

		case parser.OpSetSelLocal:
			numSel := operands[1]
			var sels []string
			for k := 0; k < numSel; k++ {
				if len(stack) > 0 {
					sels = append([]string{stack[len(stack)-1]}, sels...)
					stack = stack[:len(stack)-1]
				}
			}
			if len(stack) > 0 {
				val := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				target := fd.getLocal(operands[0])
				for _, sel := range sels {
					if strings.HasPrefix(sel, "\"") && strings.HasSuffix(sel, "\"") {
						selName := sel[1 : len(sel)-1]
						if isValidIdent(selName) {
							target += fmt.Sprintf(".%s", selName)
							continue
						}
					}
					target += fmt.Sprintf("[%s]", sel)
				}
				code = append(code, fmt.Sprintf("%s = %s", target, val))
			}

		case parser.OpGetFree:
			stack = append(stack, fd.getFree(operands[0]))

		case parser.OpSetFree:
			if len(stack) > 0 {
				val := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				code = append(code, fmt.Sprintf("%s = %s", fd.getFree(operands[0]), val))
			}

		case parser.OpSetSelFree:
			numSel := operands[1]
			var sels []string
			for k := 0; k < numSel; k++ {
				if len(stack) > 0 {
					sels = append([]string{stack[len(stack)-1]}, sels...)
					stack = stack[:len(stack)-1]
				}
			}
			if len(stack) > 0 {
				val := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				target := fd.getFree(operands[0])
				for _, sel := range sels {
					if strings.HasPrefix(sel, "\"") && strings.HasSuffix(sel, "\"") {
						selName := sel[1 : len(sel)-1]
						if isValidIdent(selName) {
							target += fmt.Sprintf(".%s", selName)
							continue
						}
					}
					target += fmt.Sprintf("[%s]", sel)
				}
				code = append(code, fmt.Sprintf("%s = %s", target, val))
			}

		case parser.OpGetFreePtr, parser.OpGetLocalPtr:
			var target string
			if op == parser.OpGetFreePtr {
				target = fd.getFree(operands[0])
			} else {
				target = fd.getLocal(operands[0])
			}
			stack = append(stack, "&"+target)

		case parser.OpGetBuiltin:
			stack = append(stack, getBuiltinName(operands[0]))

		case parser.OpBinaryOp:
			if len(stack) >= 2 {
				right := stack[len(stack)-1]
				left := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, fmt.Sprintf("(%s %s %s)", left, getBinaryOpStr(operands[0]), right))
			}

		case parser.OpEqual, parser.OpNotEqual:
			if len(stack) >= 2 {
				right := stack[len(stack)-1]
				left := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				opStr := "=="
				if op == parser.OpNotEqual {
					opStr = "!="
				}
				stack = append(stack, fmt.Sprintf("(%s %s %s)", left, opStr, right))
			}

		case parser.OpLNot:
			if len(stack) > 0 {
				val := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, "!"+val)
			}
		case parser.OpMinus:
			if len(stack) > 0 {
				val := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, "-"+val)
			}
		case parser.OpBComplement:
			if len(stack) > 0 {
				val := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, "^"+val)
			}

		case parser.OpArray:
			num := operands[0]
			var items []string
			for k := 0; k < num; k++ {
				if len(stack) > 0 {
					items = append([]string{stack[len(stack)-1]}, items...)
					stack = stack[:len(stack)-1]
				}
			}
			stack = append(stack, "["+strings.Join(items, ", ")+"]")

		case parser.OpTuple:
			num := operands[0]
			var items []string
			for k := 0; k < num; k++ {
				if len(stack) > 0 {
					items = append([]string{stack[len(stack)-1]}, items...)
					stack = stack[:len(stack)-1]
				}
			}
			if len(items) == 1 {
				stack = append(stack, "("+items[0]+",)")
			} else {
				stack = append(stack, "("+strings.Join(items, ", ")+")")
			}

		case parser.OpMap:
			num := operands[0]
			var pairs []string
			for k := 0; k < num/2; k++ {
				if len(stack) >= 2 {
					val := stack[len(stack)-1]
					key := stack[len(stack)-2]
					stack = stack[:len(stack)-2]
					pairs = append([]string{fmt.Sprintf("%s: %s", key, val)}, pairs...)
				}
			}
			stack = append(stack, "{"+strings.Join(pairs, ", ")+"}")

		case parser.OpIndex:
			if len(stack) >= 2 {
				idx := stack[len(stack)-1]
				obj := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				if strings.HasPrefix(idx, "\"") && strings.HasSuffix(idx, "\"") {
					selName := idx[1 : len(idx)-1]
					// Use dot notation only if it's a valid identifier AND not a keyword
					if isValidIdent(selName) && !isKeyword(selName) {
						stack = append(stack, fmt.Sprintf("%s.%s", obj, selName))
						continue
					}
					// Otherwise use bracket notation (for keywords or invalid identifiers)
					stack = append(stack, fmt.Sprintf("%s[%s]", obj, idx))
					continue
				}
				stack = append(stack, fmt.Sprintf("%s[%s]", obj, idx))
			}

		case parser.OpSliceIndex:
			if len(stack) >= 3 {
				high := stack[len(stack)-1]
				low := stack[len(stack)-2]
				obj := stack[len(stack)-3]
				stack = stack[:len(stack)-3]
				if low == "null" {
					low = ""
				}
				if high == "null" {
					high = ""
				}
				stack = append(stack, fmt.Sprintf("%s[%s:%s]", obj, low, high))
			}

		case parser.OpCall:
			numArgs := operands[0]
			isEllipsis := operands[1] == 1
			var args []string
			for k := 0; k < numArgs; k++ {
				if len(stack) > 0 {
					args = append([]string{stack[len(stack)-1]}, args...)
					stack = stack[:len(stack)-1]
				}
			}
			if len(stack) > 0 {
				fnStr := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				ellipStr := ""
				if isEllipsis {
					ellipStr = "..."
				}
				callExpr := fmt.Sprintf("%s(%s%s)", fnStr, strings.Join(args, ", "), ellipStr)
				stack = append(stack, callExpr)
				// fd.lastCallExpr = callExpr
			}

		case parser.OpClosure:
			constIdx := operands[0]
			numFree := operands[1]
			var captured []string
			for k := 0; k < numFree; k++ {
				if len(stack) > 0 {
					captured = append([]string{stack[len(stack)-1]}, captured...)
					stack = stack[:len(stack)-1]
				}
			}
			compiledFn := fd.bytecode.Constants[constIdx].(*tender.CompiledFunction)

			var subFreeNames []string
			for _, capName := range captured {
				name := strings.TrimPrefix(capName, "&")
				subFreeNames = append(subFreeNames, name)
			}

			subFD := NewFunctionDecompiler(compiledFn, fd.bytecode, subFreeNames)
			lines := subFD.Decompile()

			var paramNames []string
			for k := 0; k < compiledFn.NumParameters; k++ {
				paramNames = append(paramNames, subFD.getLocal(k))
			}
			if compiledFn.VarArgs {
				if len(paramNames) > 0 {
					paramNames[len(paramNames)-1] += "..."
				} else {
					paramNames = append(paramNames, "...")
				}
			}

			body := fmt.Sprintf("fn(%s) {\n", strings.Join(paramNames, ", "))
			for _, line := range lines {
				body += "    " + line + "\n"
			}
			body += "}"
			stack = append(stack, body)

		case parser.OpReturn:
			numResults := operands[0]
			var results []string
			for k := 0; k < numResults; k++ {
				if len(stack) > 0 {
					results = append([]string{stack[len(stack)-1]}, results...)
					stack = stack[:len(stack)-1]
				}
			}
			code = append(code, "return "+strings.Join(results, ", "))
			seenReturn = true
			i = endIP
			continue

		case parser.OpJump:
			targetPos := operands[0]
			resolvedBranch := false
			for idx := len(activeLoops) - 1; idx >= 0; idx-- {
				loop := activeLoops[idx]
				if targetPos == loop.ExitIP {
					code = append(code, "break")
					resolvedBranch = true
					break
				}
				if targetPos == loop.HeaderIP || targetPos == loop.PostIP || targetPos == loop.EndIP {
					code = append(code, "continue")
					resolvedBranch = true
					break
				}
			}
			if !resolvedBranch {
				code = append(code, fmt.Sprintf("// goto label_%d", targetPos))
			}

		case parser.OpIteratorInit, parser.OpIteratorNext, parser.OpIteratorKey, parser.OpIteratorValue:
			// handled by loop detection

		case parser.OpImmutable:
			if len(stack) > 0 {
				val := stack[len(stack)-1]
				stack[len(stack)-1] = "immutable(" + val + ")"
			}

		case parser.OpError:
			if len(stack) > 0 {
				val := stack[len(stack)-1]
				stack[len(stack)-1] = "error(" + val + ")"
			}

		case parser.OpStruct:
			numFields := operands[0]
			isKeyed := operands[1] == 1
			var fields []string
			if isKeyed {
				for k := 0; k < numFields; k++ {
					if len(stack) >= 2 {
						val := stack[len(stack)-1]
						key := stack[len(stack)-2]
						stack = stack[:len(stack)-2]
						fields = append([]string{fmt.Sprintf("%s: %s", key, val)}, fields...)
					}
				}
			} else {
				for k := 0; k < numFields; k++ {
					if len(stack) > 0 {
						fields = append([]string{stack[len(stack)-1]}, fields...)
						stack = stack[:len(stack)-1]
					}
				}
			}
			structType := "Struct"
			if len(stack) > 0 {
				structType = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, fmt.Sprintf("%s{%s}", structType, strings.Join(fields, ", ")))

		case parser.OpMethod:
			if len(stack) >= 3 {
				methodName := stack[len(stack)-1]
				receiverType := stack[len(stack)-2]
				methodFunc := stack[len(stack)-3]
				stack = stack[:len(stack)-3]

				if strings.HasPrefix(methodName, "\"") && strings.HasSuffix(methodName, "\"") {
					methodName = methodName[1 : len(methodName)-1]
				}

				if strings.HasPrefix(methodFunc, "fn(") {
					closeParenIdx := strings.Index(methodFunc, ")")
					if closeParenIdx != -1 {
						params := methodFunc[3:closeParenIdx]
						body := methodFunc[closeParenIdx+1:]

						receiverName := "self"
						restParams := params
						if strings.Contains(params, ",") {
							commaIdx := strings.Index(params, ",")
							receiverName = strings.TrimSpace(params[:commaIdx])
							restParams = strings.TrimSpace(params[commaIdx+1:])
						} else if len(strings.TrimSpace(params)) > 0 {
							receiverName = strings.TrimSpace(params)
							restParams = ""
						}

						code = append(code, fmt.Sprintf("fn (%s %s) %s(%s)%s", receiverName, receiverType, methodName, restParams, body))
						continue
					}
				}
				code = append(code, fmt.Sprintf("method %s.%s = %s", receiverType, methodName, methodFunc))
			}

		case parser.OpSuspend:
			// no-op
		}
	}

	// if !seenReturn {
	// for _, val := range stack {
	// code = append(code, val)
	// }
	// }

	return code, stack
}

// formatSourceCode pretty prints the decompiled source code with proper indentation
func formatSourceCode(lines []string) string {
	if len(lines) == 0 {
		return ""
	}

	var result []string
	indent := 0
	indentSize := 4

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			result = append(result, "")
			continue
		}

		// Check if this line ends with a block that decreases indentation
		dec := 0
		for _, char := range trimmed {
			if char == '}' || char == ')' {
				dec++
			} else {
				break
			}
		}

		// Apply indentation decrease before printing
		indent -= dec
		if indent < 0 {
			indent = 0
		}

		// Format with current indentation
		spaces := strings.Repeat(" ", indent*indentSize)
		formatted := spaces + trimmed
		result = append(result, formatted)

		// Check if this line starts a block that needs indentation increase
		inc := 0
		for i := len(trimmed) - 1; i >= 0; i-- {
			char := trimmed[i]
			if char == '{' {
				// Check if it's not part of a string or comment
				if i == 0 || trimmed[i-1] != '\'' {
					inc++
				}
			} else if char == '(' && i > 0 && trimmed[i-1] == 'f' && i > 1 && trimmed[i-2] == 'n' {
				// fn( doesn't increase indentation
				continue
			} else if !unicode.IsSpace(rune(char)) {
				break
			}
		}
		indent += inc
	}

	// Remove consecutive empty lines
	var cleaned []string
	for i, line := range result {
		if i > 0 && line == "" && cleaned[len(cleaned)-1] == "" {
			continue
		}
		cleaned = append(cleaned, line)
	}

	return strings.Join(cleaned, "\n")
}

func isValidIdent(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i, c := range s {
		if i == 0 {
			if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_') {
				return false
			}
		} else {
			if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '_') {
				return false
			}
		}
	}
	return true
}

func isPrintable(data []byte) bool {
	if len(data) == 0 {
		return false
	}
	for _, b := range data {
		if b > 127 {
			return false
		}
		if !unicode.IsPrint(rune(b)) && !unicode.IsSpace(rune(b)) {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: undo.exe <path-to-tdo-file> [output-directory]")
		os.Exit(1)
	}

	filePath := os.Args[1]
	outDir := "."
	if len(os.Args) >= 3 {
		outDir = os.Args[2]
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		os.Exit(1)
	}

	modules := stdlib.GetModuleMap(stdlib.AllModuleNames()...)
	bytecode := &tender.Bytecode{}
	err = bytecode.Decode(bytes.NewReader(data), modules)
	if err != nil {
		fmt.Printf("Failed to decode bytecode: %v\n", err)
		os.Exit(1)
	}

	preScanGlobals(bytecode)

	err = os.MkdirAll(outDir, 0755)
	if err != nil {
		fmt.Printf("Failed to create output directory: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Decompiling bytecode to source files in: %s\n", outDir)

	fileOutputs := make(map[string][]string)

	for idx, c := range bytecode.Constants {
		if c == nil {
			continue
		}
		if fn, ok := c.(*tender.CompiledFunction); ok {
			fileName := "unknown.td"
			for _, pos := range fn.SourceMap {
				if f := bytecode.FileSet.File(pos); f != nil {
					fileName = f.Name
					break
				}
			}
			fileName = filepath.Base(fileName)

			fd := NewFunctionDecompiler(fn, bytecode, nil)
			lines := fd.Decompile()

			var text []string
			text = append(text, fmt.Sprintf("// --- Constant Function [%d] ---", idx))

			var paramNames []string
			for k := 0; k < fn.NumParameters; k++ {
				paramNames = append(paramNames, fd.getLocal(k))
			}
			if fn.VarArgs {
				if len(paramNames) > 0 {
					paramNames[len(paramNames)-1] += "..."
				} else {
					paramNames = append(paramNames, "...")
				}
			}
			text = append(text, fmt.Sprintf("fn_const_%d := fn(%s) {", idx, strings.Join(paramNames, ", ")))
			for _, line := range lines {
				text = append(text, "    "+line)
			}
			text = append(text, "}")
			text = append(text, "")

			fileOutputs[fileName] = append(fileOutputs[fileName], text...)
		}
	}

	if bytecode.MainFunction != nil {
		fileName := "slots.td"
		for _, pos := range bytecode.MainFunction.SourceMap {
			if f := bytecode.FileSet.File(pos); f != nil {
				fileName = f.Name
				break
			}
		}
		fileName = filepath.Base(fileName)

		fd := NewFunctionDecompiler(bytecode.MainFunction, bytecode, nil)
		lines := fd.Decompile()

		var text []string
		text = append(text, "// --- Main Function ---")
		for _, line := range lines {
			text = append(text, line)
		}
		fileOutputs[fileName] = append(fileOutputs[fileName], text...)
	}

	for name, lines := range fileOutputs {
		outPath := filepath.Join(outDir, name)
		content := formatSourceCode(lines)
		err := os.WriteFile(outPath, []byte(content), 0644)
		if err != nil {
			fmt.Printf("[-] Failed to write to %s: %v\n", outPath, err)
		} else {
			fmt.Printf("[+] Successfully decompiled: %s\n", outPath)
		}
	}
}
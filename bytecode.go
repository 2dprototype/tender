package tender

import (
	"bufio"
	"encoding/gob"
	"math/big"
	"io"
	"fmt"
	"reflect"

	"github.com/2dprototype/tender/parser"
)

// Bytecode is a compiled instructions and constants.
type Bytecode struct {
	FileSet      *parser.SourceFileSet
	MainFunction *Function
	Constants    []Object
}


// func (b *Bytecode) EncodeJSON(w io.Writer) error {
    // jsonData, err := json.Marshal(b)
    // if err != nil {
        // return err
    // }
    // _, err = w.Write(jsonData)
    // return err
// }

// // DecodeJSON reads Bytecode data from the reader in JSON format.
// func (b *Bytecode) DecodeJSON(r io.Reader) error {
    // decoder := json.NewDecoder(r)
    // return decoder.Decode(b)
// }

// Encode writes Bytecode data to the writer.
func (b *Bytecode) Encode(w io.Writer) error {
	enc := gob.NewEncoder(w)
	if err := enc.Encode(b.FileSet); err != nil {
		return err
	}
	if err := enc.Encode(b.MainFunction); err != nil {
		return err
	}
	return enc.Encode(b.Constants)
}

// CountObjects returns the number of objects found in Constants.
func (b *Bytecode) CountObjects() int {
	n := 0
	for _, c := range b.Constants {
		n += CountObjects(c)
	}
	return n
}

// FormatInstructions returns human readable string representations of
// compiled instructions.
func (b *Bytecode) FormatInstructions() []string {
	return FormatInstructions(b.MainFunction.Instructions, 0)
}

// FormatConstants returns human readable string representations of
// compiled constants.
func (b *Bytecode) FormatConstants() (output []string) {
	for cidx, cn := range b.Constants {
		switch cn := cn.(type) {
		case *Function:
			output = append(output, fmt.Sprintf(
				"[% 3d] (Compiled Function|%p)", cidx, &cn))
			for _, l := range FormatInstructions(cn.Instructions, 0) {
				output = append(output, fmt.Sprintf("     %s", l))
			}
		default:
			output = append(output, fmt.Sprintf("[% 3d] %s (%s|%p)",
				cidx, cn, reflect.TypeOf(cn).Elem().Name(), &cn))
		}
	}
	return
}

// Decode reads Bytecode data from the reader.
func (b *Bytecode) Decode(r io.Reader, modules *ModuleMap) error {
	if modules == nil {
		modules = NewModuleMap()
	}

	br := bufio.NewReader(r)
	magic, err := br.Peek(4)
	if err == nil && len(magic) == 4 && magic[0] == 'T' && magic[1] == 'D' && magic[2] == 'C' && magic[3] == 1 {
		_, _, err = b.DecodeTDC(br, modules)
		return err
	}

	dec := gob.NewDecoder(br)
	if err := dec.Decode(&b.FileSet); err != nil {
		return err
	}
	// TODO: files in b.FileSet.File does not have their 'set' field properly
	//  set to b.FileSet as it's private field and not serialized by gob
	//  encoder/decoder.
	if err := dec.Decode(&b.MainFunction); err != nil {
		return err
	}
	if err := dec.Decode(&b.Constants); err != nil {
		return err
	}
	for i, v := range b.Constants {
		fv, err := fixDecodedObject(v, modules)
		if err != nil {
			return err
		}
		b.Constants[i] = fv
	}
	return nil
}

// RemoveDuplicates finds and remove the duplicate values in Constants.
// Note this function mutates Bytecode.
func (b *Bytecode) RemoveDuplicates() {
	var deduped []Object

	indexMap := make(map[int]int) // mapping from old constant index to new index
	fns := make(map[*Function]int)
	ints := make(map[int64]int)
	bigints := make(map[*big.Int]int)
	strings := make(map[string]int)
	floats := make(map[float64]int)
	bigfloats := make(map[*big.Float]int)
	complexs := make(map[complex128]int)
	chars := make(map[rune]int)
	immutableMaps := make(map[string]int) // for modules

	for curIdx, c := range b.Constants {
		switch c := c.(type) {
		case *StructType:
			foundIdx := -1
			for i, d := range deduped {
				if dt, ok := d.(*StructType); ok && dt.Equals(c) {
					foundIdx = i
					break
				}
			}
			if foundIdx != -1 {
				indexMap[curIdx] = foundIdx
			} else {
				newIdx := len(deduped)
				indexMap[curIdx] = newIdx
				deduped = append(deduped, c)
			}
		case *Function:
			if newIdx, ok := fns[c]; ok {
				indexMap[curIdx] = newIdx
			} else {
				newIdx = len(deduped)
				fns[c] = newIdx
				indexMap[curIdx] = newIdx
				deduped = append(deduped, c)
			}
		case *ImmutableMap:
			modName := inferModuleName(c)
			newIdx, ok := immutableMaps[modName]
			if modName != "" && ok {
				indexMap[curIdx] = newIdx
			} else {
				newIdx = len(deduped)
				immutableMaps[modName] = newIdx
				indexMap[curIdx] = newIdx
				deduped = append(deduped, c)
			}
		case *Int:
			if newIdx, ok := ints[c.Value]; ok {
				indexMap[curIdx] = newIdx
			} else {
				newIdx = len(deduped)
				ints[c.Value] = newIdx
				indexMap[curIdx] = newIdx
				deduped = append(deduped, c)
			}	
		case *BigInt:
			if newIdx, ok := bigints[c.Value]; ok {
				indexMap[curIdx] = newIdx
			} else {
				newIdx = len(deduped)
				bigints[c.Value] = newIdx
				indexMap[curIdx] = newIdx
				deduped = append(deduped, c)
			}
		case *String:
			if newIdx, ok := strings[c.Value]; ok {
				indexMap[curIdx] = newIdx
			} else {
				newIdx = len(deduped)
				strings[c.Value] = newIdx
				indexMap[curIdx] = newIdx
				deduped = append(deduped, c)
			}
		case *Float:
			if newIdx, ok := floats[c.Value]; ok {
				indexMap[curIdx] = newIdx
			} else {
				newIdx = len(deduped)
				floats[c.Value] = newIdx
				indexMap[curIdx] = newIdx
				deduped = append(deduped, c)
			}
		case *BigFloat:
			if newIdx, ok := bigfloats[c.Value]; ok {
				indexMap[curIdx] = newIdx
			} else {
				newIdx = len(deduped)
				bigfloats[c.Value] = newIdx
				indexMap[curIdx] = newIdx
				deduped = append(deduped, c)
			}	
		case *Complex:
			if newIdx, ok := complexs[c.Value]; ok {
				indexMap[curIdx] = newIdx
			} else {
				newIdx = len(deduped)
				complexs[c.Value] = newIdx
				indexMap[curIdx] = newIdx
				deduped = append(deduped, c)
			}
		case *Char:
			if newIdx, ok := chars[c.Value]; ok {
				indexMap[curIdx] = newIdx
			} else {
				newIdx = len(deduped)
				chars[c.Value] = newIdx
				indexMap[curIdx] = newIdx
				deduped = append(deduped, c)
			}
		default:
			panic(fmt.Errorf("unsupported top-level constant type: %s",
				c.TypeName()))
		}
	}

	// replace with de-duplicated constants
	b.Constants = deduped

	// update CONST instructions with new indexes
	// main function
	updateConstIndexes(b.MainFunction.Instructions, indexMap)
	// other compiled functions in constants
	for _, c := range b.Constants {
		switch c := c.(type) {
		case *Function:
			updateConstIndexes(c.Instructions, indexMap)
		}
	}
}

func fixDecodedObject(o Object, modules *ModuleMap) (Object, error) {
	switch o := o.(type) {
	case *Bool:
		if o.IsFalsy() {
			return FalseValue, nil
		}
		return TrueValue, nil
	case *Null:
		return NullValue, nil
	case *Array:
		for i, v := range o.Value {
			fv, err := fixDecodedObject(v, modules)
			if err != nil {
				return nil, err
			}
			o.Value[i] = fv
		}
	case *ImmutableArray:
		for i, v := range o.Value {
			fv, err := fixDecodedObject(v, modules)
			if err != nil {
				return nil, err
			}
			o.Value[i] = fv
		}
	case *Map:
		for k, v := range o.Value {
			fv, err := fixDecodedObject(v, modules)
			if err != nil {
				return nil, err
			}
			o.Value[k] = fv
		}
	case *Struct:
		for k, v := range o.Fields {
			fv, err := fixDecodedObject(v, modules)
			if err != nil {
				return nil, err
			}
			o.Fields[k] = fv
		}
	case *ImmutableMap:
		modName := inferModuleName(o)
		if mod := modules.GetBuiltinModule(modName); mod != nil {
			return mod.AsImmutableMap(modName), nil
		}

		for k, v := range o.Value {
			// encoding of user function not supported
			if _, isNativeFunction := v.(*NativeFunction); isNativeFunction {
				return nil, fmt.Errorf("user function not decodable")
			}

			fv, err := fixDecodedObject(v, modules)
			if err != nil {
				return nil, err
			}
			o.Value[k] = fv
		}
	}
	return o, nil
}

func updateConstIndexes(insts []byte, indexMap map[int]int) {
	i := 0
	for i < len(insts) {
		op := insts[i]
		numOperands := parser.OpcodeOperands[op]
		_, read := parser.ReadOperands(numOperands, insts[i+1:])

		switch op {
		case parser.OpConstant:
			curIdx := int(insts[i+2]) | int(insts[i+1])<<8
			newIdx, ok := indexMap[curIdx]
			if !ok {
				panic(fmt.Errorf("constant index not found: %d", curIdx))
			}
			copy(insts[i:], MakeInstruction(op, newIdx))
		case parser.OpClosure:
			curIdx := int(insts[i+2]) | int(insts[i+1])<<8
			numFree := int(insts[i+3])
			newIdx, ok := indexMap[curIdx]
			if !ok {
				panic(fmt.Errorf("constant index not found: %d", curIdx))
			}
			copy(insts[i:], MakeInstruction(op, newIdx, numFree))
		}

		i += 1 + read
	}
}

func inferModuleName(mod *ImmutableMap) string {
	if modName, ok := mod.Value["__module_name__"].(*String); ok {
		return modName.Value
	}
	return ""
}

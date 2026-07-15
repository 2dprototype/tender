package tender

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"math/big"
	"reflect"
	"time"

	"github.com/2dprototype/tender/parser"
)

type ObjectType byte

const (
	tagNull           ObjectType = 1
	tagBool           ObjectType = 2
	tagInt            ObjectType = 3
	tagFloat          ObjectType = 4
	tagChar           ObjectType = 5
	tagString         ObjectType = 6
	tagBytes          ObjectType = 7
	tagArray          ObjectType = 8
	tagMap            ObjectType = 9
	tagImmutableArray ObjectType = 10
	tagImmutableMap   ObjectType = 11
	tagFunction       ObjectType = 12
	tagStructType     ObjectType = 13
	tagStruct         ObjectType = 14
	tagTuple          ObjectType = 15
	tagBigInt         ObjectType = 16
	tagBigFloat       ObjectType = 17
	tagComplex        ObjectType = 18
	tagError          ObjectType = 19
	tagTime           ObjectType = 20
	tagNativeFunction ObjectType = 21
	tagBoundMethod    ObjectType = 22
)

type serializerFunc func(w *BinaryWriter, obj Object) error
type deserializerFunc func(r *BinaryReader, modules *ModuleMap) (Object, error)

type typeRegistry struct {
	tagToType   map[ObjectType]reflect.Type
	typeToTag   map[reflect.Type]ObjectType
	serializers map[reflect.Type]serializerFunc
	decoders    map[ObjectType]deserializerFunc
}

var registry = typeRegistry{
	tagToType:   make(map[ObjectType]reflect.Type),
	typeToTag:   make(map[reflect.Type]ObjectType),
	serializers: make(map[reflect.Type]serializerFunc),
	decoders:    make(map[ObjectType]deserializerFunc),
}

func RegisterSerializationType(t reflect.Type, tag ObjectType, ser serializerFunc, dec deserializerFunc) {
	registry.tagToType[tag] = t
	registry.typeToTag[t] = tag
	registry.serializers[t] = ser
	registry.decoders[tag] = dec
}

func init() {
	// Null
	RegisterSerializationType(reflect.TypeOf((*Null)(nil)), tagNull,
		func(w *BinaryWriter, obj Object) error {
			return nil
		},
		func(r *BinaryReader, modules *ModuleMap) (Object, error) {
			return NullValue, nil
		},
	)

	// Bool
	RegisterSerializationType(reflect.TypeOf((*Bool)(nil)), tagBool,
		func(w *BinaryWriter, obj Object) error {
			return w.WriteBool(!obj.(*Bool).IsFalsy())
		},
		func(r *BinaryReader, modules *ModuleMap) (Object, error) {
			val, err := r.ReadBool()
			if err != nil {
				return nil, err
			}
			if val {
				return TrueValue, nil
			}
			return FalseValue, nil
		},
	)

	// Int
	RegisterSerializationType(reflect.TypeOf((*Int)(nil)), tagInt,
		func(w *BinaryWriter, obj Object) error {
			return w.WriteVarint(obj.(*Int).Value)
		},
		func(r *BinaryReader, modules *ModuleMap) (Object, error) {
			val, err := r.ReadVarint()
			if err != nil {
				return nil, err
			}
			return &Int{Value: val}, nil
		},
	)

	// Float
	RegisterSerializationType(reflect.TypeOf((*Float)(nil)), tagFloat,
		func(w *BinaryWriter, obj Object) error {
			bits := math.Float64bits(obj.(*Float).Value)
			var buf [8]byte
			binary.LittleEndian.PutUint64(buf[:], bits)
			_, err := w.w.Write(buf[:])
			return err
		},
		func(r *BinaryReader, modules *ModuleMap) (Object, error) {
			var buf [8]byte
			for i := 0; i < 8; i++ {
				b, err := r.r.ReadByte()
				if err != nil {
					return nil, err
				}
				buf[i] = b
			}
			bits := binary.LittleEndian.Uint64(buf[:])
			return &Float{Value: math.Float64frombits(bits)}, nil
		},
	)

	// Char
	RegisterSerializationType(reflect.TypeOf((*Char)(nil)), tagChar,
		func(w *BinaryWriter, obj Object) error {
			return w.WriteVarint(int64(obj.(*Char).Value))
		},
		func(r *BinaryReader, modules *ModuleMap) (Object, error) {
			val, err := r.ReadVarint()
			if err != nil {
				return nil, err
			}
			return &Char{Value: rune(val)}, nil
		},
	)

	// String
	RegisterSerializationType(reflect.TypeOf((*String)(nil)), tagString,
		func(w *BinaryWriter, obj Object) error {
			return w.WriteString(obj.(*String).Value)
		},
		func(r *BinaryReader, modules *ModuleMap) (Object, error) {
			val, err := r.ReadString()
			if err != nil {
				return nil, err
			}
			return &String{Value: val}, nil
		},
	)

	// Bytes
	RegisterSerializationType(reflect.TypeOf((*Bytes)(nil)), tagBytes,
		func(w *BinaryWriter, obj Object) error {
			return w.WriteBytes(obj.(*Bytes).Value)
		},
		func(r *BinaryReader, modules *ModuleMap) (Object, error) {
			val, err := r.ReadBytes()
			if err != nil {
				return nil, err
			}
			return &Bytes{Value: val}, nil
		},
	)

	// Array
	RegisterSerializationType(reflect.TypeOf((*Array)(nil)), tagArray,
		func(w *BinaryWriter, obj Object) error {
			arr := obj.(*Array).Value
			if err := w.WriteUvarint(uint64(len(arr))); err != nil {
				return err
			}
			for _, elem := range arr {
				if err := w.WriteObject(elem); err != nil {
					return err
				}
			}
			return nil
		},
		func(r *BinaryReader, modules *ModuleMap) (Object, error) {
			length, err := r.ReadUvarint()
			if err != nil {
				return nil, err
			}
			arr := make([]Object, length)
			for i := uint64(0); i < length; i++ {
				elem, err := r.ReadObject(modules)
				if err != nil {
					return nil, err
				}
				arr[i] = elem
			}
			return &Array{Value: arr}, nil
		},
	)

	// Map
	RegisterSerializationType(reflect.TypeOf((*Map)(nil)), tagMap,
		func(w *BinaryWriter, obj Object) error {
			m := obj.(*Map).Value
			if err := w.WriteUvarint(uint64(len(m))); err != nil {
				return err
			}
			for k, v := range m {
				if err := w.WriteString(k); err != nil {
					return err
				}
				if err := w.WriteObject(v); err != nil {
					return err
				}
			}
			return nil
		},
		func(r *BinaryReader, modules *ModuleMap) (Object, error) {
			length, err := r.ReadUvarint()
			if err != nil {
				return nil, err
			}
			m := make(map[string]Object, length)
			for i := uint64(0); i < length; i++ {
				k, err := r.ReadString()
				if err != nil {
					return nil, err
				}
				v, err := r.ReadObject(modules)
				if err != nil {
					return nil, err
				}
				m[k] = v
			}
			return &Map{Value: m}, nil
		},
	)

	// ImmutableArray
	RegisterSerializationType(reflect.TypeOf((*ImmutableArray)(nil)), tagImmutableArray,
		func(w *BinaryWriter, obj Object) error {
			arr := obj.(*ImmutableArray).Value
			if err := w.WriteUvarint(uint64(len(arr))); err != nil {
				return err
			}
			for _, elem := range arr {
				if err := w.WriteObject(elem); err != nil {
					return err
				}
			}
			return nil
		},
		func(r *BinaryReader, modules *ModuleMap) (Object, error) {
			length, err := r.ReadUvarint()
			if err != nil {
				return nil, err
			}
			arr := make([]Object, length)
			for i := uint64(0); i < length; i++ {
				elem, err := r.ReadObject(modules)
				if err != nil {
					return nil, err
				}
				arr[i] = elem
			}
			return &ImmutableArray{Value: arr}, nil
		},
	)

	// ImmutableMap
	RegisterSerializationType(reflect.TypeOf((*ImmutableMap)(nil)), tagImmutableMap,
		func(w *BinaryWriter, obj Object) error {
			m := obj.(*ImmutableMap).Value
			if err := w.WriteUvarint(uint64(len(m))); err != nil {
				return err
			}
			for k, v := range m {
				if err := w.WriteString(k); err != nil {
					return err
				}
				if err := w.WriteObject(v); err != nil {
					return err
				}
			}
			return nil
		},
		func(r *BinaryReader, modules *ModuleMap) (Object, error) {
			length, err := r.ReadUvarint()
			if err != nil {
				return nil, err
			}
			m := make(map[string]Object, length)
			for i := uint64(0); i < length; i++ {
				k, err := r.ReadString()
				if err != nil {
					return nil, err
				}
				v, err := r.ReadObject(modules)
				if err != nil {
					return nil, err
				}
				m[k] = v
			}
			return &ImmutableMap{Value: m}, nil
		},
	)

	// Function
	RegisterSerializationType(reflect.TypeOf((*Function)(nil)), tagFunction,
		func(w *BinaryWriter, obj Object) error {
			return w.WriteFunction(obj.(*Function))
		},
		func(r *BinaryReader, modules *ModuleMap) (Object, error) {
			return r.ReadFunction()
		},
	)

	// StructType
	RegisterSerializationType(reflect.TypeOf((*StructType)(nil)), tagStructType,
		func(w *BinaryWriter, obj Object) error {
			st := obj.(*StructType)
			if err := w.WriteString(st.Name); err != nil {
				return err
			}
			if err := w.WriteUvarint(uint64(len(st.Fields))); err != nil {
				return err
			}
			for _, f := range st.Fields {
				if err := w.WriteString(f.Name); err != nil {
					return err
				}
				if err := w.WriteString(f.Type); err != nil {
					return err
				}
				if err := w.WriteString(f.Tag); err != nil {
					return err
				}
			}
			if err := w.WriteUvarint(uint64(len(st.Methods))); err != nil {
				return err
			}
			for k, v := range st.Methods {
				if err := w.WriteString(k); err != nil {
					return err
				}
				if err := w.WriteObject(v); err != nil {
					return err
				}
			}
			return nil
		},
		func(r *BinaryReader, modules *ModuleMap) (Object, error) {
			name, err := r.ReadString()
			if err != nil {
				return nil, err
			}
			numFields, err := r.ReadUvarint()
			if err != nil {
				return nil, err
			}
			fields := make([]StructField, numFields)
			for i := uint64(0); i < numFields; i++ {
				fName, err := r.ReadString()
				if err != nil {
					return nil, err
				}
				fType, err := r.ReadString()
				if err != nil {
					return nil, err
				}
				fTag, err := r.ReadString()
				if err != nil {
					return nil, err
				}
				fields[i] = StructField{Name: fName, Type: fType, Tag: fTag}
			}
			numMethods, err := r.ReadUvarint()
			if err != nil {
				return nil, err
			}
			methods := make(map[string]Object, numMethods)
			for i := uint64(0); i < numMethods; i++ {
				k, err := r.ReadString()
				if err != nil {
					return nil, err
				}
				v, err := r.ReadObject(modules)
				if err != nil {
					return nil, err
				}
				methods[k] = v
			}
			return &StructType{
				Name:    name,
				Fields:  fields,
				Methods: methods,
			}, nil
		},
	)

	// Struct
	RegisterSerializationType(reflect.TypeOf((*Struct)(nil)), tagStruct,
		func(w *BinaryWriter, obj Object) error {
			s := obj.(*Struct)
			if err := w.WriteObject(s.Type); err != nil {
				return err
			}
			if err := w.WriteUvarint(uint64(len(s.Fields))); err != nil {
				return err
			}
			for k, v := range s.Fields {
				if err := w.WriteString(k); err != nil {
					return err
				}
				if err := w.WriteObject(v); err != nil {
					return err
				}
			}
			return nil
		},
		func(r *BinaryReader, modules *ModuleMap) (Object, error) {
			stObj, err := r.ReadObject(modules)
			if err != nil {
				return nil, err
			}
			st, ok := stObj.(*StructType)
			if !ok {
				return nil, fmt.Errorf("expected StructType, got %T", stObj)
			}
			numFields, err := r.ReadUvarint()
			if err != nil {
				return nil, err
			}
			fields := make(map[string]Object, numFields)
			for i := uint64(0); i < numFields; i++ {
				k, err := r.ReadString()
				if err != nil {
					return nil, err
				}
				v, err := r.ReadObject(modules)
				if err != nil {
					return nil, err
				}
				fields[k] = v
			}
			return &Struct{
				Type:   st,
				Fields: fields,
			}, nil
		},
	)

	// Tuple
	RegisterSerializationType(reflect.TypeOf((*Tuple)(nil)), tagTuple,
		func(w *BinaryWriter, obj Object) error {
			arr := obj.(*Tuple).Value
			if err := w.WriteUvarint(uint64(len(arr))); err != nil {
				return err
			}
			for _, elem := range arr {
				if err := w.WriteObject(elem); err != nil {
					return err
				}
			}
			return nil
		},
		func(r *BinaryReader, modules *ModuleMap) (Object, error) {
			length, err := r.ReadUvarint()
			if err != nil {
				return nil, err
			}
			arr := make([]Object, length)
			for i := uint64(0); i < length; i++ {
				elem, err := r.ReadObject(modules)
				if err != nil {
					return nil, err
				}
				arr[i] = elem
			}
			return &Tuple{Value: arr}, nil
		},
	)

	// BigInt
	RegisterSerializationType(reflect.TypeOf((*BigInt)(nil)), tagBigInt,
		func(w *BinaryWriter, obj Object) error {
			bi := obj.(*BigInt)
			if err := w.WriteVarint(int64(bi.Value.Sign())); err != nil {
				return err
			}
			return w.WriteBytes(bi.Value.Bytes())
		},
		func(r *BinaryReader, modules *ModuleMap) (Object, error) {
			sign, err := r.ReadVarint()
			if err != nil {
				return nil, err
			}
			bytes, err := r.ReadBytes()
			if err != nil {
				return nil, err
			}
			val := new(big.Int).SetBytes(bytes)
			if sign < 0 {
				val.Neg(val)
			}
			return &BigInt{Value: val}, nil
		},
	)

	// BigFloat
	RegisterSerializationType(reflect.TypeOf((*BigFloat)(nil)), tagBigFloat,
		func(w *BinaryWriter, obj Object) error {
			bf := obj.(*BigFloat)
			text, err := bf.Value.MarshalText()
			if err != nil {
				return err
			}
			return w.WriteBytes(text)
		},
		func(r *BinaryReader, modules *ModuleMap) (Object, error) {
			text, err := r.ReadBytes()
			if err != nil {
				return nil, err
			}
			val := new(big.Float)
			if err := val.UnmarshalText(text); err != nil {
				return nil, err
			}
			return &BigFloat{Value: val}, nil
		},
	)

	// Complex
	RegisterSerializationType(reflect.TypeOf((*Complex)(nil)), tagComplex,
		func(w *BinaryWriter, obj Object) error {
			c := obj.(*Complex)
			re := real(c.Value)
			im := imag(c.Value)

			// Write real
			bitsRe := math.Float64bits(re)
			var bufRe [8]byte
			binary.LittleEndian.PutUint64(bufRe[:], bitsRe)
			if _, err := w.w.Write(bufRe[:]); err != nil {
				return err
			}

			// Write imag
			bitsIm := math.Float64bits(im)
			var bufIm [8]byte
			binary.LittleEndian.PutUint64(bufIm[:], bitsIm)
			if _, err := w.w.Write(bufIm[:]); err != nil {
				return err
			}
			return nil
		},
		func(r *BinaryReader, modules *ModuleMap) (Object, error) {
			// Read real
			var bufRe [8]byte
			for i := 0; i < 8; i++ {
				b, err := r.r.ReadByte()
				if err != nil {
					return nil, err
				}
				bufRe[i] = b
			}
			bitsRe := binary.LittleEndian.Uint64(bufRe[:])
			re := math.Float64frombits(bitsRe)

			// Read imag
			var bufIm [8]byte
			for i := 0; i < 8; i++ {
				b, err := r.r.ReadByte()
				if err != nil {
					return nil, err
				}
				bufIm[i] = b
			}
			bitsIm := binary.LittleEndian.Uint64(bufIm[:])
			im := math.Float64frombits(bitsIm)

			return &Complex{Value: complex(re, im)}, nil
		},
	)

	// Error
	RegisterSerializationType(reflect.TypeOf((*Error)(nil)), tagError,
		func(w *BinaryWriter, obj Object) error {
			return w.WriteObject(obj.(*Error).Value)
		},
		func(r *BinaryReader, modules *ModuleMap) (Object, error) {
			val, err := r.ReadObject(modules)
			if err != nil {
				return nil, err
			}
			return &Error{Value: val}, nil
		},
	)

	// Time
	RegisterSerializationType(reflect.TypeOf((*Time)(nil)), tagTime,
		func(w *BinaryWriter, obj Object) error {
			return w.WriteVarint(obj.(*Time).Value.UnixNano())
		},
		func(r *BinaryReader, modules *ModuleMap) (Object, error) {
			nano, err := r.ReadVarint()
			if err != nil {
				return nil, err
			}
			return &Time{Value: time.Unix(0, nano)}, nil
		},
	)

	// NativeFunction
	RegisterSerializationType(reflect.TypeOf((*NativeFunction)(nil)), tagNativeFunction,
		func(w *BinaryWriter, obj Object) error {
			nf := obj.(*NativeFunction)
			if err := w.WriteString(nf.Name); err != nil {
				return err
			}
			return w.WriteBool(nf.NeedVMObj)
		},
		func(r *BinaryReader, modules *ModuleMap) (Object, error) {
			name, err := r.ReadString()
			if err != nil {
				return nil, err
			}
			needVM, err := r.ReadBool()
			if err != nil {
				return nil, err
			}
			return &NativeFunction{
				Name:      name,
				NeedVMObj: needVM,
			}, nil
		},
	)

	// BoundMethod
	RegisterSerializationType(reflect.TypeOf((*BoundMethod)(nil)), tagBoundMethod,
		func(w *BinaryWriter, obj Object) error {
			bm := obj.(*BoundMethod)
			if err := w.WriteObject(bm.Receiver); err != nil {
				return err
			}
			return w.WriteObject(bm.Func)
		},
		func(r *BinaryReader, modules *ModuleMap) (Object, error) {
			receiver, err := r.ReadObject(modules)
			if err != nil {
				return nil, err
			}
			fn, err := r.ReadObject(modules)
			if err != nil {
				return nil, err
			}
			return &BoundMethod{
				Receiver: receiver,
				Func:     fn,
			}, nil
		},
	)
}

type BinaryWriter struct {
	w io.Writer
}

func (w *BinaryWriter) WriteVarint(v int64) error {
	var buf [binary.MaxVarintLen64]byte
	n := binary.PutVarint(buf[:], v)
	_, err := w.w.Write(buf[:n])
	return err
}

func (w *BinaryWriter) WriteUvarint(v uint64) error {
	var buf [binary.MaxVarintLen64]byte
	n := binary.PutUvarint(buf[:], v)
	_, err := w.w.Write(buf[:n])
	return err
}

func (w *BinaryWriter) WriteString(s string) error {
	if err := w.WriteUvarint(uint64(len(s))); err != nil {
		return err
	}
	_, err := w.w.Write([]byte(s))
	return err
}

func (w *BinaryWriter) WriteBytes(b []byte) error {
	if err := w.WriteUvarint(uint64(len(b))); err != nil {
		return err
	}
	_, err := w.w.Write(b)
	return err
}

func (w *BinaryWriter) WriteBool(b bool) error {
	val := byte(0)
	if b {
		val = 1
	}
	_, err := w.w.Write([]byte{val})
	return err
}

func (w *BinaryWriter) WriteFileSet(fs *parser.SourceFileSet) error {
	if fs == nil {
		return w.WriteVarint(0)
	}
	if err := w.WriteVarint(int64(fs.Base)); err != nil {
		return err
	}
	if err := w.WriteUvarint(uint64(len(fs.Files))); err != nil {
		return err
	}
	for _, f := range fs.Files {
		if err := w.WriteString(f.Name); err != nil {
			return err
		}
		if err := w.WriteVarint(int64(f.Base)); err != nil {
			return err
		}
		if err := w.WriteVarint(int64(f.Size)); err != nil {
			return err
		}
		if err := w.WriteUvarint(uint64(len(f.Lines))); err != nil {
			return err
		}
		for _, line := range f.Lines {
			if err := w.WriteVarint(int64(line)); err != nil {
				return err
			}
		}
		if err := w.WriteBytes(f.Src); err != nil {
			return err
		}
	}
	return nil
}

func (w *BinaryWriter) WriteFunction(fn *Function) error {
	if fn == nil {
		return w.WriteBool(false)
	}
	if err := w.WriteBool(true); err != nil {
		return err
	}
	if err := w.WriteBytes(fn.Instructions); err != nil {
		return err
	}
	if err := w.WriteVarint(int64(fn.NumLocals)); err != nil {
		return err
	}
	if err := w.WriteVarint(int64(fn.NumParameters)); err != nil {
		return err
	}
	if err := w.WriteBool(fn.VarArgs); err != nil {
		return err
	}
	if err := w.WriteUvarint(uint64(len(fn.SourceMap))); err != nil {
		return err
	}
	for k, v := range fn.SourceMap {
		if err := w.WriteVarint(int64(k)); err != nil {
			return err
		}
		if err := w.WriteVarint(int64(v)); err != nil {
			return err
		}
	}
	return nil
}

func (w *BinaryWriter) WriteObject(obj Object) error {
	if obj == nil {
		_, err := w.w.Write([]byte{byte(tagNull)})
		return err
	}
	t := reflect.TypeOf(obj)
	tag, registered := registry.typeToTag[t]
	if !registered {
		return fmt.Errorf("unsupported object type for binary serialization: %s", obj.TypeName())
	}
	if _, err := w.w.Write([]byte{byte(tag)}); err != nil {
		return err
	}
	return registry.serializers[t](w, obj)
}

type BinaryReader struct {
	raw io.Reader
	r   io.ByteReader
}

func NewBinaryReader(r io.Reader) *BinaryReader {
	if br, ok := r.(io.ByteReader); ok {
		return &BinaryReader{raw: r, r: br}
	}
	buf := bufio.NewReader(r)
	return &BinaryReader{raw: buf, r: buf}
}

func (r *BinaryReader) ReadVarint() (int64, error) {
	return binary.ReadVarint(r.r)
}

func (r *BinaryReader) ReadUvarint() (uint64, error) {
	return binary.ReadUvarint(r.r)
}

func (r *BinaryReader) ReadBool() (bool, error) {
	b, err := r.r.ReadByte()
	if err != nil {
		return false, err
	}
	return b != 0, nil
}

func (r *BinaryReader) ReadBytes() ([]byte, error) {
	length, err := r.ReadUvarint()
	if err != nil {
		return nil, err
	}
	buf := make([]byte, length)
	_, err = io.ReadFull(r.raw, buf)
	return buf, err
}

func (r *BinaryReader) ReadString() (string, error) {
	b, err := r.ReadBytes()
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (r *BinaryReader) ReadFileSet() (*parser.SourceFileSet, error) {
	base, err := r.ReadVarint()
	if err != nil {
		return nil, err
	}
	if base == 0 {
		return nil, nil
	}
	fs := &parser.SourceFileSet{
		Base: int(base),
	}
	numFiles, err := r.ReadUvarint()
	if err != nil {
		return nil, err
	}
	fs.Files = make([]*parser.SourceFile, numFiles)
	for i := uint64(0); i < numFiles; i++ {
		name, err := r.ReadString()
		if err != nil {
			return nil, err
		}
		fBase, err := r.ReadVarint()
		if err != nil {
			return nil, err
		}
		fSize, err := r.ReadVarint()
		if err != nil {
			return nil, err
		}
		numLines, err := r.ReadUvarint()
		if err != nil {
			return nil, err
		}
		lines := make([]int, numLines)
		for j := uint64(0); j < numLines; j++ {
			line, err := r.ReadVarint()
			if err != nil {
				return nil, err
			}
			lines[j] = int(line)
		}
		src, err := r.ReadBytes()
		if err != nil {
			return nil, err
		}
		f := &parser.SourceFile{
			Name:  name,
			Base:  int(fBase),
			Size:  int(fSize),
			Lines: lines,
			Src:   src,
		}
		f.SetFileSet(fs)
		fs.Files[i] = f
	}
	if len(fs.Files) > 0 {
		fs.LastFile = fs.Files[len(fs.Files)-1]
	}
	return fs, nil
}

func (r *BinaryReader) ReadFunction() (*Function, error) {
	hasFn, err := r.ReadBool()
	if err != nil {
		return nil, err
	}
	if !hasFn {
		return nil, nil
	}
	insts, err := r.ReadBytes()
	if err != nil {
		return nil, err
	}
	numLocals, err := r.ReadVarint()
	if err != nil {
		return nil, err
	}
	numParams, err := r.ReadVarint()
	if err != nil {
		return nil, err
	}
	varArgs, err := r.ReadBool()
	if err != nil {
		return nil, err
	}
	sourceMapLen, err := r.ReadUvarint()
	if err != nil {
		return nil, err
	}
	sourceMap := make(map[int]parser.Pos, sourceMapLen)
	for i := uint64(0); i < sourceMapLen; i++ {
		k, err := r.ReadVarint()
		if err != nil {
			return nil, err
		}
		v, err := r.ReadVarint()
		if err != nil {
			return nil, err
		}
		sourceMap[int(k)] = parser.Pos(v)
	}
	return &Function{
		Instructions:  insts,
		NumLocals:     int(numLocals),
		NumParameters: int(numParams),
		VarArgs:       varArgs,
		SourceMap:     sourceMap,
	}, nil
}

func (r *BinaryReader) ReadObject(modules *ModuleMap) (Object, error) {
	tagByte, err := r.r.ReadByte()
	if err != nil {
		return nil, err
	}
	tag := ObjectType(tagByte)
	decoder, registered := registry.decoders[tag]
	if !registered {
		return nil, fmt.Errorf("unsupported object tag for binary deserialization: %d", tag)
	}
	return decoder(r, modules)
}

func (b *Bytecode) EncodeTDC(w io.Writer, version string) error {
	bw := &BinaryWriter{w: w}

	// Write magic header: 'T', 'D', 'C', 1
	if _, err := w.Write([]byte{'T', 'D', 'C', 1}); err != nil {
		return err
	}

	// Write version string
	if err := bw.WriteString(version); err != nil {
		return err
	}

	// Write compile timestamp (UnixNano)
	if err := bw.WriteVarint(time.Now().UnixNano()); err != nil {
		return err
	}

	// Serialize FileSet
	if err := bw.WriteFileSet(b.FileSet); err != nil {
		return err
	}

	// Serialize MainFunction
	if err := bw.WriteFunction(b.MainFunction); err != nil {
		return err
	}

	// Serialize Constants
	if err := bw.WriteUvarint(uint64(len(b.Constants))); err != nil {
		return err
	}
	for _, c := range b.Constants {
		if err := bw.WriteObject(c); err != nil {
			return err
		}
	}

	return nil
}

func (b *Bytecode) DecodeTDC(r io.Reader, modules *ModuleMap) (version string, timestamp int64, err error) {
	if modules == nil {
		modules = NewModuleMap()
	}

	br := NewBinaryReader(r)

	// Read magic header
	var magic [4]byte
	_, err = io.ReadFull(br.raw, magic[:])
	if err != nil {
		return "", 0, err
	}
	if magic != [4]byte{'T', 'D', 'C', 1} {
		return "", 0, fmt.Errorf("invalid magic header")
	}

	// Read version string
	version, err = br.ReadString()
	if err != nil {
		return "", 0, err
	}

	// Read compile timestamp
	timestamp, err = br.ReadVarint()
	if err != nil {
		return "", 0, err
	}

	// Deserialize FileSet
	b.FileSet, err = br.ReadFileSet()
	if err != nil {
		return "", 0, err
	}

	// Deserialize MainFunction
	b.MainFunction, err = br.ReadFunction()
	if err != nil {
		return "", 0, err
	}

	// Deserialize Constants
	numConstants, err := br.ReadUvarint()
	if err != nil {
		return "", 0, err
	}
	b.Constants = make([]Object, numConstants)
	for i := uint64(0); i < numConstants; i++ {
		c, err := br.ReadObject(modules)
		if err != nil {
			return "", 0, err
		}
		b.Constants[i] = c
	}

	// Fix decoded constants
	for i, v := range b.Constants {
		fv, err := fixDecodedObject(v, modules)
		if err != nil {
			return "", 0, err
		}
		b.Constants[i] = fv
	}

	return version, timestamp, nil
}

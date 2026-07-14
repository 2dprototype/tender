package stdlib

import (
	"encoding/hex"
	"github.com/2dprototype/tender"
)

var hexModule = map[string]tender.Object{
	"encode": &tender.NativeFunction{Value: FuncAYRS(hex.EncodeToString)},
	"decode": &tender.NativeFunction{Value: FuncASRYE(hex.DecodeString)},
	"dump": &tender.NativeFunction{Value: FuncAYRS(hex.Dump)},
}

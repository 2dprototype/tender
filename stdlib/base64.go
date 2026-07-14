package stdlib

import (
	"encoding/base64"

	"github.com/2dprototype/tender"
)

var base64Module = map[string]tender.Object{
	"encode": &tender.NativeFunction{
		Value: FuncAYRS(base64.StdEncoding.EncodeToString),
	},
	"decode": &tender.NativeFunction{
		Value: FuncASRYE(base64.StdEncoding.DecodeString),
	},
	"raw_encode": &tender.NativeFunction{
		Value: FuncAYRS(base64.RawStdEncoding.EncodeToString),
	},
	"raw_decode": &tender.NativeFunction{
		Value: FuncASRYE(base64.RawStdEncoding.DecodeString),
	},
	"url_encode": &tender.NativeFunction{
		Value: FuncAYRS(base64.URLEncoding.EncodeToString),
	},
	"url_decode": &tender.NativeFunction{
		Value: FuncASRYE(base64.URLEncoding.DecodeString),
	},
	"raw_url_encode": &tender.NativeFunction{
		Value: FuncAYRS(base64.RawURLEncoding.EncodeToString),
	},
	"raw_url_decode": &tender.NativeFunction{
		Value: FuncASRYE(base64.RawURLEncoding.DecodeString),
	},
}

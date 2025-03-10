package stdlib

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/2dprototype/tender"
)

var stringsModule = map[string]tender.Object{
	"re_match": &tender.UserFunction{
		Name:  "re_match",
		Value: stringsREMatch,
	}, // re_match(pattern, strings) => bool/error
	"re_find": &tender.UserFunction{
		Name:  "re_find",
		Value: stringsREFind,
	}, // re_find(pattern, strings, count) => [[{strings:,begin:,end:}]]/null
	"re_replace": &tender.UserFunction{
		Name:  "re_replace",
		Value: stringsREReplace,
	}, // re_replace(pattern, strings, repl) => string/error
	"re_split": &tender.UserFunction{
		Name:  "re_split",
		Value: stringsRESplit,
	}, // re_split(pattern, strings, count) => [string]/error
	"re_compile": &tender.UserFunction{
		Name:  "re_compile",
		Value: stringsRECompile,
	}, // re_compile(pattern) => Regexp/error
	"compare": &tender.UserFunction{
		Name:  "compare",
		Value: FuncASSRI(strings.Compare),
	}, // compare(a, b) => int
	"contains": &tender.UserFunction{
		Name:  "contains",
		Value: FuncASSRB(strings.Contains),
	}, // contains(s, substr) => bool
	"contains_any": &tender.UserFunction{
		Name:  "contains_any",
		Value: FuncASSRB(strings.ContainsAny),
	}, // contains_any(s, chars) => bool
	"count": &tender.UserFunction{
		Name:  "count",
		Value: FuncASSRI(strings.Count),
	}, // count(s, substr) => int
	"equal_fold": &tender.UserFunction{
		Name:  "equal_fold",
		Value: FuncASSRB(strings.EqualFold),
	}, // "equal_fold(s, t) => bool
	"fields": &tender.UserFunction{
		Name:  "fields",
		Value: FuncASRSs(strings.Fields),
	}, // fields(s) => [string]
	"has_prefix": &tender.UserFunction{
		Name:  "has_prefix",
		Value: FuncASSRB(strings.HasPrefix),
	}, // has_prefix(s, prefix) => bool
	"has_suffix": &tender.UserFunction{
		Name:  "has_suffix",
		Value: FuncASSRB(strings.HasSuffix),
	}, // has_suffix(s, suffix) => bool
	"index": &tender.UserFunction{
		Name:  "index",
		Value: FuncASSRI(strings.Index),
	}, // index(s, substr) => int
	"index_any": &tender.UserFunction{
		Name:  "index_any",
		Value: FuncASSRI(strings.IndexAny),
	}, // index_any(s, chars) => int
	"join": &tender.UserFunction{
		Name:  "join",
		Value: stringsJoin,
	}, // join(arr, sep) => string
	"last_index": &tender.UserFunction{
		Name:  "last_index",
		Value: FuncASSRI(strings.LastIndex),
	}, // last_index(s, substr) => int
	"last_index_any": &tender.UserFunction{
		Name:  "last_index_any",
		Value: FuncASSRI(strings.LastIndexAny),
	}, // last_index_any(s, chars) => int
	"repeat": &tender.UserFunction{
		Name:  "repeat",
		Value: stringsRepeat,
	}, // repeat(s, count) => string
	"replace": &tender.UserFunction{
		Name:  "replace",
		Value: stringsReplace,
	}, // replace(s, old, new, n) => string
	"substr": &tender.UserFunction{
		Name:  "substr",
		Value: stringsSubstring,
	}, // substr(s, lower, upper) => string
	"split": &tender.UserFunction{
		Name:  "split",
		Value: FuncASSRSs(strings.Split),
	}, // split(s, sep) => [string]
	"split_after": &tender.UserFunction{
		Name:  "split_after",
		Value: FuncASSRSs(strings.SplitAfter),
	}, // split_after(s, sep) => [string]
	"split_after_n": &tender.UserFunction{
		Name:  "split_after_n",
		Value: FuncASSIRSs(strings.SplitAfterN),
	}, // split_after_n(s, sep, n) => [string]
	"split_n": &tender.UserFunction{
		Name:  "split_n",
		Value: FuncASSIRSs(strings.SplitN),
	}, // split_n(s, sep, n) => [string]
	"title": &tender.UserFunction{
		Name:  "title",
		Value: FuncASRS(strings.Title),
	}, // title(s) => string
	"to_lower": &tender.UserFunction{
		Name:  "to_lower",
		Value: FuncASRS(strings.ToLower),
	}, // to_lower(s) => string
	"to_title": &tender.UserFunction{
		Name:  "to_title",
		Value: FuncASRS(strings.ToTitle),
	}, // to_title(s) => string
	"to_upper": &tender.UserFunction{
		Name:  "to_upper",
		Value: FuncASRS(strings.ToUpper),
	}, // to_upper(s) => string
	"pad_left": &tender.UserFunction{
		Name:  "pad_left",
		Value: stringsPadLeft,
	}, // pad_left(s, pad_len, pad_with) => string
	"pad_right": &tender.UserFunction{
		Name:  "pad_right",
		Value: stringsPadRight,
	}, // pad_right(s, pad_len, pad_with) => string
	"trim": &tender.UserFunction{
		Name:  "trim",
		Value: FuncASSRS(strings.Trim),
	}, // trim(s, cutset) => string
	"trim_left": &tender.UserFunction{
		Name:  "trim_left",
		Value: FuncASSRS(strings.TrimLeft),
	}, // trim_left(s, cutset) => string
	"trim_prefix": &tender.UserFunction{
		Name:  "trim_prefix",
		Value: FuncASSRS(strings.TrimPrefix),
	}, // trim_prefix(s, prefix) => string
	"trim_right": &tender.UserFunction{
		Name:  "trim_right",
		Value: FuncASSRS(strings.TrimRight),
	}, // trim_right(s, cutset) => string
	"trim_space": &tender.UserFunction{
		Name:  "trim_space",
		Value: FuncASRS(strings.TrimSpace),
	}, // trim_space(s) => string
	"trim_suffix": &tender.UserFunction{
		Name:  "trim_suffix",
		Value: FuncASSRS(strings.TrimSuffix),
	}, // trim_suffix(s, suffix) => string
	"atoi": &tender.UserFunction{
		Name:  "atoi",
		Value: FuncASRIE(strconv.Atoi),
	}, // atoi(str) => int/error
	"format_bool": &tender.UserFunction{
		Name:  "format_bool",
		Value: stringsFormatBool,
	}, // format_bool(b) => string
	"format_float": &tender.UserFunction{
		Name:  "format_float",
		Value: stringsFormatFloat,
	}, // format_float(f, fmt, prec, bits) => string
	"format_int": &tender.UserFunction{
		Name:  "format_int",
		Value: stringsFormatInt,
	}, // format_int(i, base) => string
	"itoa": &tender.UserFunction{
		Name:  "itoa",
		Value: FuncAIRS(strconv.Itoa),
	}, // itoa(i) => string
	"parse_bool": &tender.UserFunction{
		Name:  "parse_bool",
		Value: stringsParseBool,
	}, // parse_bool(str) => bool/error
	"parse_float": &tender.UserFunction{
		Name:  "parse_float",
		Value: stringsParseFloat,
	}, // parse_float(str, bits) => float/error
	"parse_int": &tender.UserFunction{
		Name:  "parse_int",
		Value: stringsParseInt,
	}, // parse_int(str, base, bits) => int/error
	"quote": &tender.UserFunction{
		Name:  "quote",
		Value: FuncASRS(strconv.Quote),
	}, // quote(str) => string
	"unquote": &tender.UserFunction{
		Name:  "unquote",
		Value: FuncASRSE(strconv.Unquote),
	}, // unquote(str) => string/error
}

func stringsREMatch(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 2 {
		err = tender.ErrWrongNumArguments
		return
	}

	s1, ok := tender.ToString(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := tender.ToString(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	matched, err := regexp.MatchString(s1, s2)
	if err != nil {
		ret = wrapError(err)
		return
	}

	if matched {
		ret = tender.TrueValue
	} else {
		ret = tender.FalseValue
	}

	return
}

func stringsREFind(args ...tender.Object) (ret tender.Object, err error) {
	numArgs := len(args)
	if numArgs != 2 && numArgs != 3 {
		err = tender.ErrWrongNumArguments
		return
	}

	s1, ok := tender.ToString(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	re, err := regexp.Compile(s1)
	if err != nil {
		ret = wrapError(err)
		return
	}

	s2, ok := tender.ToString(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	if numArgs < 3 {
		m := re.FindStringSubmatchIndex(s2)
		if m == nil {
			ret = tender.NullValue
			return
		}

		arr := &tender.Array{}
		for i := 0; i < len(m); i += 2 {
			arr.Value = append(arr.Value,
				&tender.ImmutableMap{Value: map[string]tender.Object{
					"strings":  &tender.String{Value: s2[m[i]:m[i+1]]},
					"begin": &tender.Int{Value: int64(m[i])},
					"end":   &tender.Int{Value: int64(m[i+1])},
				}})
		}

		ret = &tender.Array{Value: []tender.Object{arr}}

		return
	}

	i3, ok := tender.ToInt(args[2])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}
	m := re.FindAllStringSubmatchIndex(s2, i3)
	if m == nil {
		ret = tender.NullValue
		return
	}

	arr := &tender.Array{}
	for _, m := range m {
		subMatch := &tender.Array{}
		for i := 0; i < len(m); i += 2 {
			subMatch.Value = append(subMatch.Value,
				&tender.ImmutableMap{Value: map[string]tender.Object{
					"strings":  &tender.String{Value: s2[m[i]:m[i+1]]},
					"begin": &tender.Int{Value: int64(m[i])},
					"end":   &tender.Int{Value: int64(m[i+1])},
				}})
		}

		arr.Value = append(arr.Value, subMatch)
	}

	ret = arr

	return
}

func stringsREReplace(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 3 {
		err = tender.ErrWrongNumArguments
		return
	}

	s1, ok := tender.ToString(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := tender.ToString(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	s3, ok := tender.ToString(args[2])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "string(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}

	re, err := regexp.Compile(s1)
	if err != nil {
		ret = wrapError(err)
	} else {
		s, ok := doStringsRegexpReplace(re, s2, s3)
		if !ok {
			return nil, tender.ErrStringLimit
		}

		ret = &tender.String{Value: s}
	}

	return
}

func stringsRESplit(args ...tender.Object) (ret tender.Object, err error) {
	numArgs := len(args)
	if numArgs != 2 && numArgs != 3 {
		err = tender.ErrWrongNumArguments
		return
	}

	s1, ok := tender.ToString(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := tender.ToString(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	var i3 = -1
	if numArgs > 2 {
		i3, ok = tender.ToInt(args[2])
		if !ok {
			err = tender.ErrInvalidArgumentType{
				Name:     "third",
				Expected: "int(compatible)",
				Found:    args[2].TypeName(),
			}
			return
		}
	}

	re, err := regexp.Compile(s1)
	if err != nil {
		ret = wrapError(err)
		return
	}

	arr := &tender.Array{}
	for _, s := range re.Split(s2, i3) {
		arr.Value = append(arr.Value, &tender.String{Value: s})
	}

	ret = arr

	return
}

func stringsRECompile(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	s1, ok := tender.ToString(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	re, err := regexp.Compile(s1)
	if err != nil {
		ret = wrapError(err)
	} else {
		ret = makeStringsRegexp(re)
	}

	return
}

func stringsReplace(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 4 {
		err = tender.ErrWrongNumArguments
		return
	}

	s1, ok := tender.ToString(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := tender.ToString(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	s3, ok := tender.ToString(args[2])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "string(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}

	i4, ok := tender.ToInt(args[3])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "int(compatible)",
			Found:    args[3].TypeName(),
		}
		return
	}

	s, ok := doStringsReplace(s1, s2, s3, i4)
	if !ok {
		err = tender.ErrStringLimit
		return
	}

	ret = &tender.String{Value: s}

	return
}

func stringsSubstring(args ...tender.Object) (ret tender.Object, err error) {
	argslen := len(args)
	if argslen != 2 && argslen != 3 {
		err = tender.ErrWrongNumArguments
		return
	}

	s1, ok := tender.ToString(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := tender.ToInt(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	strlen := len(s1)
	i3 := strlen
	if argslen == 3 {
		i3, ok = tender.ToInt(args[2])
		if !ok {
			err = tender.ErrInvalidArgumentType{
				Name:     "third",
				Expected: "int(compatible)",
				Found:    args[2].TypeName(),
			}
			return
		}
	}

	if i2 > i3 {
		err = tender.ErrInvalidIndexType
		return
	}

	if i2 < 0 {
		i2 = 0
	} else if i2 > strlen {
		i2 = strlen
	}

	if i3 < 0 {
		i3 = 0
	} else if i3 > strlen {
		i3 = strlen
	}

	ret = &tender.String{Value: s1[i2:i3]}

	return
}

func stringsPadLeft(args ...tender.Object) (ret tender.Object, err error) {
	argslen := len(args)
	if argslen != 2 && argslen != 3 {
		err = tender.ErrWrongNumArguments
		return
	}

	s1, ok := tender.ToString(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := tender.ToInt(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	if i2 > tender.MaxStringLen {
		return nil, tender.ErrStringLimit
	}

	sLen := len(s1)
	if sLen >= i2 {
		ret = &tender.String{Value: s1}
		return
	}

	s3 := " "
	if argslen == 3 {
		s3, ok = tender.ToString(args[2])
		if !ok {
			err = tender.ErrInvalidArgumentType{
				Name:     "third",
				Expected: "string(compatible)",
				Found:    args[2].TypeName(),
			}
			return
		}
	}

	padStrLen := len(s3)
	if padStrLen == 0 {
		ret = &tender.String{Value: s1}
		return
	}

	padCount := ((i2 - padStrLen) / padStrLen) + 1
	retStr := strings.Repeat(s3, padCount) + s1
	ret = &tender.String{Value: retStr[len(retStr)-i2:]}

	return
}

func stringsPadRight(args ...tender.Object) (ret tender.Object, err error) {
	argslen := len(args)
	if argslen != 2 && argslen != 3 {
		err = tender.ErrWrongNumArguments
		return
	}

	s1, ok := tender.ToString(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := tender.ToInt(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	if i2 > tender.MaxStringLen {
		return nil, tender.ErrStringLimit
	}

	sLen := len(s1)
	if sLen >= i2 {
		ret = &tender.String{Value: s1}
		return
	}

	s3 := " "
	if argslen == 3 {
		s3, ok = tender.ToString(args[2])
		if !ok {
			err = tender.ErrInvalidArgumentType{
				Name:     "third",
				Expected: "string(compatible)",
				Found:    args[2].TypeName(),
			}
			return
		}
	}

	padStrLen := len(s3)
	if padStrLen == 0 {
		ret = &tender.String{Value: s1}
		return
	}

	padCount := ((i2 - padStrLen) / padStrLen) + 1
	retStr := s1 + strings.Repeat(s3, padCount)
	ret = &tender.String{Value: retStr[:i2]}

	return
}

func stringsRepeat(args ...tender.Object) (ret tender.Object, err error) {
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

	i2, ok := tender.ToInt(args[1])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
	}

	if len(s1)*i2 > tender.MaxStringLen {
		return nil, tender.ErrStringLimit
	}

	return &tender.String{Value: strings.Repeat(s1, i2)}, nil
}

func stringsJoin(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 2 {
		return nil, tender.ErrWrongNumArguments
	}

	var slen int
	var ss1 []string
	switch arg0 := args[0].(type) {
	case *tender.Array:
		for idx, a := range arg0.Value {
			as, ok := tender.ToString(a)
			if !ok {
				return nil, tender.ErrInvalidArgumentType{
					Name:     fmt.Sprintf("first[%d]", idx),
					Expected: "string(compatible)",
					Found:    a.TypeName(),
				}
			}
			slen += len(as)
			ss1 = append(ss1, as)
		}
	case *tender.ImmutableArray:
		for idx, a := range arg0.Value {
			as, ok := tender.ToString(a)
			if !ok {
				return nil, tender.ErrInvalidArgumentType{
					Name:     fmt.Sprintf("first[%d]", idx),
					Expected: "string(compatible)",
					Found:    a.TypeName(),
				}
			}
			slen += len(as)
			ss1 = append(ss1, as)
		}
	default:
		return nil, tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "array",
			Found:    args[0].TypeName(),
		}
	}

	s2, ok := tender.ToString(args[1])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
	}

	// make sure output length does not exceed the limit
	if slen+len(s2)*(len(ss1)-1) > tender.MaxStringLen {
		return nil, tender.ErrStringLimit
	}

	return &tender.String{Value: strings.Join(ss1, s2)}, nil
}

func stringsFormatBool(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	b1, ok := args[0].(*tender.Bool)
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bool",
			Found:    args[0].TypeName(),
		}
		return
	}

	if b1 == tender.TrueValue {
		ret = &tender.String{Value: "true"}
	} else {
		ret = &tender.String{Value: "false"}
	}

	return
}

func stringsFormatFloat(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 4 {
		err = tender.ErrWrongNumArguments
		return
	}

	f1, ok := args[0].(*tender.Float)
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "float",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := tender.ToString(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	i3, ok := tender.ToInt(args[2])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}

	i4, ok := tender.ToInt(args[3])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "int(compatible)",
			Found:    args[3].TypeName(),
		}
		return
	}

	ret = &tender.String{Value: strconv.FormatFloat(f1.Value, s2[0], i3, i4)}

	return
}

func stringsFormatInt(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 2 {
		err = tender.ErrWrongNumArguments
		return
	}

	i1, ok := args[0].(*tender.Int)
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := tender.ToInt(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &tender.String{Value: strconv.FormatInt(i1.Value, i2)}

	return
}

func stringsParseBool(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	s1, ok := args[0].(*tender.String)
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
		return
	}

	parsed, err := strconv.ParseBool(s1.Value)
	if err != nil {
		ret = wrapError(err)
		return
	}

	if parsed {
		ret = tender.TrueValue
	} else {
		ret = tender.FalseValue
	}

	return
}

func stringsParseFloat(args ...tender.Object) (tender.Object, error) {
	var err error
	if len(args) != 2 {
		err = tender.ErrWrongNumArguments
		return nil, err
	}

	s1, ok := args[0].(*tender.String)
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
		return nil, err
	}

	i2, ok := tender.ToInt(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return nil, err
	}

	parsed, err := strconv.ParseFloat(s1.Value, i2)
	if err != nil {
		return wrapError(err), nil
	}

	return &tender.Float{Value: parsed}, nil
}

func stringsParseInt(args ...tender.Object) (tender.Object, error) {
	var err error
	if len(args) != 3 {
		err = tender.ErrWrongNumArguments
		return nil, err
	}

	s1, ok := args[0].(*tender.String)
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
		return nil, err
	}

	i2, ok := tender.ToInt(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return nil, err
	}

	i3, ok := tender.ToInt(args[2])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return nil, err
	}

	parsed, err := strconv.ParseInt(s1.Value, i2, i3)
	if err != nil {
		return wrapError(err), nil
	}
	
	return &tender.Int{Value: parsed}, nil
}

// Modified implementation of strings.Replace
// to limit the maximum length of output string.
func doStringsReplace(s, old, new string, n int) (string, bool) {
	if old == new || n == 0 {
		return s, true // avoid allocation
	}

	// Compute number of replacements.
	if m := strings.Count(s, old); m == 0 {
		return s, true // avoid allocation
	} else if n < 0 || m < n {
		n = m
	}

	// Apply replacements to buffer.
	t := make([]byte, len(s)+n*(len(new)-len(old)))
	w := 0
	start := 0
	for i := 0; i < n; i++ {
		j := start
		if len(old) == 0 {
			if i > 0 {
				_, wid := utf8.DecodeRuneInString(s[start:])
				j += wid
			}
		} else {
			j += strings.Index(s[start:], old)
		}

		ssj := s[start:j]
		if w+len(ssj)+len(new) > tender.MaxStringLen {
			return "", false
		}

		w += copy(t[w:], ssj)
		w += copy(t[w:], new)
		start = j + len(old)
	}

	ss := s[start:]
	if w+len(ss) > tender.MaxStringLen {
		return "", false
	}

	w += copy(t[w:], ss)

	return string(t[0:w]), true
}

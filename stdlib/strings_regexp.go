package stdlib

import (
	"regexp"

	"github.com/2dprototype/tender"
)

func makeStringsRegexp(re *regexp.Regexp) *tender.ImmutableMap {
	return &tender.ImmutableMap{
		Value: map[string]tender.Object{
			// match(strings) => bool
			"match": &tender.NativeFunction{
				Value: func(args ...tender.Object) (
					ret tender.Object,
					err error,
				) {
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

					if re.MatchString(s1) {
						ret = tender.TrueValue
					} else {
						ret = tender.FalseValue
					}

					return
				},
			},

			// find(strings) 			=> array(array({strings:,begin:,end:}))/null
			// find(strings, maxCount) => array(array({strings:,begin:,end:}))/null
			"find": &tender.NativeFunction{
				Value: func(args ...tender.Object) (
					ret tender.Object,
					err error,
				) {
					numArgs := len(args)
					if numArgs != 1 && numArgs != 2 {
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

					if numArgs == 1 {
						m := re.FindStringSubmatchIndex(s1)
						if m == nil {
							ret = tender.NullValue
							return
						}

						arr := &tender.Array{}
						for i := 0; i < len(m); i += 2 {
							arr.Value = append(arr.Value,
								&tender.ImmutableMap{
									Value: map[string]tender.Object{
										"strings": &tender.String{
											Value: s1[m[i]:m[i+1]],
										},
										"begin": &tender.Int{
											Value: int64(m[i]),
										},
										"end": &tender.Int{
											Value: int64(m[i+1]),
										},
									}})
						}

						ret = &tender.Array{Value: []tender.Object{arr}}

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
					m := re.FindAllStringSubmatchIndex(s1, i2)
					if m == nil {
						ret = tender.NullValue
						return
					}

					arr := &tender.Array{}
					for _, m := range m {
						subMatch := &tender.Array{}
						for i := 0; i < len(m); i += 2 {
							subMatch.Value = append(subMatch.Value,
								&tender.ImmutableMap{
									Value: map[string]tender.Object{
										"strings": &tender.String{
											Value: s1[m[i]:m[i+1]],
										},
										"begin": &tender.Int{
											Value: int64(m[i]),
										},
										"end": &tender.Int{
											Value: int64(m[i+1]),
										},
									}})
						}

						arr.Value = append(arr.Value, subMatch)
					}

					ret = arr

					return
				},
			},

			// replace(src, repl) => string
			"replace": &tender.NativeFunction{
				Value: func(args ...tender.Object) (
					ret tender.Object,
					err error,
				) {
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

					s, ok := doStringsRegexpReplace(re, s1, s2)
					if !ok {
						return nil, tender.ErrStringLimit
					}

					ret = &tender.String{Value: s}

					return
				},
			},

			// split(strings) 			 => array(string)
			// split(strings, maxCount) => array(string)
			"split": &tender.NativeFunction{
				Value: func(args ...tender.Object) (
					ret tender.Object,
					err error,
				) {
					numArgs := len(args)
					if numArgs != 1 && numArgs != 2 {
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

					var i2 = -1
					if numArgs > 1 {
						i2, ok = tender.ToInt(args[1])
						if !ok {
							err = tender.ErrInvalidArgumentType{
								Name:     "second",
								Expected: "int(compatible)",
								Found:    args[1].TypeName(),
							}
							return
						}
					}

					arr := &tender.Array{}
					for _, s := range re.Split(s1, i2) {
						arr.Value = append(arr.Value,
							&tender.String{Value: s})
					}

					ret = arr

					return
				},
			},
		},
	}
}

// Size-limit checking implementation of regexp.ReplaceAllString.
func doStringsRegexpReplace(re *regexp.Regexp, src, repl string) (string, bool) {
	idx := 0
	out := ""
	for _, m := range re.FindAllStringSubmatchIndex(src, -1) {
		var exp []byte
		exp = re.ExpandString(exp, repl, src, m)
		if len(out)+m[0]-idx+len(exp) > tender.MaxStringLen {
			return "", false
		}
		out += src[idx:m[0]] + string(exp)
		idx = m[1]
	}
	if idx < len(src) {
		if len(out)+len(src)-idx > tender.MaxStringLen {
			return "", false
		}
		out += src[idx:]
	}
	return out, true
}

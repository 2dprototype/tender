package stdlib

import (
	"github.com/2dprototype/tender"
	"github.com/gonutz/wui/v2"
)

func wuiMessageBox(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 2 {
		err = tender.ErrWrongNumArguments
		return
	}

	caption, ok := tender.ToString(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "caption",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	text, ok := tender.ToString(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "text",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	wui.MessageBox(caption, text)
	return tender.NullValue, nil
}

func wuiMessageBoxError(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 2 {
		err = tender.ErrWrongNumArguments
		return
	}

	caption, ok := tender.ToString(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "caption",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	text, ok := tender.ToString(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "text",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	wui.MessageBoxError(caption, text)
	return tender.NullValue, nil
}

func wuiMessageBoxInfo(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 2 {
		err = tender.ErrWrongNumArguments
		return
	}

	caption, ok := tender.ToString(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "caption",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	text, ok := tender.ToString(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "text",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	wui.MessageBoxInfo(caption, text)
	return tender.NullValue, nil
}

func wuiMessageBoxWarning(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 2 {
		err = tender.ErrWrongNumArguments
		return
	}

	caption, ok := tender.ToString(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "caption",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	text, ok := tender.ToString(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "text",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	wui.MessageBoxWarning(caption, text)
	return tender.NullValue, nil
}

func wuiMessageBoxQuestion(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 2 {
		err = tender.ErrWrongNumArguments
		return
	}

	caption, ok := tender.ToString(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "caption",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	text, ok := tender.ToString(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "text",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	wui.MessageBoxQuestion(caption, text)
	return tender.NullValue, nil
}

func wuiMessageBoxOKCancel(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 2 {
		err = tender.ErrWrongNumArguments
		return
	}

	caption, ok := tender.ToString(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "caption",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	text, ok := tender.ToString(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "text",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	result := wui.MessageBoxOKCancel(caption, text)
	if result {
		ret = tender.TrueValue
	} else {
		ret = tender.FalseValue
	}
	return
}

func wuiMessageBoxYesNo(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 2 {
		err = tender.ErrWrongNumArguments
		return
	}

	caption, ok := tender.ToString(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "caption",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	text, ok := tender.ToString(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "text",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	result := wui.MessageBoxYesNo(caption, text)
	if result {
		ret = tender.TrueValue
	} else {
		ret = tender.FalseValue
	}
	return
}

func wuiMessageBoxCustom(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 3 {
		err = tender.ErrWrongNumArguments
		return
	}

	caption, ok := tender.ToString(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "caption",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	text, ok := tender.ToString(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "text",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	flags, ok := tender.ToInt(args[2])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "flags",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}

	result := wui.MessageBoxCustom(caption, text, uint(flags))
	ret = &tender.Int{Value: int64(result)}
	return
}

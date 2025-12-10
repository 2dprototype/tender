package stdlib

import (
	"fmt"
	
	"github.com/2dprototype/tender"
	"github.com/gonutz/wui/v2"
)

type WUIButton struct {
	tender.ObjectImpl
	Value *wui.Button
}

func (b *WUIButton) TypeName() string { return "button" }
func (b *WUIButton) String() string   { return "<button>" }
func (b *WUIButton) Copy() tender.Object {
	return &WUIButton{Value: b.Value}
}

type WUICheckBox struct {
	tender.ObjectImpl
	Value *wui.CheckBox
}

func (c *WUICheckBox) TypeName() string { return "checkbox" }
func (c *WUICheckBox) String() string   { return "<checkbox>" }
func (c *WUICheckBox) Copy() tender.Object {
	return &WUICheckBox{Value: c.Value}
}

type WUILabel struct {
	tender.ObjectImpl
	Value *wui.Label
}

func (l *WUILabel) TypeName() string { return "label" }
func (l *WUILabel) String() string   { return "<label>" }
func (l *WUILabel) Copy() tender.Object {
	return &WUILabel{Value: l.Value}
}

type WUIEditLine struct {
	tender.ObjectImpl
	Value *wui.EditLine
}

func (e *WUIEditLine) TypeName() string { return "editline" }
func (e *WUIEditLine) String() string   { return "<editline>" }
func (e *WUIEditLine) Copy() tender.Object {
	return &WUIEditLine{Value: e.Value}
}

type WUITextEdit struct {
	tender.ObjectImpl
	Value *wui.TextEdit
}

func (t *WUITextEdit) TypeName() string { return "textedit" }
func (t *WUITextEdit) String() string   { return "<textedit>" }
func (t *WUITextEdit) Copy() tender.Object {
	return &WUITextEdit{Value: t.Value}
}

type WUIComboBox struct {
	tender.ObjectImpl
	Value *wui.ComboBox
}

func (c *WUIComboBox) TypeName() string { return "combobox" }
func (c *WUIComboBox) String() string   { return "<combobox>" }
func (c *WUIComboBox) Copy() tender.Object {
	return &WUIComboBox{Value: c.Value}
}



type WUISlider struct {
	tender.ObjectImpl
	Value *wui.Slider
}

func (s *WUISlider) TypeName() string { return "slider" }
func (s *WUISlider) String() string   { return "<slider>" }
func (s *WUISlider) Copy() tender.Object {
	return &WUISlider{Value: s.Value}
}

type WUIProgressBar struct {
	tender.ObjectImpl
	Value *wui.ProgressBar
}

func (p *WUIProgressBar) TypeName() string { return "progressbar" }
func (p *WUIProgressBar) String() string   { return "<progressbar>" }
func (p *WUIProgressBar) Copy() tender.Object {
	return &WUIProgressBar{Value: p.Value}
}

type WUIRadioButton struct {
	tender.ObjectImpl
	Value *wui.RadioButton
}

func (r *WUIRadioButton) TypeName() string { return "radiobutton" }
func (r *WUIRadioButton) String() string   { return "<radiobutton>" }
func (r *WUIRadioButton) Copy() tender.Object {
	return &WUIRadioButton{Value: r.Value}
}


type WUIIntUpDown struct {
	tender.ObjectImpl
	Value *wui.IntUpDown
}

func (i *WUIIntUpDown) TypeName() string { return "intupdown" }
func (i *WUIIntUpDown) String() string   { return "<intupdown>" }
func (i *WUIIntUpDown) Copy() tender.Object {
	return &WUIIntUpDown{Value: i.Value}
}

type WUIFloatUpDown struct {
	tender.ObjectImpl
	Value *wui.FloatUpDown
}

func (f *WUIFloatUpDown) TypeName() string { return "floatupdown" }
func (f *WUIFloatUpDown) String() string   { return "<floatupdown>" }
func (f *WUIFloatUpDown) Copy() tender.Object {
	return &WUIFloatUpDown{Value: f.Value}
}

type WUIPanel struct {
	tender.ObjectImpl
	Value *wui.Panel
}

func (p *WUIPanel) TypeName() string { return "panel" }
func (p *WUIPanel) String() string   { return "<panel>" }
func (p *WUIPanel) Copy() tender.Object {
	return &WUIPanel{Value: p.Value}
}

type WUIPaintBox struct {
	tender.ObjectImpl
	Value *wui.PaintBox
}

func (p *WUIPaintBox) TypeName() string { return "paintbox" }
func (p *WUIPaintBox) String() string   { return "<paintbox>" }
func (p *WUIPaintBox) Copy() tender.Object {
	return &WUIPaintBox{Value: p.Value}
}


type WUIStringList struct {
	tender.ObjectImpl
	Value *wui.StringList
}

func (s *WUIStringList) TypeName() string { return "stringlist" }
func (s *WUIStringList) String() string   { return "<stringlist>" }
func (s *WUIStringList) Copy() tender.Object {
	return &WUIStringList{Value: s.Value}
}

type WUIStringTable struct {
	tender.ObjectImpl
	Value *wui.StringTable
}

func (s *WUIStringTable) TypeName() string { return "stringtable" }
func (s *WUIStringTable) String() string   { return "<stringtable>" }
func (s *WUIStringTable) Copy() tender.Object {
	return &WUIStringTable{Value: s.Value}
}



func wuiNewButton(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 0 {
		err = tender.ErrWrongNumArguments
		return
	}

	button := wui.NewButton()
	return &WUIButton{Value: button}, nil
}

func wuiNewCheckBox(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 0 {
		err = tender.ErrWrongNumArguments
		return
	}

	checkbox := wui.NewCheckBox()
	return &WUICheckBox{Value: checkbox}, nil
}

func wuiNewLabel(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 0 {
		err = tender.ErrWrongNumArguments
		return
	}

	label := wui.NewLabel()
	return &WUILabel{Value: label}, nil
}

func wuiNewEditLine(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 0 {
		err = tender.ErrWrongNumArguments
		return
	}

	editLine := wui.NewEditLine()
	return &WUIEditLine{Value: editLine}, nil
}

func wuiNewTextEdit(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 0 {
		err = tender.ErrWrongNumArguments
		return
	}

	textEdit := wui.NewTextEdit()
	return &WUITextEdit{Value: textEdit}, nil
}

func wuiNewComboBox(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 0 {
		err = tender.ErrWrongNumArguments
		return
	}

	comboBox := wui.NewComboBox()
	return &WUIComboBox{Value: comboBox}, nil
}

func wuiNewStringList(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 0 {
		err = tender.ErrWrongNumArguments
		return
	}

	stringList := wui.NewStringList()
	return &WUIStringList{Value: stringList}, nil
}

func wuiNewStringTable(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) < 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	header1, ok := tender.ToString(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "first header",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	headers := []string{header1}
	for i := 1; i < len(args); i++ {
		header, ok := tender.ToString(args[i])
		if !ok {
			err = tender.ErrInvalidArgumentType{
				Name:     fmt.Sprintf("header %d", i+1),
				Expected: "string(compatible)",
				Found:    args[i].TypeName(),
			}
			return
		}
		headers = append(headers, header)
	}

	stringTable := wui.NewStringTable(header1, headers[1:]...)
	return &WUIStringTable{Value: stringTable}, nil
}

func wuiNewSlider(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 0 {
		err = tender.ErrWrongNumArguments
		return
	}

	slider := wui.NewSlider()
	return &WUISlider{Value: slider}, nil
}

func wuiNewProgressBar(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 0 {
		err = tender.ErrWrongNumArguments
		return
	}

	progressBar := wui.NewProgressBar()
	return &WUIProgressBar{Value: progressBar}, nil
}

func wuiNewRadioButton(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 0 {
		err = tender.ErrWrongNumArguments
		return
	}

	radioButton := wui.NewRadioButton()
	return &WUIRadioButton{Value: radioButton}, nil
}

func wuiNewIntUpDown(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 0 {
		err = tender.ErrWrongNumArguments
		return
	}

	intUpDown := wui.NewIntUpDown()
	return &WUIIntUpDown{Value: intUpDown}, nil
}

func wuiNewFloatUpDown(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 0 {
		err = tender.ErrWrongNumArguments
		return
	}

	floatUpDown := wui.NewFloatUpDown()
	return &WUIFloatUpDown{Value: floatUpDown}, nil
}

func wuiNewPanel(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 0 {
		err = tender.ErrWrongNumArguments
		return
	}

	panel := wui.NewPanel()
	return &WUIPanel{Value: panel}, nil
}

func wuiNewPaintBox(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 0 {
		err = tender.ErrWrongNumArguments
		return
	}

	paintBox := wui.NewPaintBox()
	return &WUIPaintBox{Value: paintBox}, nil
}


// Helper function to extract control from wrapper
func extractControl(obj tender.Object) (wui.Control, bool) {
	switch c := obj.(type) {
	case *WUIButton:
		return c.Value, true
	case *WUICheckBox:
		return c.Value, true
	case *WUILabel:
		return c.Value, true
	case *WUIEditLine:
		return c.Value, true
	case *WUITextEdit:
		return c.Value, true
	case *WUIComboBox:
		return c.Value, true
	case *WUIStringList:
		return c.Value, true
	case *WUIStringTable:
		return c.Value, true
	case *WUISlider:
		return c.Value, true
	case *WUIProgressBar:
		return c.Value, true
	case *WUIRadioButton:
		return c.Value, true
	case *WUIIntUpDown:
		return c.Value, true
	case *WUIFloatUpDown:
		return c.Value, true
	case *WUIPanel:
		return c.Value, true
	case *WUIPaintBox:
		return c.Value, true
	default:
		return nil, false
	}
}









func (b *WUIButton) IndexGet(index tender.Object) (res tender.Object, err error) {
	strIdx, ok := index.(*tender.String)
	if !ok {
		return nil, tender.ErrInvalidIndexType
	}

	switch strIdx.Value {
	case "set_text":
		res = &tender.BuiltinFunction{
			Value: FuncASR(b.Value.SetText),
		}
	case "text":
		res = &tender.BuiltinFunction{
			Value: FuncARS(b.Value.Text),
		}
	case "set_font":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				font, ok := extractFont(args[0])
				if !ok {
					return nil, tender.ErrInvalidArgumentType{
						Name:     "font",
						Expected: "font",
						Found:    args[0].TypeName(),
					}
				}
				b.Value.SetFont(font)
				return tender.NullValue, nil
			},
		}
	case "font":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				font := b.Value.Font()
				return &WUIFont{Value: font}, nil
			},
		}
	case "focus":
		res = &tender.BuiltinFunction{
			Value: FuncAR(b.Value.Focus),
		}
	case "has_focus":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				if b.Value.HasFocus() {
					return tender.TrueValue, nil
				}
				return tender.FalseValue, nil
			},
		}
	case "set_onclick":
		res = &tender.BuiltinFunction{
			NeedVMObj: true,
			Value: func(args ...tender.Object) (tender.Object, error) {
				vm := args[0].(*tender.VMObj).Value
				args = args[1:] // the first arg is VMObj inserted by VM
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				b.Value.SetOnClick(func(){
					tender.WrapFuncCall(vm, args[0])
				})
				return tender.NullValue, nil
			},
		}
	case "onclick":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				// Return nil as we can't return a Go function to script
				return tender.NullValue, nil
			},
		}
	case "set_on_tab_focus":
		res = &tender.BuiltinFunction{
			NeedVMObj: true,
			Value: func(args ...tender.Object) (tender.Object, error) {
				vm := args[0].(*tender.VMObj).Value
				args = args[1:] // the first arg is VMObj inserted by VM
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				b.Value.SetOnTabFocus(func(){
					tender.WrapFuncCall(vm, args[0])
				})
				return tender.NullValue, nil
			},
		}
	case "on_tab_focus":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				// Return nil as we can't return a Go function to script
				return tender.NullValue, nil
			},
		}
	case "set_bounds":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 4 {
					return nil, tender.ErrWrongNumArguments
				}
				x, _ := tender.ToInt(args[0])
				y, _ := tender.ToInt(args[1])
				width, _ := tender.ToInt(args[2])
				height, _ := tender.ToInt(args[3])
				b.Value.SetBounds(x, y, width, height)
				return tender.NullValue, nil
			},
		}
	case "set_enabled":
		res = &tender.BuiltinFunction{
			Value: FuncABR(b.Value.SetEnabled),
		}
	case "enabled":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				if b.Value.Enabled() {
					return tender.TrueValue, nil
				}
				return tender.FalseValue, nil
			},
		}
	case "set_visible":
		res = &tender.BuiltinFunction{
			Value: FuncABR(b.Value.SetVisible),
		}
	case "visible":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				if b.Value.Visible() {
					return tender.TrueValue, nil
				}
				return tender.FalseValue, nil
			},
		}
	}
	return
}

// WUICheckBox updates
func (c *WUICheckBox) IndexGet(index tender.Object) (res tender.Object, err error) {
	strIdx, ok := index.(*tender.String)
	if !ok {
		return nil, tender.ErrInvalidIndexType
	}

	switch strIdx.Value {
	case "set_text":
		res = &tender.BuiltinFunction{
			Value: FuncASR(c.Value.SetText),
		}
	case "text":
		res = &tender.BuiltinFunction{
			Value: FuncARS(c.Value.Text),
		}
	case "set_checked":
		res = &tender.BuiltinFunction{
			Value: FuncABR(c.Value.SetChecked),
		}
	case "checked":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				if c.Value.Checked() {
					return tender.TrueValue, nil
				}
				return tender.FalseValue, nil
			},
		}
	case "set_font":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				font, ok := extractFont(args[0])
				if !ok {
					return nil, tender.ErrInvalidArgumentType{
						Name:     "font",
						Expected: "font",
						Found:    args[0].TypeName(),
					}
				}
				c.Value.SetFont(font)
				return tender.NullValue, nil
			},
		}
	case "font":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				font := c.Value.Font()
				return &WUIFont{Value: font}, nil
			},
		}
	case "focus":
		res = &tender.BuiltinFunction{
			Value: FuncAR(c.Value.Focus),
		}
	case "has_focus":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				if c.Value.HasFocus() {
					return tender.TrueValue, nil
				}
				return tender.FalseValue, nil
			},
		}
	case "set_on_change":
		res = &tender.BuiltinFunction{
			NeedVMObj: true,
			Value: func(args ...tender.Object) (tender.Object, error) {
				vm := args[0].(*tender.VMObj).Value
				args = args[1:]
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				c.Value.SetOnChange(func(checked bool){
					var arg tender.Object
					if checked {
						arg = tender.TrueValue
					} else {
						arg = tender.FalseValue
					}
					tender.WrapFuncCall(vm, args[0], arg)
				})
				return tender.NullValue, nil
			},
		}
	case "set_on_tab_focus":
		res = &tender.BuiltinFunction{
			NeedVMObj: true,
			Value: func(args ...tender.Object) (tender.Object, error) {
				vm := args[0].(*tender.VMObj).Value
				args = args[1:]
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				c.Value.SetOnTabFocus(func(){
					tender.WrapFuncCall(vm, args[0])
				})
				return tender.NullValue, nil
			},
		}
	case "on_tab_focus":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				return tender.NullValue, nil
			},
		}
	case "set_bounds":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 4 {
					return nil, tender.ErrWrongNumArguments
				}
				x, _ := tender.ToInt(args[0])
				y, _ := tender.ToInt(args[1])
				width, _ := tender.ToInt(args[2])
				height, _ := tender.ToInt(args[3])
				c.Value.SetBounds(x, y, width, height)
				return tender.NullValue, nil
			},
		}
	}
	return
}

// WUIComboBox updates
func (c *WUIComboBox) IndexGet(index tender.Object) (res tender.Object, err error) {
	strIdx, ok := index.(*tender.String)
	if !ok {
		return nil, tender.ErrInvalidIndexType
	}

	switch strIdx.Value {
	case "set_text":
		res = &tender.BuiltinFunction{
			Value: FuncASR(c.Value.SetText),
		}
	case "text":
		res = &tender.BuiltinFunction{
			Value: FuncARS(c.Value.Text),
		}
	case "set_items":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				arr, ok := args[0].(*tender.Array)
				if !ok {
					return nil, tender.ErrInvalidArgumentType{
						Name:     "items",
						Expected: "array",
						Found:    args[0].TypeName(),
					}
				}
				items := make([]string, len(arr.Value))
				for i, item := range arr.Value {
					items[i], _ = tender.ToString(item)
				}
				c.Value.SetItems(items)
				return tender.NullValue, nil
			},
		}
	case "add_item":
		res = &tender.BuiltinFunction{
			Value: FuncASR(c.Value.AddItem),
		}
	case "clear":
		res = &tender.BuiltinFunction{
			Value: FuncAR(c.Value.Clear),
		}
	case "set_selected_index":
		res = &tender.BuiltinFunction{
			Value: FuncAIR(c.Value.SetSelectedIndex),
		}
	case "selected_index":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				return &tender.Int{Value: int64(c.Value.SelectedIndex())}, nil
			},
		}
	case "items":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				items := c.Value.Items()
				tenderItems := make([]tender.Object, len(items))
				for i, item := range items {
					tenderItems[i] = &tender.String{Value: item}
				}
				return &tender.Array{Value: tenderItems}, nil
			},
		}
	case "focus":
		res = &tender.BuiltinFunction{
			Value: FuncAR(c.Value.Focus),
		}
	case "has_focus":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				if c.Value.HasFocus() {
					return tender.TrueValue, nil
				}
				return tender.FalseValue, nil
			},
		}
	case "set_font":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				font, ok := extractFont(args[0])
				if !ok {
					return nil, tender.ErrInvalidArgumentType{
						Name:     "font",
						Expected: "font",
						Found:    args[0].TypeName(),
					}
				}
				c.Value.SetFont(font)
				return tender.NullValue, nil
			},
		}
	case "font":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				font := c.Value.Font()
				return &WUIFont{Value: font}, nil
			},
		}
	case "set_on_change":
		res = &tender.BuiltinFunction{
			NeedVMObj: true,
			Value: func(args ...tender.Object) (tender.Object, error) {
				vm := args[0].(*tender.VMObj).Value
				args = args[1:]
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				c.Value.SetOnChange(func(newIndex int){
					tender.WrapFuncCall(vm, args[0], &tender.Int{Value: int64(newIndex)})
				})
				return tender.NullValue, nil
			},
		}
	case "set_on_tab_focus":
		res = &tender.BuiltinFunction{
			NeedVMObj: true,
			Value: func(args ...tender.Object) (tender.Object, error) {
				vm := args[0].(*tender.VMObj).Value
				args = args[1:]
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				c.Value.SetOnTabFocus(func(){
					tender.WrapFuncCall(vm, args[0])
				})
				return tender.NullValue, nil
			},
		}
	case "on_tab_focus":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				return tender.NullValue, nil
			},
		}
	case "set_bounds":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 4 {
					return nil, tender.ErrWrongNumArguments
				}
				x, _ := tender.ToInt(args[0])
				y, _ := tender.ToInt(args[1])
				width, _ := tender.ToInt(args[2])
				height, _ := tender.ToInt(args[3])
				c.Value.SetBounds(x, y, width, height)
				return tender.NullValue, nil
			},
		}
	}
	return
}

// WUIEditLine updates
func (e *WUIEditLine) IndexGet(index tender.Object) (res tender.Object, err error) {
	strIdx, ok := index.(*tender.String)
	if !ok {
		return nil, tender.ErrInvalidIndexType
	}

	switch strIdx.Value {
	case "set_text":
		res = &tender.BuiltinFunction{
			Value: FuncASR(e.Value.SetText),
		}
	case "text":
		res = &tender.BuiltinFunction{
			Value: FuncARS(e.Value.Text),
		}
	case "set_bounds":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 4 {
					return nil, tender.ErrWrongNumArguments
				}
				x, _ := tender.ToInt(args[0])
				y, _ := tender.ToInt(args[1])
				width, _ := tender.ToInt(args[2])
				height, _ := tender.ToInt(args[3])
				e.Value.SetBounds(x, y, width, height)
				return tender.NullValue, nil
			},
		}
	case "set_font":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				font, ok := extractFont(args[0])
				if !ok {
					return nil, tender.ErrInvalidArgumentType{
						Name:     "font",
						Expected: "font",
						Found:    args[0].TypeName(),
					}
				}
				e.Value.SetFont(font)
				return tender.NullValue, nil
			},
		}
	case "character_limit":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				return &tender.Int{Value: int64(e.Value.CharacterLimit())}, nil
			},
		}
	case "set_character_limit":
		res = &tender.BuiltinFunction{
			Value: FuncAIR(e.Value.SetCharacterLimit),
		}
	case "cursor_position":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				start, end := e.Value.CursorPosition()
				return &tender.Array{Value: []tender.Object{
					&tender.Int{Value: int64(start)},
					&tender.Int{Value: int64(end)},
				}}, nil
			},
		}
	case "set_cursor_position":
		res = &tender.BuiltinFunction{
			Value: FuncAIR(e.Value.SetCursorPosition),
		}
	case "is_password":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				if e.Value.IsPassword() {
					return tender.TrueValue, nil
				}
				return tender.FalseValue, nil
			},
		}
	case "set_is_password":
		res = &tender.BuiltinFunction{
			Value: FuncABR(e.Value.SetIsPassword),
		}
	case "read_only":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				if e.Value.ReadOnly() {
					return tender.TrueValue, nil
				}
				return tender.FalseValue, nil
			},
		}
	case "set_read_only":
		res = &tender.BuiltinFunction{
			Value: FuncABR(e.Value.SetReadOnly),
		}
	case "select_all":
		res = &tender.BuiltinFunction{
			Value: FuncAR(e.Value.SelectAll),
		}
	case "set_selection":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, tender.ErrWrongNumArguments
				}
				start, _ := tender.ToInt(args[0])
				end, _ := tender.ToInt(args[1])
				e.Value.SetSelection(start, end)
				return tender.NullValue, nil
			},
		}
	case "set_on_tab_focus":
		res = &tender.BuiltinFunction{
			NeedVMObj: true,
			Value: func(args ...tender.Object) (tender.Object, error) {
				vm := args[0].(*tender.VMObj).Value
				args = args[1:]
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				e.Value.SetOnTabFocus(func(){
					tender.WrapFuncCall(vm, args[0])
				})
				return tender.NullValue, nil
			},
		}
	case "set_on_text_change":
		res = &tender.BuiltinFunction{
			NeedVMObj: true,
			Value: func(args ...tender.Object) (tender.Object, error) {
				vm := args[0].(*tender.VMObj).Value
				args = args[1:]
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				e.Value.SetOnTextChange(func(){
					tender.WrapFuncCall(vm, args[0])
				})
				return tender.NullValue, nil
			},
		}
	case "on_tab_focus":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				return tender.NullValue, nil
			},
		}
	case "on_text_change":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				return tender.NullValue, nil
			},
		}
	}
	return
}

// WUIFloatUpDown updates
func (f *WUIFloatUpDown) IndexGet(index tender.Object) (res tender.Object, err error) {
	strIdx, ok := index.(*tender.String)
	if !ok {
		return nil, tender.ErrInvalidIndexType
	}

	switch strIdx.Value {
	case "set_min":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				min, _ := tender.ToFloat64(args[0])
				f.Value.SetMin(min)
				return tender.NullValue, nil
			},
		}
	case "set_max":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				max, _ := tender.ToFloat64(args[0])
				f.Value.SetMax(max)
				return tender.NullValue, nil
			},
		}
	case "set_value":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				val, _ := tender.ToFloat64(args[0])
				f.Value.SetValue(val)
				return tender.NullValue, nil
			},
		}
	case "value":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				return &tender.Float{Value: float64(f.Value.Value())}, nil
			},
		}
	case "min":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				return &tender.Float{Value: f.Value.Min()}, nil
			},
		}
	case "max":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				return &tender.Float{Value: f.Value.Max()}, nil
			},
		}
	case "min_max":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				min, max := f.Value.MinMax()
				return &tender.Array{Value: []tender.Object{
					&tender.Float{Value: min},
					&tender.Float{Value: max},
				}}, nil
			},
		}
	case "set_min_max":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, tender.ErrWrongNumArguments
				}
				min, _ := tender.ToFloat64(args[0])
				max, _ := tender.ToFloat64(args[1])
				f.Value.SetMinMax(min, max)
				return tender.NullValue, nil
			},
		}
	case "precision":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				return &tender.Int{Value: int64(f.Value.Precision())}, nil
			},
		}
	case "set_precision":
		res = &tender.BuiltinFunction{
			Value: FuncAIR(f.Value.SetPrecision),
		}
	case "cursor_position":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				start, end := f.Value.CursorPosition()
				return &tender.Array{Value: []tender.Object{
					&tender.Int{Value: int64(start)},
					&tender.Int{Value: int64(end)},
				}}, nil
			},
		}
	case "set_cursor_position":
		res = &tender.BuiltinFunction{
			Value: FuncAIR(f.Value.SetCursorPosition),
		}
	case "select_all":
		res = &tender.BuiltinFunction{
			Value: FuncAR(f.Value.SelectAll),
		}
	case "set_selection":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, tender.ErrWrongNumArguments
				}
				start, _ := tender.ToInt(args[0])
				end, _ := tender.ToInt(args[1])
				f.Value.SetSelection(start, end)
				return tender.NullValue, nil
			},
		}
	case "set_on_tab_focus":
		res = &tender.BuiltinFunction{
			NeedVMObj: true,
			Value: func(args ...tender.Object) (tender.Object, error) {
				vm := args[0].(*tender.VMObj).Value
				args = args[1:]
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				f.Value.SetOnTabFocus(func(){
					tender.WrapFuncCall(vm, args[0])
				})
				return tender.NullValue, nil
			},
		}
	case "set_on_value_change":
		res = &tender.BuiltinFunction{
			NeedVMObj: true,
			Value: func(args ...tender.Object) (tender.Object, error) {
				vm := args[0].(*tender.VMObj).Value
				args = args[1:]
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				f.Value.SetOnValueChange(func(value float64){
					tender.WrapFuncCall(vm, args[0], &tender.Float{Value: value})
				})
				return tender.NullValue, nil
			},
		}
	case "on_tab_focus":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				return tender.NullValue, nil
			},
		}
	case "on_value_change":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				return tender.NullValue, nil
			},
		}
	case "set_bounds":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 4 {
					return nil, tender.ErrWrongNumArguments
				}
				x, _ := tender.ToInt(args[0])
				y, _ := tender.ToInt(args[1])
				width, _ := tender.ToInt(args[2])
				height, _ := tender.ToInt(args[3])
				f.Value.SetBounds(x, y, width, height)
				return tender.NullValue, nil
			},
		}
	// Position and size methods
	case "set_pos":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, tender.ErrWrongNumArguments
				}
				x, _ := tender.ToInt(args[0])
				y, _ := tender.ToInt(args[1])
				f.Value.SetPos(x, y)
				return tender.NullValue, nil
			},
		}
	case "set_size":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, tender.ErrWrongNumArguments
				}
				width, _ := tender.ToInt(args[0])
				height, _ := tender.ToInt(args[1])
				f.Value.SetSize(width, height)
				return tender.NullValue, nil
			},
		}
	case "set_x":
		res = &tender.BuiltinFunction{
			Value: FuncAIR(f.Value.SetX),
		}
	case "set_y":
		res = &tender.BuiltinFunction{
			Value: FuncAIR(f.Value.SetY),
		}
	case "set_width":
		res = &tender.BuiltinFunction{
			Value: FuncAIR(f.Value.SetWidth),
		}
	case "set_height":
		res = &tender.BuiltinFunction{
			Value: FuncAIR(f.Value.SetHeight),
		}
	}
	return
}

// WUIIntUpDown updates - similar to FloatUpDown but with int values
func (i *WUIIntUpDown) IndexGet(index tender.Object) (res tender.Object, err error) {
	strIdx, ok := index.(*tender.String)
	if !ok {
		return nil, tender.ErrInvalidIndexType
	}

	switch strIdx.Value {
	case "set_min":
		res = &tender.BuiltinFunction{
			Value: FuncAIR(i.Value.SetMin),
		}
	case "set_max":
		res = &tender.BuiltinFunction{
			Value: FuncAIR(i.Value.SetMax),
		}
	case "set_value":
		res = &tender.BuiltinFunction{
			Value: FuncAIR(i.Value.SetValue),
		}
	case "value":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				return &tender.Int{Value: int64(i.Value.Value())}, nil
			},
		}
	case "min":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				return &tender.Int{Value: int64(i.Value.Min())}, nil
			},
		}
	case "max":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				return &tender.Int{Value: int64(i.Value.Max())}, nil
			},
		}
	case "min_max":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				min, max := i.Value.MinMax()
				return &tender.Array{Value: []tender.Object{
					&tender.Int{Value: int64(min)},
					&tender.Int{Value: int64(max)},
				}}, nil
			},
		}
	case "set_min_max":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, tender.ErrWrongNumArguments
				}
				min, _ := tender.ToInt(args[0])
				max, _ := tender.ToInt(args[1])
				i.Value.SetMinMax(min, max)
				return tender.NullValue, nil
			},
		}
	case "cursor_position":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				start, end := i.Value.CursorPosition()
				return &tender.Array{Value: []tender.Object{
					&tender.Int{Value: int64(start)},
					&tender.Int{Value: int64(end)},
				}}, nil
			},
		}
	case "set_cursor_position":
		res = &tender.BuiltinFunction{
			Value: FuncAIR(i.Value.SetCursorPosition),
		}
	case "select_all":
		res = &tender.BuiltinFunction{
			Value: FuncAR(i.Value.SelectAll),
		}
	case "set_selection":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, tender.ErrWrongNumArguments
				}
				start, _ := tender.ToInt(args[0])
				end, _ := tender.ToInt(args[1])
				i.Value.SetSelection(start, end)
				return tender.NullValue, nil
			},
		}
	case "set_on_tab_focus":
		res = &tender.BuiltinFunction{
			NeedVMObj: true,
			Value: func(args ...tender.Object) (tender.Object, error) {
				vm := args[0].(*tender.VMObj).Value
				args = args[1:]
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				i.Value.SetOnTabFocus(func(){
					tender.WrapFuncCall(vm, args[0])
				})
				return tender.NullValue, nil
			},
		}
	case "set_on_value_change":
		res = &tender.BuiltinFunction{
			NeedVMObj: true,
			Value: func(args ...tender.Object) (tender.Object, error) {
				vm := args[0].(*tender.VMObj).Value
				args = args[1:]
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				i.Value.SetOnValueChange(func(value int){
					tender.WrapFuncCall(vm, args[0], &tender.Int{Value: int64(value)})
				})
				return tender.NullValue, nil
			},
		}
	case "on_tab_focus":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				return tender.NullValue, nil
			},
		}
	case "on_value_change":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				return tender.NullValue, nil
			},
		}
	case "set_bounds":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 4 {
					return nil, tender.ErrWrongNumArguments
				}
				x, _ := tender.ToInt(args[0])
				y, _ := tender.ToInt(args[1])
				width, _ := tender.ToInt(args[2])
				height, _ := tender.ToInt(args[3])
				i.Value.SetBounds(x, y, width, height)
				return tender.NullValue, nil
			},
		}
	// Position and size methods
	case "set_pos":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, tender.ErrWrongNumArguments
				}
				x, _ := tender.ToInt(args[0])
				y, _ := tender.ToInt(args[1])
				i.Value.SetPos(x, y)
				return tender.NullValue, nil
			},
		}
	case "set_size":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, tender.ErrWrongNumArguments
				}
				width, _ := tender.ToInt(args[0])
				height, _ := tender.ToInt(args[1])
				i.Value.SetSize(width, height)
				return tender.NullValue, nil
			},
		}
	case "set_x":
		res = &tender.BuiltinFunction{
			Value: FuncAIR(i.Value.SetX),
		}
	case "set_y":
		res = &tender.BuiltinFunction{
			Value: FuncAIR(i.Value.SetY),
		}
	case "set_width":
		res = &tender.BuiltinFunction{
			Value: FuncAIR(i.Value.SetWidth),
		}
	case "set_height":
		res = &tender.BuiltinFunction{
			Value: FuncAIR(i.Value.SetHeight),
		}
	case "set_visible":
		res = &tender.BuiltinFunction{
			Value: FuncABR(i.Value.SetVisible),
		}
	case "visible":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				if i.Value.Visible() {
					return tender.TrueValue, nil
				}
				return tender.FalseValue, nil
			},
		}
	}
	return
}

// WUILabel updates
func (l *WUILabel) IndexGet(index tender.Object) (res tender.Object, err error) {
	strIdx, ok := index.(*tender.String)
	if !ok {
		return nil, tender.ErrInvalidIndexType
	}

	switch strIdx.Value {
	case "set_text":
		res = &tender.BuiltinFunction{
			Value: FuncASR(l.Value.SetText),
		}
	case "text":
		res = &tender.BuiltinFunction{
			Value: FuncARS(l.Value.Text),
		}
	case "set_alignment":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				alignment, _ := tender.ToInt(args[0])
				l.Value.SetAlignment(wui.TextAlignment(alignment))
				return tender.NullValue, nil
			},
		}
	case "alignment":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				return &tender.Int{Value: int64(l.Value.Alignment())}, nil
			},
		}
	case "set_font":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				font, ok := extractFont(args[0])
				if !ok {
					return nil, tender.ErrInvalidArgumentType{
						Name:     "font",
						Expected: "font",
						Found:    args[0].TypeName(),
					}
				}
				l.Value.SetFont(font)
				return tender.NullValue, nil
			},
		}
	case "font":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				font := l.Value.Font()
				return &WUIFont{Value: font}, nil
			},
		}
	case "focus":
		res = &tender.BuiltinFunction{
			Value: FuncAR(l.Value.Focus),
		}
	case "has_focus":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				if l.Value.HasFocus() {
					return tender.TrueValue, nil
				}
				return tender.FalseValue, nil
			},
		}
	case "set_bounds":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 4 {
					return nil, tender.ErrWrongNumArguments
				}
				x, _ := tender.ToInt(args[0])
				y, _ := tender.ToInt(args[1])
				width, _ := tender.ToInt(args[2])
				height, _ := tender.ToInt(args[3])
				l.Value.SetBounds(x, y, width, height)
				return tender.NullValue, nil
			},
		}
	}
	return
}

// WUIPaintBox updates - needs WUIFont and WUICanvas wrapper types
func (p *WUIPaintBox) IndexGet(index tender.Object) (res tender.Object, err error) {
	strIdx, ok := index.(*tender.String)
	if !ok {
		return nil, tender.ErrInvalidIndexType
	}

	switch strIdx.Value {
	case "set_bounds":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 4 {
					return nil, tender.ErrWrongNumArguments
				}
				x, _ := tender.ToInt(args[0])
				y, _ := tender.ToInt(args[1])
				width, _ := tender.ToInt(args[2])
				height, _ := tender.ToInt(args[3])
				p.Value.SetBounds(x, y, width, height)
				return tender.NullValue, nil
			},
		}
	case "anchors":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				horizontal, vertical := p.Value.Anchors()
				return &tender.Array{Value: []tender.Object{
					&tender.Int{Value: int64(horizontal)},
					&tender.Int{Value: int64(vertical)},
				}}, nil
			},
		}
	case "bounds":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				x, y, width, height := p.Value.Bounds()
				return &tender.Array{Value: []tender.Object{
					&tender.Int{Value: int64(x)},
					&tender.Int{Value: int64(y)},
					&tender.Int{Value: int64(width)},
					&tender.Int{Value: int64(height)},
				}}, nil
			},
		}
	case "enabled":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				if p.Value.Enabled() {
					return tender.TrueValue, nil
				}
				return tender.FalseValue, nil
			},
		}
	case "set_enabled":
		res = &tender.BuiltinFunction{
			Value: FuncABR(p.Value.SetEnabled),
		}
	case "handle":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				return &tender.Int{Value: int64(p.Value.Handle())}, nil
			},
		}
	case "height":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				return &tender.Int{Value: int64(p.Value.Height())}, nil
			},
		}
	case "horizontal_anchor":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				return &tender.Int{Value: int64(p.Value.HorizontalAnchor())}, nil
			},
		}
	case "set_on_mouse_move":
		res = &tender.BuiltinFunction{
			NeedVMObj: true,
			Value: func(args ...tender.Object) (tender.Object, error) {
				vm := args[0].(*tender.VMObj).Value
				args = args[1:]
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				p.Value.SetOnMouseMove(func(x, y int){
					tender.WrapFuncCall(vm, args[0], 
						&tender.Int{Value: int64(x)},
						&tender.Int{Value: int64(y)})
				})
				return tender.NullValue, nil
			},
		}
	case "set_on_paint":
		res = &tender.BuiltinFunction{
			NeedVMObj: true,
			Value: func(args ...tender.Object) (tender.Object, error) {
				vm := args[0].(*tender.VMObj).Value
				args = args[1:]
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				p.Value.SetOnPaint(func(canvas *wui.Canvas){
					// You'll need to create a WUICanvas wrapper type
					wrapper := &WUICanvas{Value: canvas}
					tender.WrapFuncCall(vm, args[0], wrapper)
				})
				return tender.NullValue, nil
			},
		}
	case "set_on_resize":
		res = &tender.BuiltinFunction{
			NeedVMObj: true,
			Value: func(args ...tender.Object) (tender.Object, error) {
				vm := args[0].(*tender.VMObj).Value
				args = args[1:]
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				p.Value.SetOnResize(func(){
					tender.WrapFuncCall(vm, args[0])
				})
				return tender.NullValue, nil
			},
		}
	case "on_resize":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				return tender.NullValue, nil
			},
		}
	case "paint":
		res = &tender.BuiltinFunction{
			Value: FuncAR(p.Value.Paint),
		}
	case "parent":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				parent := p.Value.Parent()
				if parent == nil {
					return tender.NullValue, nil
				}
				// Need to wrap parent container - this is complex
				return tender.NullValue, nil
			},
		}
	case "position":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				x, y := p.Value.Position()
				return &tender.Array{Value: []tender.Object{
					&tender.Int{Value: int64(x)},
					&tender.Int{Value: int64(y)},
				}}, nil
			},
		}
	case "set_anchors":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, tender.ErrWrongNumArguments
				}
				horizontal, _ := tender.ToInt(args[0])
				vertical, _ := tender.ToInt(args[1])
				p.Value.SetAnchors(wui.Anchor(horizontal), wui.Anchor(vertical))
				return tender.NullValue, nil
			},
		}
	case "set_height":
		res = &tender.BuiltinFunction{
			Value: FuncAIR(p.Value.SetHeight),
		}
	case "set_horizontal_anchor":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				anchor, _ := tender.ToInt(args[0])
				p.Value.SetHorizontalAnchor(wui.Anchor(anchor))
				return tender.NullValue, nil
			},
		}
	case "set_position":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, tender.ErrWrongNumArguments
				}
				x, _ := tender.ToInt(args[0])
				y, _ := tender.ToInt(args[1])
				p.Value.SetPosition(x, y)
				return tender.NullValue, nil
			},
		}
	case "set_size":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, tender.ErrWrongNumArguments
				}
				width, _ := tender.ToInt(args[0])
				height, _ := tender.ToInt(args[1])
				p.Value.SetSize(width, height)
				return tender.NullValue, nil
			},
		}
	case "set_vertical_anchor":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				anchor, _ := tender.ToInt(args[0])
				p.Value.SetVerticalAnchor(wui.Anchor(anchor))
				return tender.NullValue, nil
			},
		}
	case "set_visible":
		res = &tender.BuiltinFunction{
			Value: FuncABR(p.Value.SetVisible),
		}
	case "set_width":
		res = &tender.BuiltinFunction{
			Value: FuncAIR(p.Value.SetWidth),
		}
	case "set_x":
		res = &tender.BuiltinFunction{
			Value: FuncAIR(p.Value.SetX),
		}
	case "set_y":
		res = &tender.BuiltinFunction{
			Value: FuncAIR(p.Value.SetY),
		}
	case "size":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				width, height := p.Value.Size()
				return &tender.Array{Value: []tender.Object{
					&tender.Int{Value: int64(width)},
					&tender.Int{Value: int64(height)},
				}}, nil
			},
		}
	case "vertical_anchor":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				return &tender.Int{Value: int64(p.Value.VerticalAnchor())}, nil
			},
		}
	case "visible":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				if p.Value.Visible() {
					return tender.TrueValue, nil
				}
				return tender.FalseValue, nil
			},
		}
	case "width":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				return &tender.Int{Value: int64(p.Value.Width())}, nil
			},
		}
	case "x":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				return &tender.Int{Value: int64(p.Value.X())}, nil
			},
		}
	case "y":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				return &tender.Int{Value: int64(p.Value.Y())}, nil
			},
		}
	}
	return
}



type WUICanvas struct {
	tender.ObjectImpl
	Value *wui.Canvas
}

func (c *WUICanvas) TypeName() string { return "canvas" }
func (c *WUICanvas) String() string   { return "<canvas>" }
func (c *WUICanvas) Copy() tender.Object {
	return &WUICanvas{Value: c.Value}
}

func (c *WUICanvas) IndexGet(index tender.Object) (res tender.Object, err error) {
    strIdx, ok := index.(*tender.String)
    if !ok {
        return nil, tender.ErrInvalidIndexType
    }

    switch strIdx.Value {
    case "arc":
        res = &tender.BuiltinFunction{
            Value: func(args ...tender.Object) (tender.Object, error) {
                if len(args) != 7 {
                    return nil, tender.ErrWrongNumArguments
                }
                x, _ := tender.ToInt(args[0])
                y, _ := tender.ToInt(args[1])
                width, _ := tender.ToInt(args[2])
                height, _ := tender.ToInt(args[3])
                fromClockAngle, _ := tender.ToFloat64(args[4])
                dAngle, _ := tender.ToFloat64(args[5])
                color, ok := extractColor(args[6])
                if !ok {
                    return nil, tender.ErrInvalidArgumentType{
                        Name:     "color",
                        Expected: "color",
                        Found:    args[6].TypeName(),
                    }
                }
                c.Value.Arc(x, y, width, height, fromClockAngle, dAngle, color)
                return tender.NullValue, nil
            },
        }
    case "clear_draw_regions":
        res = &tender.BuiltinFunction{
            Value: FuncAR(c.Value.ClearDrawRegions),
        }
    case "draw_ellipse":
        res = &tender.BuiltinFunction{
            Value: func(args ...tender.Object) (tender.Object, error) {
                if len(args) != 5 {
                    return nil, tender.ErrWrongNumArguments
                }
                x, _ := tender.ToInt(args[0])
                y, _ := tender.ToInt(args[1])
                width, _ := tender.ToInt(args[2])
                height, _ := tender.ToInt(args[3])
                color, ok := extractColor(args[4])
                if !ok {
                    return nil, tender.ErrInvalidArgumentType{
                        Name:     "color",
                        Expected: "color",
                        Found:    args[4].TypeName(),
                    }
                }
                c.Value.DrawEllipse(x, y, width, height, color)
                return tender.NullValue, nil
            },
        }
    case "draw_image":
        res = &tender.BuiltinFunction{
            Value: func(args ...tender.Object) (tender.Object, error) {
                if len(args) != 4 {
                    return nil, tender.ErrWrongNumArguments
                }
                img, ok := args[0].(*WUIImage)
                if !ok {
                    return nil, tender.ErrInvalidArgumentType{
                        Name:     "image",
                        Expected: "image",
                        Found:    args[0].TypeName(),
                    }
                }
                rect, ok := args[1].(*WUIRectangle)
                if !ok {
                    return nil, tender.ErrInvalidArgumentType{
                        Name:     "rectangle",
                        Expected: "rectangle",
                        Found:    args[1].TypeName(),
                    }
                }
                destX, _ := tender.ToInt(args[2])
                destY, _ := tender.ToInt(args[3])
                c.Value.DrawImage(img.Value, rect.Value, destX, destY)
                return tender.NullValue, nil
            },
        }
    case "draw_pie":
        res = &tender.BuiltinFunction{
            Value: func(args ...tender.Object) (tender.Object, error) {
                if len(args) != 7 {
                    return nil, tender.ErrWrongNumArguments
                }
                x, _ := tender.ToInt(args[0])
                y, _ := tender.ToInt(args[1])
                width, _ := tender.ToInt(args[2])
                height, _ := tender.ToInt(args[3])
                fromClockAngle, _ := tender.ToFloat64(args[4])
                dAngle, _ := tender.ToFloat64(args[5])
                color, ok := extractColor(args[6])
                if !ok {
                    return nil, tender.ErrInvalidArgumentType{
                        Name:     "color",
                        Expected: "color",
                        Found:    args[6].TypeName(),
                    }
                }
                c.Value.DrawPie(x, y, width, height, fromClockAngle, dAngle, color)
                return tender.NullValue, nil
            },
        }
    case "draw_rect":
        res = &tender.BuiltinFunction{
            Value: func(args ...tender.Object) (tender.Object, error) {
                if len(args) != 5 {
                    return nil, tender.ErrWrongNumArguments
                }
                x, _ := tender.ToInt(args[0])
                y, _ := tender.ToInt(args[1])
                width, _ := tender.ToInt(args[2])
                height, _ := tender.ToInt(args[3])
                color, ok := extractColor(args[4])
                if !ok {
                    return nil, tender.ErrInvalidArgumentType{
                        Name:     "color",
                        Expected: "color",
                        Found:    args[4].TypeName(),
                    }
                }
                c.Value.DrawRect(x, y, width, height, color)
                return tender.NullValue, nil
            },
        }
    case "fill_ellipse":
        res = &tender.BuiltinFunction{
            Value: func(args ...tender.Object) (tender.Object, error) {
                if len(args) != 5 {
                    return nil, tender.ErrWrongNumArguments
                }
                x, _ := tender.ToInt(args[0])
                y, _ := tender.ToInt(args[1])
                width, _ := tender.ToInt(args[2])
                height, _ := tender.ToInt(args[3])
                color, ok := extractColor(args[4])
                if !ok {
                    return nil, tender.ErrInvalidArgumentType{
                        Name:     "color",
                        Expected: "color",
                        Found:    args[4].TypeName(),
                    }
                }
                c.Value.FillEllipse(x, y, width, height, color)
                return tender.NullValue, nil
            },
        }
    case "fill_pie":
        res = &tender.BuiltinFunction{
            Value: func(args ...tender.Object) (tender.Object, error) {
                if len(args) != 7 {
                    return nil, tender.ErrWrongNumArguments
                }
                x, _ := tender.ToInt(args[0])
                y, _ := tender.ToInt(args[1])
                width, _ := tender.ToInt(args[2])
                height, _ := tender.ToInt(args[3])
                fromClockAngle, _ := tender.ToFloat64(args[4])
                dAngle, _ := tender.ToFloat64(args[5])
                color, ok := extractColor(args[6])
                if !ok {
                    return nil, tender.ErrInvalidArgumentType{
                        Name:     "color",
                        Expected: "color",
                        Found:    args[6].TypeName(),
                    }
                }
                c.Value.FillPie(x, y, width, height, fromClockAngle, dAngle, color)
                return tender.NullValue, nil
            },
        }
    case "fill_rect":
        res = &tender.BuiltinFunction{
            Value: func(args ...tender.Object) (tender.Object, error) {
                if len(args) != 5 {
                    return nil, tender.ErrWrongNumArguments
                }
                x, _ := tender.ToInt(args[0])
                y, _ := tender.ToInt(args[1])
                width, _ := tender.ToInt(args[2])
                height, _ := tender.ToInt(args[3])
                color, ok := extractColor(args[4])
                if !ok {
                    return nil, tender.ErrInvalidArgumentType{
                        Name:     "color",
                        Expected: "color",
                        Found:    args[4].TypeName(),
                    }
                }
                c.Value.FillRect(x, y, width, height, color)
                return tender.NullValue, nil
            },
        }
    case "handle":
        res = &tender.BuiltinFunction{
            Value: func(args ...tender.Object) (tender.Object, error) {
                if len(args) != 0 {
                    return nil, tender.ErrWrongNumArguments
                }
                return &tender.Int{Value: int64(c.Value.Handle())}, nil
            },
        }
    case "height":
        res = &tender.BuiltinFunction{
            Value: func(args ...tender.Object) (tender.Object, error) {
                if len(args) != 0 {
                    return nil, tender.ErrWrongNumArguments
                }
                return &tender.Int{Value: int64(c.Value.Height())}, nil
            },
        }
    case "line":
        res = &tender.BuiltinFunction{
            Value: func(args ...tender.Object) (tender.Object, error) {
                if len(args) != 5 {
                    return nil, tender.ErrWrongNumArguments
                }
                x1, _ := tender.ToInt(args[0])
                y1, _ := tender.ToInt(args[1])
                x2, _ := tender.ToInt(args[2])
                y2, _ := tender.ToInt(args[3])
                color, ok := extractColor(args[4])
                if !ok {
                    return nil, tender.ErrInvalidArgumentType{
                        Name:     "color",
                        Expected: "color",
                        Found:    args[4].TypeName(),
                    }
                }
                c.Value.Line(x1, y1, x2, y2, color)
                return tender.NullValue, nil
            },
        }
    case "polygon":
        res = &tender.BuiltinFunction{
            Value: func(args ...tender.Object) (tender.Object, error) {
                if len(args) != 2 {
                    return nil, tender.ErrWrongNumArguments
                }
                arr, ok := args[0].(*tender.Array)
                if !ok {
                    return nil, tender.ErrInvalidArgumentType{
                        Name:     "points",
                        Expected: "array",
                        Found:    args[0].TypeName(),
                    }
                }
                points := make([]wui.Point, len(arr.Value))
                for i, obj := range arr.Value {
                    pointArr, ok := obj.(*tender.Array)
                    if !ok || len(pointArr.Value) != 2 {
                        return nil, tender.ErrInvalidArgumentType{
                            Name:     fmt.Sprintf("point %d", i),
                            Expected: "array of 2 ints",
                            Found:    obj.TypeName(),
                        }
                    }
                    x, _ := tender.ToInt32(pointArr.Value[0])
                    y, _ := tender.ToInt32(pointArr.Value[1])
                    points[i] = wui.Point{X: x, Y: y}
                }
                color, ok := extractColor(args[1])
                if !ok {
                    return nil, tender.ErrInvalidArgumentType{
                        Name:     "color",
                        Expected: "color",
                        Found:    args[1].TypeName(),
                    }
                }
                c.Value.Polygon(points, color)
                return tender.NullValue, nil
            },
        }
    case "polyline":
        res = &tender.BuiltinFunction{
            Value: func(args ...tender.Object) (tender.Object, error) {
                if len(args) != 2 {
                    return nil, tender.ErrWrongNumArguments
                }
                arr, ok := args[0].(*tender.Array)
                if !ok {
                    return nil, tender.ErrInvalidArgumentType{
                        Name:     "points",
                        Expected: "array",
                        Found:    args[0].TypeName(),
                    }
                }
                points := make([]wui.Point, len(arr.Value))
                for i, obj := range arr.Value {
                    pointArr, ok := obj.(*tender.Array)
                    if !ok || len(pointArr.Value) != 2 {
                        return nil, tender.ErrInvalidArgumentType{
                            Name:     fmt.Sprintf("point %d", i),
                            Expected: "array of 2 ints",
                            Found:    obj.TypeName(),
                        }
                    }
                    x, _ := tender.ToInt32(pointArr.Value[0])
                    y, _ := tender.ToInt32(pointArr.Value[1])
                    points[i] = wui.Point{X: x, Y: y}
                }
                color, ok := extractColor(args[1])
                if !ok {
                    return nil, tender.ErrInvalidArgumentType{
                        Name:     "color",
                        Expected: "color",
                        Found:    args[1].TypeName(),
                    }
                }
                c.Value.Polyline(points, color)
                return tender.NullValue, nil
            },
        }
    case "pop_draw_region":
        res = &tender.BuiltinFunction{
            Value: FuncAR(c.Value.PopDrawRegion),
        }
    case "push_draw_region":
        res = &tender.BuiltinFunction{
            Value: func(args ...tender.Object) (tender.Object, error) {
                if len(args) != 4 {
                    return nil, tender.ErrWrongNumArguments
                }
                x, _ := tender.ToInt(args[0])
                y, _ := tender.ToInt(args[1])
                width, _ := tender.ToInt(args[2])
                height, _ := tender.ToInt(args[3])
                c.Value.PushDrawRegion(x, y, width, height)
                return tender.NullValue, nil
            },
        }
    case "set_font":
        res = &tender.BuiltinFunction{
            Value: func(args ...tender.Object) (tender.Object, error) {
                if len(args) != 1 {
                    return nil, tender.ErrWrongNumArguments
                }
                font, ok := extractFont(args[0])
                if !ok {
                    return nil, tender.ErrInvalidArgumentType{
                        Name:     "font",
                        Expected: "font",
                        Found:    args[0].TypeName(),
                    }
                }
                c.Value.SetFont(font)
                return tender.NullValue, nil
            },
        }
    case "size":
        res = &tender.BuiltinFunction{
            Value: func(args ...tender.Object) (tender.Object, error) {
                if len(args) != 0 {
                    return nil, tender.ErrWrongNumArguments
                }
                width, height := c.Value.Size()
                return &tender.Array{Value: []tender.Object{
                    &tender.Int{Value: int64(width)},
                    &tender.Int{Value: int64(height)},
                }}, nil
            },
        }
    case "text_extent":
        res = &tender.BuiltinFunction{
            Value: func(args ...tender.Object) (tender.Object, error) {
                if len(args) != 1 {
                    return nil, tender.ErrWrongNumArguments
                }
                s, ok := tender.ToString(args[0])
                if !ok {
                    return nil, tender.ErrInvalidArgumentType{
                        Name:     "text",
                        Expected: "string",
                        Found:    args[0].TypeName(),
                    }
                }
                width, height := c.Value.TextExtent(s)
                return &tender.Array{Value: []tender.Object{
                    &tender.Int{Value: int64(width)},
                    &tender.Int{Value: int64(height)},
                }}, nil
            },
        }
    case "text_out":
        res = &tender.BuiltinFunction{
            Value: func(args ...tender.Object) (tender.Object, error) {
                if len(args) != 4 {
                    return nil, tender.ErrWrongNumArguments
                }
                x, _ := tender.ToInt(args[0])
                y, _ := tender.ToInt(args[1])
                s, ok := tender.ToString(args[2])
                if !ok {
                    return nil, tender.ErrInvalidArgumentType{
                        Name:     "text",
                        Expected: "string",
                        Found:    args[2].TypeName(),
                    }
                }
                color, ok := extractColor(args[3])
                if !ok {
                    return nil, tender.ErrInvalidArgumentType{
                        Name:     "color",
                        Expected: "color",
                        Found:    args[3].TypeName(),
                    }
                }
                c.Value.TextOut(x, y, s, color)
                return tender.NullValue, nil
            },
        }
    case "text_rect":
        res = &tender.BuiltinFunction{
            Value: func(args ...tender.Object) (tender.Object, error) {
                if len(args) != 6 {
                    return nil, tender.ErrWrongNumArguments
                }
                x, _ := tender.ToInt(args[0])
                y, _ := tender.ToInt(args[1])
                w, _ := tender.ToInt(args[2])
                h, _ := tender.ToInt(args[3])
                s, ok := tender.ToString(args[4])
                if !ok {
                    return nil, tender.ErrInvalidArgumentType{
                        Name:     "text",
                        Expected: "string",
                        Found:    args[4].TypeName(),
                    }
                }
                color, ok := extractColor(args[5])
                if !ok {
                    return nil, tender.ErrInvalidArgumentType{
                        Name:     "color",
                        Expected: "color",
                        Found:    args[5].TypeName(),
                    }
                }
                c.Value.TextRect(x, y, w, h, s, color)
                return tender.NullValue, nil
            },
        }
    case "text_rect_extent":
        res = &tender.BuiltinFunction{
            Value: func(args ...tender.Object) (tender.Object, error) {
                if len(args) != 2 {
                    return nil, tender.ErrWrongNumArguments
                }
                s, ok := tender.ToString(args[0])
                if !ok {
                    return nil, tender.ErrInvalidArgumentType{
                        Name:     "text",
                        Expected: "string",
                        Found:    args[0].TypeName(),
                    }
                }
                givenWidth, _ := tender.ToInt(args[1])
                width, height := c.Value.TextRectExtent(s, givenWidth)
                return &tender.Array{Value: []tender.Object{
                    &tender.Int{Value: int64(width)},
                    &tender.Int{Value: int64(height)},
                }}, nil
            },
        }
    case "text_rect_format":
        res = &tender.BuiltinFunction{
            Value: func(args ...tender.Object) (tender.Object, error) {
                if len(args) != 7 {
                    return nil, tender.ErrWrongNumArguments
                }
                x, _ := tender.ToInt(args[0])
                y, _ := tender.ToInt(args[1])
                w, _ := tender.ToInt(args[2])
                h, _ := tender.ToInt(args[3])
                s, ok := tender.ToString(args[4])
                if !ok {
                    return nil, tender.ErrInvalidArgumentType{
                        Name:     "text",
                        Expected: "string",
                        Found:    args[4].TypeName(),
                    }
                }
                format, ok := extractFormat(args[5])
                if !ok {
                    return nil, tender.ErrInvalidArgumentType{
                        Name:     "format",
                        Expected: "format",
                        Found:    args[5].TypeName(),
                    }
                }
                color, ok := extractColor(args[6])
                if !ok {
                    return nil, tender.ErrInvalidArgumentType{
                        Name:     "color",
                        Expected: "color",
                        Found:    args[6].TypeName(),
                    }
                }
                c.Value.TextRectFormat(x, y, w, h, s, format, color)
                return tender.NullValue, nil
            },
        }
    case "width":
        res = &tender.BuiltinFunction{
            Value: func(args ...tender.Object) (tender.Object, error) {
                if len(args) != 0 {
                    return nil, tender.ErrWrongNumArguments
                }
                return &tender.Int{Value: int64(c.Value.Width())}, nil
            },
        }
    }
    return
}


// Color wrapper
type WUIColor struct {
	tender.ObjectImpl
	Value wui.Color
}

func (c *WUIColor) TypeName() string { return "color" }
func (c *WUIColor) String() string   { return fmt.Sprintf("<color r:%d g:%d b:%d>", c.Value.R(), c.Value.G(), c.Value.B()) }
func (c *WUIColor) Copy() tender.Object {
	return &WUIColor{Value: c.Value}
}

// Image wrapper
type WUIImage struct {
	tender.ObjectImpl
	Value *wui.Image
}

func (i *WUIImage) TypeName() string { return "image" }
func (i *WUIImage) String() string   { return "<image>" }
func (i *WUIImage) Copy() tender.Object {
	return &WUIImage{Value: i.Value}
}

// Rectangle wrapper
type WUIRectangle struct {
	tender.ObjectImpl
	Value wui.Rectangle
}

func (r *WUIRectangle) TypeName() string { return "rectangle" }
func (r *WUIRectangle) String() string   { 
	x, y, w, h := r.Value.X, r.Value.Y, r.Value.Width, r.Value.Height
	return fmt.Sprintf("<rectangle x:%d y:%d w:%d h:%d>", x, y, w, h)
}
func (r *WUIRectangle) Copy() tender.Object {
	return &WUIRectangle{Value: r.Value}
}

// Format wrapper
type WUIFormat struct {
	tender.ObjectImpl
	Value wui.Format
}

func (f *WUIFormat) TypeName() string { return "format" }
func (f *WUIFormat) String() string   { return "<format>" }
func (f *WUIFormat) Copy() tender.Object {
	return &WUIFormat{Value: f.Value}
}

// Helper functions for extraction

func extractColor(obj tender.Object) (wui.Color, bool) {
	switch c := obj.(type) {
	case *WUIColor:
		return c.Value, true
	default:
		// Try to create color from integers or other representations
		// For now, return a default color and false
		return wui.Color(0), false
	}
}

func extractFormat(obj tender.Object) (wui.Format, bool) {
	switch f := obj.(type) {
	case *WUIFormat:
		return f.Value, true
	default:
		return wui.Format(0), false
	}
}

// Function to create RGB color
func wuiRGB(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 3 {
		err = tender.ErrWrongNumArguments
		return
	}

	r, ok1 := tender.ToInt(args[0])
	g, ok2 := tender.ToInt(args[1])
	b, ok3 := tender.ToInt(args[2])
	
	if !ok1 || !ok2 || !ok3 {
		err = tender.ErrInvalidArgumentType{
			Name:     "color components",
			Expected: "integers",
			Found:    fmt.Sprintf("%s, %s, %s", args[0].TypeName(), args[1].TypeName(), args[2].TypeName()),
		}
		return
	}

	color := wui.RGB(uint8(r), uint8(g), uint8(b))
	return &WUIColor{Value: color}, nil
}

// Function to create rectangle
func wuiRect(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 4 {
		err = tender.ErrWrongNumArguments
		return
	}

	x, ok1 := tender.ToInt(args[0])
	y, ok2 := tender.ToInt(args[1])
	width, ok3 := tender.ToInt(args[2])
	height, ok4 := tender.ToInt(args[3])
	
	if !ok1 || !ok2 || !ok3 || !ok4 {
		err = tender.ErrInvalidArgumentType{
			Name:     "rectangle components",
			Expected: "integers",
			Found:    fmt.Sprintf("%s, %s, %s, %s", args[0].TypeName(), args[1].TypeName(), args[2].TypeName(), args[3].TypeName()),
		}
		return
	}

	rect := wui.Rect(x, y, width, height)
	return &WUIRectangle{Value: rect}, nil
}
package stdlib

import (
	"fmt"
	"image"
	"bytes"

	"github.com/2dprototype/tender"
	"github.com/gonutz/wui/v2"
)

var wuiModule = map[string]tender.Object{
	"new_window": &tender.UserFunction{
		Name:  "new_window",
		Value: wuiNewWindow,
	}, // new_window() => Window
	"message_box": &tender.UserFunction{
		Name:  "message_box",
		Value: wuiMessageBox,
	}, // message_box(caption, text) => nil
	"message_box_error": &tender.UserFunction{
		Name:  "message_box_error",
		Value: wuiMessageBoxError,
	}, // message_box_error(caption, text) => nil
	"message_box_info": &tender.UserFunction{
		Name:  "message_box_info",
		Value: wuiMessageBoxInfo,
	}, // message_box_info(caption, text) => nil
	"message_box_warning": &tender.UserFunction{
		Name:  "message_box_warning",
		Value: wuiMessageBoxWarning,
	}, // message_box_warning(caption, text) => nil
	"message_box_question": &tender.UserFunction{
		Name:  "message_box_question",
		Value: wuiMessageBoxQuestion,
	}, // message_box_question(caption, text) => nil
	"message_box_ok_cancel": &tender.UserFunction{
		Name:  "message_box_ok_cancel",
		Value: wuiMessageBoxOKCancel,
	}, // message_box_ok_cancel(caption, text) => bool
	"message_box_yes_no": &tender.UserFunction{
		Name:  "message_box_yes_no",
		Value: wuiMessageBoxYesNo,
	}, // message_box_yes_no(caption, text) => bool
	"message_box_custom": &tender.UserFunction{
		Name:  "message_box_custom",
		Value: wuiMessageBoxCustom,
	}, // message_box_custom(caption, text, flags) => int
	"new_button": &tender.UserFunction{
		Name:  "new_button",
		Value: wuiNewButton,
	}, // new_button() => Button
	"new_checkbox": &tender.UserFunction{
		Name:  "new_checkbox",
		Value: wuiNewCheckBox,
	}, // new_checkbox() => CheckBox
	"new_label": &tender.UserFunction{
		Name:  "new_label",
		Value: wuiNewLabel,
	}, // new_label() => Label
	"new_edit_line": &tender.UserFunction{
		Name:  "new_edit_line",
		Value: wuiNewEditLine,
	}, // new_edit_line() => EditLine
	"new_text_edit": &tender.UserFunction{
		Name:  "new_text_edit",
		Value: wuiNewTextEdit,
	}, // new_text_edit() => TextEdit
	"new_combo_box": &tender.UserFunction{
		Name:  "new_combo_box",
		Value: wuiNewComboBox,
	}, // new_combo_box() => ComboBox
	"new_string_list": &tender.UserFunction{
		Name:  "new_string_list",
		Value: wuiNewStringList,
	}, // new_string_list() => StringList
	"new_string_table": &tender.UserFunction{
		Name:  "new_string_table",
		Value: wuiNewStringTable,
	}, // new_string_table(header1, ...) => StringTable
	"new_slider": &tender.UserFunction{
		Name:  "new_slider",
		Value: wuiNewSlider,
	}, // new_slider() => Slider
	"new_progress_bar": &tender.UserFunction{
		Name:  "new_progress_bar",
		Value: wuiNewProgressBar,
	}, // new_progress_bar() => ProgressBar
	"new_radio_button": &tender.UserFunction{
		Name:  "new_radio_button",
		Value: wuiNewRadioButton,
	}, // new_radio_button() => RadioButton
	"new_int_up_down": &tender.UserFunction{
		Name:  "new_int_up_down",
		Value: wuiNewIntUpDown,
	}, // new_int_up_down() => IntUpDown
	"new_float_up_down": &tender.UserFunction{
		Name:  "new_float_up_down",
		Value: wuiNewFloatUpDown,
	}, // new_float_up_down() => FloatUpDown
	"new_panel": &tender.UserFunction{
		Name:  "new_panel",
		Value: wuiNewPanel,
	}, // new_panel() => Panel
	"new_paint_box": &tender.UserFunction{
		Name:  "new_paint_box",
		Value: wuiNewPaintBox,
	}, // new_paint_box() => PaintBox
	"new_file_open_dialog": &tender.UserFunction{
		Name:  "new_file_open_dialog",
		Value: wuiNewFileOpenDialog,
	}, // new_file_open_dialog() => FileOpenDialog
	"new_file_save_dialog": &tender.UserFunction{
		Name:  "new_file_save_dialog",
		Value: wuiNewFileSaveDialog,
	}, // new_file_save_dialog() => FileSaveDialog
	"new_folder_select_dialog": &tender.UserFunction{
		Name:  "new_folder_select_dialog",
		Value: wuiNewFolderSelectDialog,
	}, // new_folder_select_dialog() => FolderSelectDialog
	"rgb_color": &tender.UserFunction{
		Name:  "rgb_color",
		Value: wuiRGBColor,
	}, // rgb_color(r, g, b) => Color
	"new_font": &tender.UserFunction{
		Name:  "new_font",
		Value: wuiNewFont,
	}, // new_font(desc) => Font/error
	"new_cursor_from_image": &tender.UserFunction{
		Name:  "new_cursor_from_image",
		Value: wuiNewCursorFromImage,
	}, // new_cursor_from_image(image_bytes, x, y) => Cursor/error
	"new_icon_from_image": &tender.UserFunction{
		Name:  "new_icon_from_image",
		Value: wuiNewIconFromImage,
	}, // new_icon_from_image(image_bytes) => Icon/error
	"new_menu": &tender.UserFunction{
		Name:  "new_menu",
		Value: wuiNewMenu,
	}, // new_menu(name) => Menu
	"new_menu_string": &tender.UserFunction{
		Name:  "new_menu_string",
		Value: wuiNewMenuString,
	}, // new_menu_string(text) => MenuString
	"new_menu_separator": &tender.UserFunction{
		Name:  "new_menu_separator",
		Value: wuiNewMenuSeparator,
	}, // new_menu_separator() => MenuItem
	"new_main_menu": &tender.UserFunction{
		Name:  "new_main_menu",
		Value: wuiNewMainMenu,
	}, // new_main_menu() => Menu
	"enabled": &tender.UserFunction{
		Name:  "enabled",
		Value: wuiEnabled,
	}, // enabled(control) => bool
	"visible": &tender.UserFunction{
		Name:  "visible",
		Value: wuiVisible,
	}, // visible(control) => bool
}

// Wrapper types for WUI controls
type WUIWindow struct {
	tender.ObjectImpl
	Value *wui.Window
}

func (w *WUIWindow) TypeName() string { return "window" }
func (w *WUIWindow) String() string   { return "<window>" }
func (w *WUIWindow) Copy() tender.Object {
	return &WUIWindow{Value: w.Value}
}

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

type WUIFileOpenDialog struct {
	tender.ObjectImpl
	Value *wui.FileOpenDialog
}

func (f *WUIFileOpenDialog) TypeName() string { return "fileopendialog" }
func (f *WUIFileOpenDialog) String() string   { return "<fileopendialog>" }
func (f *WUIFileOpenDialog) Copy() tender.Object {
	return &WUIFileOpenDialog{Value: f.Value}
}

type WUIFileSaveDialog struct {
	tender.ObjectImpl
	Value *wui.FileSaveDialog
}

func (f *WUIFileSaveDialog) TypeName() string { return "filesavedialog" }
func (f *WUIFileSaveDialog) String() string   { return "<filesavedialog>" }
func (f *WUIFileSaveDialog) Copy() tender.Object {
	return &WUIFileSaveDialog{Value: f.Value}
}

type WUIFolderSelectDialog struct {
	tender.ObjectImpl
	Value *wui.FolderSelectDialog
}

func (f *WUIFolderSelectDialog) TypeName() string { return "folderselectdialog" }
func (f *WUIFolderSelectDialog) String() string   { return "<folderselectdialog>" }
func (f *WUIFolderSelectDialog) Copy() tender.Object {
	return &WUIFolderSelectDialog{Value: f.Value}
}

type WUIColor struct {
	tender.ObjectImpl
	Value wui.Color
}

func (c *WUIColor) TypeName() string { return "color" }
func (c *WUIColor) String() string   { return "<color>" }
func (c *WUIColor) Copy() tender.Object {
	return &WUIColor{Value: c.Value}
}

type WUIFont struct {
	tender.ObjectImpl
	Value *wui.Font
}

func (f *WUIFont) TypeName() string { return "font" }
func (f *WUIFont) String() string   { return "<font>" }
func (f *WUIFont) Copy() tender.Object {
	return &WUIFont{Value: f.Value}
}

type WUICursor struct {
	tender.ObjectImpl
	Value *wui.Cursor
}

func (c *WUICursor) TypeName() string { return "cursor" }
func (c *WUICursor) String() string   { return "<cursor>" }
func (c *WUICursor) Copy() tender.Object {
	return &WUICursor{Value: c.Value}
}

type WUIIcon struct {
	tender.ObjectImpl
	Value *wui.Icon
}

func (i *WUIIcon) TypeName() string { return "icon" }
func (i *WUIIcon) String() string   { return "<icon>" }
func (i *WUIIcon) Copy() tender.Object {
	return &WUIIcon{Value: i.Value}
}

type WUIMenu struct {
	tender.ObjectImpl
	Value *wui.Menu
}

func (m *WUIMenu) TypeName() string { return "menu" }
func (m *WUIMenu) String() string   { return "<menu>" }
func (m *WUIMenu) Copy() tender.Object {
	return &WUIMenu{Value: m.Value}
}

type WUIMenuString struct {
	tender.ObjectImpl
	Value *wui.MenuString
}

func (m *WUIMenuString) TypeName() string { return "menustring" }
func (m *WUIMenuString) String() string   { return "<menustring>" }
func (m *WUIMenuString) Copy() tender.Object {
	return &WUIMenuString{Value: m.Value}
}

type WUIMenuItem struct {
	tender.ObjectImpl
	Value wui.MenuItem
}

func (m *WUIMenuItem) TypeName() string { return "menuitem" }
func (m *WUIMenuItem) String() string   { return "<menuitem>" }
func (m *WUIMenuItem) Copy() tender.Object {
	return &WUIMenuItem{Value: m.Value}
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

// Helper function to extract font from wrapper
func extractFont(obj tender.Object) (*wui.Font, bool) {
	if font, ok := obj.(*WUIFont); ok {
		return font.Value, true
	}
	return nil, false
}

func wuiNewWindow(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 0 {
		err = tender.ErrWrongNumArguments
		return
	}

	window := wui.NewWindow()
	return &WUIWindow{Value: window}, nil
}

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

func wuiNewFileOpenDialog(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 0 {
		err = tender.ErrWrongNumArguments
		return
	}

	fileOpenDialog := wui.NewFileOpenDialog()
	return &WUIFileOpenDialog{Value: fileOpenDialog}, nil
}

func wuiNewFileSaveDialog(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 0 {
		err = tender.ErrWrongNumArguments
		return
	}

	fileSaveDialog := wui.NewFileSaveDialog()
	return &WUIFileSaveDialog{Value: fileSaveDialog}, nil
}

func wuiNewFolderSelectDialog(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 0 {
		err = tender.ErrWrongNumArguments
		return
	}

	folderSelectDialog := wui.NewFolderSelectDialog()
	return &WUIFolderSelectDialog{Value: folderSelectDialog}, nil
}

func wuiRGBColor(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 3 {
		err = tender.ErrWrongNumArguments
		return
	}

	r, ok := tender.ToInt(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "red",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	g, ok := tender.ToInt(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "green",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	b, ok := tender.ToInt(args[2])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "blue",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}

	color := wui.RGB(uint8(r), uint8(g), uint8(b))
	return &WUIColor{Value: color}, nil
}

func wuiNewFont(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	descMap, ok := args[0].(*tender.Map)
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "font_desc",
			Expected: "map",
			Found:    args[0].TypeName(),
		}
		return
	}

	desc := wui.FontDesc{}

	if val, ok := descMap.Value["name"]; ok {
		desc.Name, _ = tender.ToString(val)
	}
	if val, ok := descMap.Value["height"]; ok {
		desc.Height, _ = tender.ToInt(val)
	}
	if val, ok := descMap.Value["bold"]; ok {
		desc.Bold, _ = tender.ToBool(val)
	}
	if val, ok := descMap.Value["italic"]; ok {
		desc.Italic, _ = tender.ToBool(val)
	}
	if val, ok := descMap.Value["underlined"]; ok {
		desc.Underlined, _ = tender.ToBool(val)
	}
	if val, ok := descMap.Value["striked_out"]; ok {
		desc.StrikedOut, _ = tender.ToBool(val)
	}

	font, err := wui.NewFont(desc)
	if err != nil {
		ret = wrapError(err)
		return
	}

	return &WUIFont{Value: font}, nil
}

func wuiNewCursorFromImage(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 3 {
		err = tender.ErrWrongNumArguments
		return
	}

	imageBytes, ok := tender.ToByteSlice(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "image",
			Expected: "bytes",
			Found:    args[0].TypeName(),
		}
		return
	}

	x, ok := tender.ToInt(args[1])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "x",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	y, ok := tender.ToInt(args[2])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "y",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}

	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		ret = wrapError(err)
		return
	}

	cursor, err := wui.NewCursorFromImage(img, x, y)
	if err != nil {
		ret = wrapError(err)
		return
	}

	return &WUICursor{Value: cursor}, nil
}

func wuiNewIconFromImage(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	imageBytes, ok := tender.ToByteSlice(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "image",
			Expected: "bytes",
			Found:    args[0].TypeName(),
		}
		return
	}

	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		ret = wrapError(err)
		return
	}

	icon, err := wui.NewIconFromImage(img)
	if err != nil {
		ret = wrapError(err)
		return
	}

	return &WUIIcon{Value: icon}, nil
}

func wuiNewMenu(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	name, ok := tender.ToString(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "name",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	menu := wui.NewMenu(name)
	return &WUIMenu{Value: menu}, nil
}

func wuiNewMenuString(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	text, ok := tender.ToString(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "text",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	menuString := wui.NewMenuString(text)
	return &WUIMenuString{Value: menuString}, nil
}

func wuiNewMenuSeparator(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 0 {
		err = tender.ErrWrongNumArguments
		return
	}

	menuSeparator := wui.NewMenuSeparator()
	return &WUIMenuItem{Value: menuSeparator}, nil
}

func wuiNewMainMenu(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 0 {
		err = tender.ErrWrongNumArguments
		return
	}

	mainMenu := wui.NewMainMenu()
	return &WUIMenu{Value: mainMenu}, nil
}

func wuiEnabled(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	control, ok := extractControl(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "control",
			Expected: "control",
			Found:    args[0].TypeName(),
		}
		return
	}

	if control.Enabled() {
		ret = tender.TrueValue
	} else {
		ret = tender.FalseValue
	}
	return
}

func wuiVisible(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		err = tender.ErrWrongNumArguments
		return
	}

	control, ok := extractControl(args[0])
	if !ok {
		err = tender.ErrInvalidArgumentType{
			Name:     "control",
			Expected: "control",
			Found:    args[0].TypeName(),
		}
		return
	}

	if control.Visible() {
		ret = tender.TrueValue
	} else {
		ret = tender.FalseValue
	}
	return
}

// Implement IndexGet for all wrapper types to expose their methods
func (w *WUIWindow) IndexGet(index tender.Object) (res tender.Object, err error) {
	strIdx, ok := index.(*tender.String)
	if !ok {
		return nil, tender.ErrInvalidIndexType
	}

	switch strIdx.Value {
	case "show":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				err := w.Value.Show()
				if err != nil {
					return wrapError(err), nil
				}
				return tender.NullValue, nil
			},
		}
	case "show_modal":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				err := w.Value.ShowModal()
				if err != nil {
					return wrapError(err), nil
				}
				return tender.NullValue, nil
			},
		}
	case "close":
		res = &tender.BuiltinFunction{
			Value: FuncAR(w.Value.Close),
		}
	case "set_title":
		res = &tender.BuiltinFunction{
			Value: FuncASR(w.Value.SetTitle),
		}
	case "title":
		res = &tender.BuiltinFunction{
			Value: FuncARS(w.Value.Title),
		}
	case "set_size":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, tender.ErrWrongNumArguments
				}
				width, _ := tender.ToInt(args[0])
				height, _ := tender.ToInt(args[1])
				w.Value.SetSize(width, height)
				return tender.NullValue, nil
			},
		}
	case "size":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				width, height := w.Value.Size()
				return &tender.Array{
					Value: []tender.Object{
						&tender.Int{Value: int64(width)},
						&tender.Int{Value: int64(height)},
					},
				}, nil
			},
		}
	case "add":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				control, ok := extractControl(args[0])
				if !ok {
					return nil, tender.ErrInvalidArgumentType{
						Name:     "control",
						Expected: "control",
						Found:    args[0].TypeName(),
					}
				}
				w.Value.Add(control)
				return tender.NullValue, nil
			},
		}
	case "remove":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				control, ok := extractControl(args[0])
				if !ok {
					return nil, tender.ErrInvalidArgumentType{
						Name:     "control",
						Expected: "control",
						Found:    args[0].TypeName(),
					}
				}
				w.Value.Remove(control)
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
				w.Value.SetPosition(x, y)
				return tender.NullValue, nil
			},
		}
	case "position":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				x, y := w.Value.Position()
				return &tender.Array{
					Value: []tender.Object{
						&tender.Int{Value: int64(x)},
						&tender.Int{Value: int64(y)},
					},
				}, nil
			},
		}
	}
	return
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
	// case "set_enabled":
		// res = &tender.BuiltinFunction{
			// Value: FuncABR(b.Value.SetEnabled),
		// }
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
	// case "set_visible":
		// res = &tender.BuiltinFunction{
			// Value: FuncABR(b.Value.SetVisible),
		// }
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
	// case "set_checked":
		// res = &tender.BuiltinFunction{
			// Value: FuncABR(c.Value.SetChecked),
		// }
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
	}
	return
}

func (t *WUITextEdit) IndexGet(index tender.Object) (res tender.Object, err error) {
	strIdx, ok := index.(*tender.String)
	if !ok {
		return nil, tender.ErrInvalidIndexType
	}

	switch strIdx.Value {
	case "set_text":
		res = &tender.BuiltinFunction{
			Value: FuncASR(t.Value.SetText),
		}
	case "text":
		res = &tender.BuiltinFunction{
			Value: FuncARS(t.Value.Text),
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
				t.Value.SetBounds(x, y, width, height)
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
				t.Value.SetFont(font)
				return tender.NullValue, nil
			},
		}
	}
	return
}

func (c *WUIComboBox) IndexGet(index tender.Object) (res tender.Object, err error) {
	strIdx, ok := index.(*tender.String)
	if !ok {
		return nil, tender.ErrInvalidIndexType
	}

	switch strIdx.Value {
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

func (s *WUIStringList) IndexGet(index tender.Object) (res tender.Object, err error) {
	strIdx, ok := index.(*tender.String)
	if !ok {
		return nil, tender.ErrInvalidIndexType
	}

	switch strIdx.Value {
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
				s.Value.SetItems(items)
				return tender.NullValue, nil
			},
		}
	case "add_item":
		res = &tender.BuiltinFunction{
			Value: FuncASR(s.Value.AddItem),
		}
	case "clear":
		res = &tender.BuiltinFunction{
			Value: FuncAR(s.Value.Clear),
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
				s.Value.SetBounds(x, y, width, height)
				return tender.NullValue, nil
			},
		}
	}
	return
}

func (s *WUIStringTable) IndexGet(index tender.Object) (res tender.Object, err error) {
	strIdx, ok := index.(*tender.String)
	if !ok {
		return nil, tender.ErrInvalidIndexType
	}

	switch strIdx.Value {
	// case "set_column_count":
		// res = &tender.BuiltinFunction{
			// Value: FuncAIR(s.Value.SetColumnCount),
		// }
	// case "set_column_width":
		// res = &tender.BuiltinFunction{
			// Value: func(args ...tender.Object) (tender.Object, error) {
				// if len(args) != 2 {
					// return nil, tender.ErrWrongNumArguments
				// }
				// col, _ := tender.ToInt(args[0])
				// width, _ := tender.ToInt(args[1])
				// s.Value.SetColumnWidth(col, width)
				// return tender.NullValue, nil
			// },
		// }
	// case "add_row":
		// res = &tender.BuiltinFunction{
			// Value: func(args ...tender.Object) (tender.Object, error) {
				// if len(args) != 1 {
					// return nil, tender.ErrWrongNumArguments
				// }
				// arr, ok := args[0].(*tender.Array)
				// if !ok {
					// return nil, tender.ErrInvalidArgumentType{
						// Name:     "row",
						// Expected: "array",
						// Found:    args[0].TypeName(),
					// }
				// }
				// row := make([]string, len(arr.Value))
				// for i, item := range arr.Value {
					// row[i], _ = tender.ToString(item)
				// }
				// s.Value.AddRow(row)
				// return tender.NullValue, nil
			// },
		// }
	case "clear":
		res = &tender.BuiltinFunction{
			Value: FuncAR(s.Value.Clear),
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
				s.Value.SetBounds(x, y, width, height)
				return tender.NullValue, nil
			},
		}
	}
	return
}

func (s *WUISlider) IndexGet(index tender.Object) (res tender.Object, err error) {
	strIdx, ok := index.(*tender.String)
	if !ok {
		return nil, tender.ErrInvalidIndexType
	}

	switch strIdx.Value {
	case "set_min":
		res = &tender.BuiltinFunction{
			Value: FuncAIR(s.Value.SetMin),
		}
	case "set_max":
		res = &tender.BuiltinFunction{
			Value: FuncAIR(s.Value.SetMax),
		}
	// case "set_value":
		// res = &tender.BuiltinFunction{
			// Value: FuncAIR(s.Value.SetValue),
		// }
	// case "value":
		// res = &tender.BuiltinFunction{
			// Value: func(args ...tender.Object) (tender.Object, error) {
				// if len(args) != 0 {
					// return nil, tender.ErrWrongNumArguments
				// }
				// return &tender.Int{Value: int64(s.Value.Value())}, nil
			// },
		// }
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
				s.Value.SetBounds(x, y, width, height)
				return tender.NullValue, nil
			},
		}
	}
	return
}

func (p *WUIProgressBar) IndexGet(index tender.Object) (res tender.Object, err error) {
	strIdx, ok := index.(*tender.String)
	if !ok {
		return nil, tender.ErrInvalidIndexType
	}

	switch strIdx.Value {
	// case "set_min":
		// res = &tender.BuiltinFunction{
			// Value: FuncAIR(p.Value.SetMin),
		// }
	// case "set_max":
		// res = &tender.BuiltinFunction{
			// Value: FuncAIR(p.Value.SetMax),
		// }
	// case "set_value":
		// res = &tender.BuiltinFunction{
			// Value: FuncAIR(p.Value.SetValue),
		// }
	case "value":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				return &tender.Int{Value: int64(p.Value.Value())}, nil
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
				p.Value.SetBounds(x, y, width, height)
				return tender.NullValue, nil
			},
		}
	}
	return
}

func (r *WUIRadioButton) IndexGet(index tender.Object) (res tender.Object, err error) {
	strIdx, ok := index.(*tender.String)
	if !ok {
		return nil, tender.ErrInvalidIndexType
	}

	switch strIdx.Value {
	case "set_text":
		res = &tender.BuiltinFunction{
			Value: FuncASR(r.Value.SetText),
		}
	case "text":
		res = &tender.BuiltinFunction{
			Value: FuncARS(r.Value.Text),
		}
	// case "set_checked":
		// res = &tender.BuiltinFunction{
			// Value: FuncABR(r.Value.SetChecked),
		// }
	case "checked":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, tender.ErrWrongNumArguments
				}
				if r.Value.Checked() {
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
				r.Value.SetBounds(x, y, width, height)
				return tender.NullValue, nil
			},
		}
	}
	return
}

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
	}
	return
}

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
	}
	return
}

func (p *WUIPanel) IndexGet(index tender.Object) (res tender.Object, err error) {
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
	case "add":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				control, ok := extractControl(args[0])
				if !ok {
					return nil, tender.ErrInvalidArgumentType{
						Name:     "control",
						Expected: "control",
						Found:    args[0].TypeName(),
					}
				}
				p.Value.Add(control)
				return tender.NullValue, nil
			},
		}
	case "remove":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				control, ok := extractControl(args[0])
				if !ok {
					return nil, tender.ErrInvalidArgumentType{
						Name:     "control",
						Expected: "control",
						Found:    args[0].TypeName(),
					}
				}
				p.Value.Remove(control)
				return tender.NullValue, nil
			},
		}
	}
	return
}

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
	}
	return
}

// func (f *WUIFileOpenDialog) IndexGet(index tender.Object) (res tender.Object, err error) {
	// strIdx, ok := index.(*tender.String)
	// if !ok {
		// return nil, tender.ErrInvalidIndexType
	// }

	// switch strIdx.Value {
	// case "execute":
		// res = &tender.BuiltinFunction{
			// Value: func(args ...tender.Object) (tender.Object, error) {
				// if len(args) != 0 {
					// return nil, tender.ErrWrongNumArguments
				// }
				// if f.Value.Execute() {
					// return tender.TrueValue, nil
				// }
				// return tender.FalseValue, nil
			// },
		// }
	// case "file":
		// res = &tender.BuiltinFunction{
			// Value: func(args ...tender.Object) (tender.Object, error) {
				// if len(args) != 0 {
					// return nil, tender.ErrWrongNumArguments
				// }
				// return &tender.String{Value: f.Value.File()}, nil
			// },
		// }
	// case "set_file":
		// res = &tender.BuiltinFunction{
			// Value: FuncASR(f.Value.SetFile),
		// }
	// case "set_filter":
		// res = &tender.BuiltinFunction{
			// Value: FuncASR(f.Value.SetFilter),
		// }
	// }
	// return
// }

// func (f *WUIFileSaveDialog) IndexGet(index tender.Object) (res tender.Object, err error) {
	// strIdx, ok := index.(*tender.String)
	// if !ok {
		// return nil, tender.ErrInvalidIndexType
	// }

	// switch strIdx.Value {
	// case "execute":
		// res = &tender.BuiltinFunction{
			// Value: func(args ...tender.Object) (tender.Object, error) {
				// if len(args) != 0 {
					// return nil, tender.ErrWrongNumArguments
				// }
				// if f.Value.Execute() {
					// return tender.TrueValue, nil
				// }
				// return tender.FalseValue, nil
			// },
		// }
	// case "file":
		// res = &tender.BuiltinFunction{
			// Value: func(args ...tender.Object) (tender.Object, error) {
				// if len(args) != 0 {
					// return nil, tender.ErrWrongNumArguments
				// }
				// return &tender.String{Value: f.Value.File()}, nil
			// },
		// }
	// case "set_file":
		// res = &tender.BuiltinFunction{
			// Value: FuncASR(f.Value.SetFile),
		// }
	// case "set_filter":
		// res = &tender.BuiltinFunction{
			// Value: FuncASR(f.Value.SetFilter),
		// }
	// }
	// return
// }

// func (f *WUIFolderSelectDialog) IndexGet(index tender.Object) (res tender.Object, err error) {
	// strIdx, ok := index.(*tender.String)
	// if !ok {
		// return nil, tender.ErrInvalidIndexType
	// }

	// switch strIdx.Value {
	// case "execute":
		// res = &tender.BuiltinFunction{
			// Value: func(args ...tender.Object) (tender.Object, error) {
				// if len(args) != 0 {
					// return nil, tender.ErrWrongNumArguments
				// }
				// if f.Value.Execute() {
					// return tender.TrueValue, nil
				// }
				// return tender.FalseValue, nil
			// },
		// }
	// case "folder":
		// res = &tender.BuiltinFunction{
			// Value: func(args ...tender.Object) (tender.Object, error) {
				// if len(args) != 0 {
					// return nil, tender.ErrWrongNumArguments
				// }
				// return &tender.String{Value: f.Value.Folder()}, nil
			// },
		// }
	// case "set_folder":
		// res = &tender.BuiltinFunction{
			// Value: FuncASR(f.Value.SetFolder),
		// }
	// }
	// return
// }

func (c *WUIColor) IndexGet(index tender.Object) (res tender.Object, err error) {
	strIdx, ok := index.(*tender.String)
	if !ok {
		return nil, tender.ErrInvalidIndexType
	}

	switch strIdx.Value {
	case "r":
		res = &tender.Int{Value: int64(c.Value.R())}
	case "g":
		res = &tender.Int{Value: int64(c.Value.G())}
	case "b":
		res = &tender.Int{Value: int64(c.Value.B())}
	}
	return
}

// func (f *WUIFont) IndexGet(index tender.Object) (res tender.Object, err error) {
	// strIdx, ok := index.(*tender.String)
	// if !ok {
		// return nil, tender.ErrInvalidIndexType
	// }

	// switch strIdx.Value {
	// case "height":
		// res = &tender.BuiltinFunction{
			// Value: func(args ...tender.Object) (tender.Object, error) {
				// if len(args) != 0 {
					// return nil, tender.ErrWrongNumArguments
				// }
				// return &tender.Int{Value: int64(f.Value.Height())}, nil
			// },
		// }
	// }
	// return
// }

// func (c *WUICursor) IndexGet(index tender.Object) (res tender.Object, err error) {
	// strIdx, ok := index.(*tender.String)
	// if !ok {
		// return nil, tender.ErrInvalidIndexType
	// }

	// // Cursor doesn't have methods in wui, but we can add if needed
	// return
// }

// func (i *WUIIcon) IndexGet(index tender.Object) (res tender.Object, err error) {
	// strIdx, ok := index.(*tender.String)
	// if !ok {
		// return nil, tender.ErrInvalidIndexType
	// }

	// // Icon doesn't have methods in wui, but we can add if needed
	// return
// }

func (m *WUIMenu) IndexGet(index tender.Object) (res tender.Object, err error) {
	strIdx, ok := index.(*tender.String)
	if !ok {
		return nil, tender.ErrInvalidIndexType
	}

	switch strIdx.Value {
	case "add":
		res = &tender.BuiltinFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, tender.ErrWrongNumArguments
				}
				menuItem, ok := args[0].(*WUIMenuItem)
				if !ok {
					return nil, tender.ErrInvalidArgumentType{
						Name:     "menuitem",
						Expected: "menuitem",
						Found:    args[0].TypeName(),
					}
				}
				m.Value.Add(menuItem.Value)
				return tender.NullValue, nil
			},
		}
	}
	return
}

func (m *WUIMenuString) IndexGet(index tender.Object) (res tender.Object, err error) {
	strIdx, ok := index.(*tender.String)
	if !ok {
		return nil, tender.ErrInvalidIndexType
	}

	switch strIdx.Value {
	case "set_text":
		res = &tender.BuiltinFunction{
			Value: FuncASR(m.Value.SetText),
		}
	case "text":
		res = &tender.BuiltinFunction{
			Value: FuncARS(m.Value.Text),
		}
	// case "set_on_click":
		// res = &tender.BuiltinFunction{
			// Value: FuncAFR(m.Value.SetOnClick),
		// }
	}
	return
}

// func (m *WUIMenuItem) IndexGet(index tender.Object) (res tender.Object, err error) {
	// strIdx, ok := index.(*tender.String)
	// if !ok {
		// return nil, tender.ErrInvalidIndexType
	// }

	// // MenuItem is an interface, methods depend on concrete type
	// return
// }

// Common helper functions for WUI module
func wuiSetBounds(control interface{ SetBounds(x, y, width, height int) }, args []tender.Object) (tender.Object, error) {
	if len(args) != 4 {
		return nil, tender.ErrWrongNumArguments
	}
	x, _ := tender.ToInt(args[0])
	y, _ := tender.ToInt(args[1])
	width, _ := tender.ToInt(args[2])
	height, _ := tender.ToInt(args[3])
	control.SetBounds(x, y, width, height)
	return tender.NullValue, nil
}

func wuiSetPosition(control interface{ SetPosition(x, y int) }, args []tender.Object) (tender.Object, error) {
	if len(args) != 2 {
		return nil, tender.ErrWrongNumArguments
	}
	x, _ := tender.ToInt(args[0])
	y, _ := tender.ToInt(args[1])
	control.SetPosition(x, y)
	return tender.NullValue, nil
}

func wuiSetSize(control interface{ SetSize(width, height int) }, args []tender.Object) (tender.Object, error) {
	if len(args) != 2 {
		return nil, tender.ErrWrongNumArguments
	}
	width, _ := tender.ToInt(args[0])
	height, _ := tender.ToInt(args[1])
	control.SetSize(width, height)
	return tender.NullValue, nil
}
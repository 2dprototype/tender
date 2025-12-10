package stdlib

import (
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
	"new_editline": &tender.UserFunction{
		Name:  "new_editline",
		Value: wuiNewEditLine,
	}, // new_edit_line() => EditLine
	"new_textedit": &tender.UserFunction{
		Name:  "new_textedit",
		Value: wuiNewTextEdit,
	}, // new_text_edit() => TextEdit
	"new_combobox": &tender.UserFunction{
		Name:  "new_combo_box",
		Value: wuiNewComboBox,
	}, // new_combo_box() => ComboBox
	"new_stringlist": &tender.UserFunction{
		Name:  "new_stringlist",
		Value: wuiNewStringList,
	}, // new_string_list() => StringList
	"new_stringtable": &tender.UserFunction{
		Name:  "new_stringtable",
		Value: wuiNewStringTable,
	}, // new_string_table(header1, ...) => StringTable
	"new_slider": &tender.UserFunction{
		Name:  "new_slider",
		Value: wuiNewSlider,
	}, // new_slider() => Slider
	"new_progressbar": &tender.UserFunction{
		Name:  "new_progressbar",
		Value: wuiNewProgressBar,
	}, // new_progress_bar() => ProgressBar
	"new_radiobutton": &tender.UserFunction{
		Name:  "new_radiobutton",
		Value: wuiNewRadioButton,
	}, // new_radio_button() => RadioButton
	"new_intupdown": &tender.UserFunction{
		Name:  "new_intupdown",
		Value: wuiNewIntUpDown,
	}, // new_int_up_down() => IntUpDown
	"new_floatupdown": &tender.UserFunction{
		Name:  "new_floatupdown",
		Value: wuiNewFloatUpDown,
	}, // new_float_up_down() => FloatUpDown
	"new_panel": &tender.UserFunction{
		Name:  "new_panel",
		Value: wuiNewPanel,
	}, // new_panel() => Panel
	"new_paintbox": &tender.UserFunction{
		Name:  "new_paintbox",
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
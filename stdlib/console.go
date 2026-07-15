package stdlib

import (
	"os"
	
	"github.com/2dprototype/tender"
	"github.com/2dprototype/tender/v/colorable"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

var consoleModule = map[string]tender.Object{
	"is_terminal": &tender.NativeFunction{
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrWrongNumArguments
			}
			// Get file mode
			mode, err := os.Stdout.Stat()
			if err != nil {
				return tender.NullValue, nil
			}
			// Check if it's a character device
			if mode.Mode()&os.ModeCharDevice != 0 {
				return tender.TrueValue, nil
			} else {
				return tender.FalseValue, nil
			}
		},
	},	
	"stdout": &tender.NativeFunction{
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrWrongNumArguments
			}
			return &IOWriter{Value: colorable.NewColorableStdout()}, nil
		},
	},
	"stderr": &tender.NativeFunction{
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrWrongNumArguments
			}
			return &IOWriter{Value: colorable.NewColorableStderr()}, nil
		},
	},
	"style": &tender.NativeFunction{
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) < 1 {
				return nil, tender.ErrWrongNumArguments
			}

			text, ok := tender.ToString(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgumentType{
					Name:     "text",
					Expected: "string",
					Found:    args[0].TypeName(),
				}
			}

			style := buildLipglossStyle(args[1:])
			return &tender.String{Value: style.Render(text)}, nil
		},
	},
	"table": &tender.NativeFunction{
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) < 2 {
				return nil, tender.ErrWrongNumArguments
			}

			// Parse headers
			headersArr, ok := args[0].(*tender.Array)
			if !ok {
				return nil, tender.ErrInvalidArgumentType{
					Name:     "headers",
					Expected: "array",
					Found:    args[0].TypeName(),
				}
			}
			var headers []string
			for _, h := range headersArr.Value {
				if s, ok := tender.ToString(h); ok {
					headers = append(headers, s)
				}
			}

			// Parse rows
			rowsArr, ok := args[1].(*tender.Array)
			if !ok {
				return nil, tender.ErrInvalidArgumentType{
					Name:     "rows",
					Expected: "array",
					Found:    args[1].TypeName(),
				}
			}
			var rows [][]string
			for _, rObj := range rowsArr.Value {
				rArr, ok := rObj.(*tender.Array)
				if !ok {
					continue
				}
				var row []string
				for _, cell := range rArr.Value {
					if s, ok := tender.ToString(cell); ok {
						row = append(row, s)
					}
				}
				rows = append(rows, row)
			}

			t := table.New().
				Headers(headers...).
				Rows(rows...)

			// Optional table styling
			if len(args) > 2 {
				if styleMap, ok := args[2].(*tender.Map); ok {
					if borderVal, exists := styleMap.Value["border"]; exists {
						if s, ok := tender.ToString(borderVal); ok {
							switch s {
							case "normal":
								t.Border(lipgloss.NormalBorder())
							case "rounded":
								t.Border(lipgloss.RoundedBorder())
							case "thick":
								t.Border(lipgloss.ThickBorder()) 
							case "double":
								t.Border(lipgloss.DoubleBorder())
							case "none":
								t.Border(lipgloss.HiddenBorder())
							}
						}
					}
				}
			}

			return &tender.String{Value: t.Render()}, nil
		},
	},
}

// Helper function to keep style parsing clean and reusable
func buildLipglossStyle(args []tender.Object) lipgloss.Style {
	style := lipgloss.NewStyle()
	for _, arg := range args {
		switch prop := arg.(type) {
		case *tender.Map:
			for k, v := range prop.Value {
				switch k {
				case "color", "foreground":
					if s, ok := v.(*tender.String); ok {
						style = style.Foreground(lipgloss.Color(s.Value))
					}
				case "background":
					if s, ok := v.(*tender.String); ok {
						style = style.Background(lipgloss.Color(s.Value))
					}
				case "bold":
					if b, ok := tender.ToBool(v); ok {
						style = style.Bold(b)
					}
				case "italic":
					if b, ok := tender.ToBool(v); ok {
						style = style.Italic(b)
					}
				case "underline":
					if b, ok := tender.ToBool(v); ok {
						style = style.Underline(b)
					}
				case "strikethrough":
					if b, ok := tender.ToBool(v); ok {
						style = style.Strikethrough(b)
					}
				case "faint":
					if b, ok := tender.ToBool(v); ok {
						style = style.Faint(b)
					}
				case "blink":
					if b, ok := tender.ToBool(v); ok {
						style = style.Blink(b)
					}
				case "reverse":
					if b, ok := tender.ToBool(v); ok {
						style = style.Reverse(b)
					}
				case "width":
					if i, ok := tender.ToInt(v); ok {
						style = style.Width(i)
					}
				case "height":
					if i, ok := tender.ToInt(v); ok {
						style = style.Height(i)
					}
				case "align":
					if s, ok := v.(*tender.String); ok {
						switch s.Value {
						case "left":
							style = style.Align(lipgloss.Left)
						case "center":
							style = style.Align(lipgloss.Center)
						case "right":
							style = style.Align(lipgloss.Right)
						}
					}
				case "border":
					if s, ok := v.(*tender.String); ok {
						switch s.Value {
						case "normal":
							style = style.BorderStyle(lipgloss.NormalBorder())
						case "rounded":
							style = style.BorderStyle(lipgloss.RoundedBorder())
						case "thick":
							style = style.BorderStyle(lipgloss.ThickBorder())
						case "double":
							style = style.BorderStyle(lipgloss.DoubleBorder())
						}
					}
				case "border_top":
					if b, ok := tender.ToBool(v); ok {
						style = style.BorderTop(b)
					}
				case "border_bottom":
					if b, ok := tender.ToBool(v); ok {
						style = style.BorderBottom(b)
					}
				case "border_left":
					if b, ok := tender.ToBool(v); ok {
						style = style.BorderLeft(b)
					}
				case "border_right":
					if b, ok := tender.ToBool(v); ok {
						style = style.BorderRight(b)
					}
				case "margin":
					if i, ok := tender.ToInt(v); ok {
						style = style.Margin(i)
					}
				case "padding":
					if i, ok := tender.ToInt(v); ok {
						style = style.Padding(i)
					}
				}
			}
		}
	}
	return style
}
package stdlib

import (
	"github.com/2dprototype/tender/v/colorable"
	"github.com/2dprototype/tender"
	"github.com/charmbracelet/lipgloss"
)

var colorsModule = map[string]tender.Object{
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
			
			style := lipgloss.NewStyle()
			for _, arg := range args[1:] {
				switch prop := arg.(type) {
				case *tender.Map:
					for k, v := range prop.Value {
						switch k {
						case "color":
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
			
			return &tender.String{Value: style.Render(text)}, nil
		},
	},
}
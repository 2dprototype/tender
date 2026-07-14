package stdlib

import (
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/event/touch"
	"github.com/2dprototype/tender"
)


func eventToObject(event interface{}) tender.Object {
	switch e := event.(type) {
		case lifecycle.Event:
			return  &tender.ImmutableMap{
				Value: map[string]tender.Object{
					"type": &tender.String{Value: "lifecycle"},
					"to": &tender.Int{Value: int64(e.To)},
					"from": &tender.Int{Value: int64(e.From)},
					"string": &tender.NativeFunction{
						Value: FuncARS(e.String),
					},
				},
			}	
		case mouse.Event:
			return  &tender.ImmutableMap{
				Value: map[string]tender.Object{
					"type": &tender.String{Value: "mouse"},
					"x": &tender.Int{Value: int64(e.X)},
					"y": &tender.Int{Value: int64(e.Y)},
					"button": &tender.Int{Value: int64(e.Button)},
					"is_wheel": &tender.NativeFunction{
						Value: FuncARB(e.Button.IsWheel),
					},
					"direction": &tender.Int{Value: int64(e.Direction)},
					"direction_string": &tender.NativeFunction{
						Value: FuncARS(e.Direction.String),
					},
					"modifiers": &tender.Int{Value: int64(e.Modifiers)},
				},
			}
		case key.Event:
			return  &tender.ImmutableMap{
				Value: map[string]tender.Object{
					"type": &tender.String{Value: "key"},
					"rune": &tender.Char{Value: e.Rune},
					"code": &tender.Int{Value: int64(e.Code)},
					"direction": &tender.Int{Value: int64(e.Direction)},
					"modifiers": &tender.Int{Value: int64(e.Modifiers)},
					"string": &tender.NativeFunction{
						Value: FuncARS(e.String),
					},
				},
			}
		case paint.Event:
			return  &tender.ImmutableMap{
				Value: map[string]tender.Object{
					"type": &tender.String{Value: "paint"},
					"bool": tender.FromBool(e.External),
				},
			}
		case size.Event:
			return  &tender.ImmutableMap{
				Value: map[string]tender.Object{
					"type": &tender.String{Value: "size"},
					"width_px": &tender.Int{Value: int64(e.WidthPx)},
					"height_px": &tender.Int{Value: int64(e.HeightPx)},		
					"width_pt": &tender.Float{Value: float64(e.WidthPt)},
					"height_pt": &tender.Float{Value: float64(e.HeightPt)},
					"pixels_per_pt": &tender.Float{Value: float64(e.PixelsPerPt)},
					"orientation": &tender.Int{Value: int64(e.Orientation)},
				},
			}
		case touch.Event:
			return  &tender.ImmutableMap{
				Value: map[string]tender.Object{
					"type": &tender.String{Value: "touch"},
					"x": &tender.Float{Value: float64(e.X)},
					"y": &tender.Float{Value: float64(e.Y)},		
					"sequence": &tender.Int{Value: int64(e.Sequence)},
					"touch_type": &tender.Int{Value: int64(e.Type)},
				},
			}
		default:
			return nil
	}
}

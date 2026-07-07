//go:build windows

package stdlib

import (
	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
	"bytes"
	// "time"
	"github.com/2dprototype/tender"
	"os"
)

var otoCtx *oto.Context

var audioModule = map[string]tender.Object{
	"init": &tender.UserFunction{Name: "init", Value: audioInit},
	"player": &tender.UserFunction{Name: "player", Value: audioPlayer},
}

func audioInit(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 0 {
		return nil, tender.ErrWrongNumArguments
	}
    op := &oto.NewContextOptions{}
    op.SampleRate = 44100
	op.ChannelCount = 2
	// if len(args) >= 1 {
		// SampleRate, ok := tender.ToInt(args[0])
		// if ok {
			// op.SampleRate = SampleRate
		// } else {
			// op.SampleRate = 44100
		// }
	// }
	
	// if len(args) == 2 {
		// ChannelCount, ok := tender.ToInt(args[1])
		// if ok && ChannelCount != 0 {
			// op.ChannelCount = ChannelCount
		// } else {
			// op.ChannelCount = 2
		// }
	// }
	
    op.Format = oto.FormatSignedInt16LE
	var readyChan chan struct{}
    otoCtx, readyChan, err = oto.NewContext(op)
	if err != nil {
	return wrapError(err), nil
	}
	<-readyChan
	return &tender.ImmutableMap{
		Value: map[string]tender.Object{
			"suspend": &tender.UserFunction{
			Name:  "suspend",
			Value: FuncARE(otoCtx.Suspend),
			},	
			"resume": &tender.UserFunction{
				Name:  "resume",
				Value: FuncARE(otoCtx.Resume),
			},	
		},
	}, nil
}

func audioPlayer(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}
	
    var data []byte
    switch arg := args[0].(type) {
    case *tender.Bytes:
        data = arg.Value
    case *tender.String:
        data, err = os.ReadFile(tender.ResolvePath(arg.Value))
        if err != nil {
            return wrapError(err), nil
        }
    default:
        return nil, tender.ErrInvalidArgumentType{
            Name: "audio_data",
            Expected: "bytes or string (path)",
            Found: args[0].TypeName(),
        }
    }
	
	buf := bytes.NewReader(data)
	
	decoded, err := mp3.NewDecoder(buf)
	if err != nil {
		return wrapError(err), nil
	}
	
	player := otoCtx.NewPlayer(decoded)
	
	return &tender.ImmutableMap{
		Value: map[string]tender.Object{
			"decoder" : &tender.ImmutableMap{
				Value: map[string]tender.Object{
					"length": &tender.UserFunction{
						Name:  "length",
						Value: FuncARI64(decoded.Length),
					},	
					"sample_rate": &tender.UserFunction{
						Name:  "sample_rate",
						Value: FuncARI(decoded.SampleRate),
					},	
					"seek": &tender.UserFunction{
						Name:  "seek",
						Value: FuncAI64IRI64E(decoded.Seek),
					},	
				},
			},
			"play": &tender.UserFunction{
				Name:  "play",
				Value: FuncAR(player.Play),
			},		
			"pause": &tender.UserFunction{
				Name:  "pause",
				Value: FuncAR(player.Pause),
			},	
			"is_playing": &tender.UserFunction{
				Name:  "is_playing",
				Value: FuncARB(player.IsPlaying),
			},	
			"close": &tender.UserFunction{
				Name:  "close",
				Value: FuncARE(player.Close),
			},	
			"err": &tender.UserFunction{
				Name:  "err",
				Value: FuncARE(player.Err),
			},	
			"reset": &tender.UserFunction{
				Name:  "reset",
				Value: FuncAR(player.Reset),
			},	
			"buffered_size": &tender.UserFunction{
				Name:  "buffered_size",
				Value: FuncARI(player.BufferedSize),
			},	
			"set_buffer_size": &tender.UserFunction{
				Name:  "set_buffer_size",
				Value: FuncAIR(player.SetBufferSize),
			},	
			"set_volume": &tender.UserFunction{
				Name:  "set_volume",
				Value: FuncAFR(player.SetVolume),
			},
			"volume": &tender.UserFunction{
				Name:  "volume",
				Value: FuncARF(player.Volume),
			},	
			"seek": &tender.UserFunction{
				Name:  "seek",
				Value: FuncAI64IRI64E(player.Seek),
			},
		},
	}, nil
}

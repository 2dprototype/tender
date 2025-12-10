package stdlib

import (
	"github.com/2dprototype/tender"
)

// BuiltinModules are builtin type standard library modules.
var BuiltinModules = map[string]map[string]tender.Object{
	"math":         mathModule,
	"cmplx":        cmplxModule,
	"os":           osModule,
	"strings":      stringsModule,
	"times":        timesModule,
	"rand":         randModule,
	"fmt":          fmtModule,
	"json":         jsonModule,
	"base64":       base64Module,
	"hex":          hexModule,
	"colors":       colorsModule,
	"gzip":         gzipModule,
	"zip":          zipModule,
	"tar":          tarModule,
	"bufio":        bufioModule,
	"crypto":       cryptoModule,
	"path":         pathModule,
	"image":        imageModule,
	"canvas":       canvasModule,
	"dll":          dllModule,
	"io":           ioModule,
	"audio":        audioModule,
	"net":          netModule,
	"http":         httpModule,
	"websocket":    websocketModule,
	"gob":          gobModule,
	"csv":          csvModule,
	"xml":          xmlModule,
	// "wui":          wuiModule,
}

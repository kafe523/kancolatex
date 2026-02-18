package main

//#include <stdlib.h>
//#include <stdint.h>
import "C"
import (
	"fmt"
	"os"
	"runtime/cgo"
	"strings"

	"github.com/kafe523/kancolatex/internal"
	"github.com/kafe523/kancolatex/internal/access"
	"github.com/kafe523/kancolatex/internal/macro"
)

//export InitMacroContext
func InitMacroContext(cNoroPath *C.char) C.uintptr_t {
	var mc macro.MacroContext

	noroPath := strings.TrimSpace(C.GoString(cNoroPath))

	var noroRaw []byte

	file, err := os.ReadFile(noroPath)
	if err != nil {
		panic(err)
	}
	noroRaw = file

	err = mc.Populate(&noroRaw)
	if err != nil {
		panic(err)
	}

	h := cgo.NewHandle(mc)

	return C.uintptr_t(h)
}

func generalAccess(mcp C.uintptr_t, argPat *C.char, mode string) *C.char {
	h := cgo.Handle(mcp)
	val := h.Value()

	mc, ok := val.(macro.MacroContext)
	if !ok {
		fmt.Fprintln(os.Stderr, "mcp value is not MacroContext")
		return C.CString("")
	}

	var pattern string

	switch mode {
	case "MACRO":
		pattern = mc.GetMacroPattern(C.GoString(argPat))
	case "PATTERN":
		pattern = C.GoString(argPat)
	default:
		fmt.Fprintln(os.Stderr, "unknown mode", mode)
		return C.CString("")
	}

	result := access.Access(&mc, access.PatternTokenizer(pattern))

	if internal.IsInterfaceNil(result) {
		result = ""
	}

	return C.CString(fmt.Sprintf("%v", result))
}

//export AccessMacro
func AccessMacro(mcp C.uintptr_t, mac *C.char) *C.char {
	return generalAccess(mcp, mac, "MACRO")
}

//export AccessPattern
func AccessPattern(mcp C.uintptr_t, pat *C.char) *C.char {
	return generalAccess(mcp, pat, "PATTERN")
}

//export AccessKeyStr
func AccessKeyStr(mcp C.uintptr_t) *C.char {
	h := cgo.Handle(mcp)
	val := h.Value()

	mc, ok := val.(macro.MacroContext)
	if !ok {
		fmt.Fprintln(os.Stderr, "mcp value is not MacroContext")
		return C.CString("")
	}

	return C.CString(mc.AccessKeyStr())
}

func main() {}

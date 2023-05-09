package AnsiTerm

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"golang.org/x/term"
)

const (
	HOME        = "\033[H"
	CLRSCR      = "\033[2J"
	ERASE_LINE  = "\033[K"
	ERASE_EOL   = "\033[0K"
	ERASE_SOL   = "\033[1K"
	RESET       = "\033[0m"
	RESET_COLOR = "\033[32m"
	RESET_LINE  = "\r\033[2K"
)

var (
	ThinHLine    = "─"
	FrameHLine   = "━"
	FrameVLine   = "┃"
	FrameOpenR   = "┓"
	FrameTLineR  = "┫"
	FrameCloseR  = "┛"
	FrameOpenL   = "┏"
	FrameTLineL  = "┣"
	FrameCloseL  = "┗"
	FrameOHLine  = "═"
	FrameOVLine  = "║"
	FrameOOpenR  = "╗"
	FrameOTLineR = "╣"
	FrameOCloseR = "╝"
	FrameOOpenL  = "╔"
	FrameOTLineL = "╠"
	FrameOCloseL = "╚"
	Harrow       = "⋙"
	Rarrow       = "⋙"
	Larrow       = "⋘"
	BulletChar   = "•"
	MarkChar     = "★"
	ContStr      = "…"
)

func AsciiChars() {
	ThinHLine = "-"
	FrameHLine = "-"
	FrameVLine = "|"
	FrameOpenR = "+"
	FrameTLineR = "+"
	FrameCloseR = "+"
	FrameOpenL = "+"
	FrameTLineL = "+"
	FrameCloseL = "+"
	FrameOHLine = "="
	FrameOVLine = "|"
	FrameOOpenR = "+"
	FrameOTLineR = "+"
	FrameOCloseR = "+"
	FrameOOpenL = "+"
	FrameOTLineL = "+"
	FrameOCloseL = "+"
	Harrow = ">>>"
	Rarrow = ">>>"
	Larrow = "<<<"
	BulletChar = "*"
	MarkChar = "*"
	ContStr = "..."

	color.NoColor = true
}

var Out *bufio.Writer = bufio.NewWriter(os.Stdout)

//
// Terminal handling: ----------------------------------------------
//

// IsTTY checks for interactive terminal
func IsTTY() bool {
	return term.IsTerminal(int(os.Stdout.Fd()))
}

// HasColor checks color capabilities
func HasColor() bool {
	return os.Getenv("TERM") != "dumb" && IsTTY()
}

// GetSize retruen the current terminal size
func GetSize() (int, int, error) {
	return term.GetSize(int(os.Stdout.Fd()))
}

func IsColorTerm() bool {
	noColor := false

	if HasColor() {
		noColor = false
	} else {
		noColor = true
	}

	if !IsTTY() {
		noColor = true
	}

	return noColor
}

//
// Screen management: ------------------------------------------------
//

// ClearScr clears screen
func ClearScr() {
	Out.WriteString(HOME)
	Out.Flush()
}

// Reset screen
func Reset() {
	Out.WriteString(RESET)
	Out.Flush()
}

//
// Cursor movement: ------------------------------------------------
//

// CursorPos moves cursor to position (x,y)
func CursorPos(x int, y int) {
	Out.WriteString(fmt.Sprintf("\033[%d;%dH", y, x))
	Out.Flush()
}

// CursorUp moves cursor up relative the current position
func CursorUp(count int) {
	Out.WriteString(fmt.Sprintf("\033[%dA", count))
	Out.Flush()
}

// CursorDown moves cursor down relative the current position
func CursorDown(count int) {
	Out.WriteString(fmt.Sprintf("\033[%dB", count))
	Out.Flush()
}

// CursorRight moves cursor forward relative the current position
func CursorRight(count int) {
	Out.WriteString(fmt.Sprintf("\033[%dC", count))
	Out.Flush()
}

// CursorLeft moves cursor backward relative the current position
func CursorLeft(count int) {
	Out.WriteString(fmt.Sprintf("\033[%dD", count))
	Out.Flush()
}

// StartOfLine returns cursor to start of line
func StartOfLine() {
	Out.WriteString(RESET_LINE)
	Out.Flush()
}

//
// Color functions: ------------------------------------------------
//

var Normal = color.New(color.Reset).SprintFunc()

var Bold = color.New(color.Bold).SprintFunc()
var Faint = color.New(color.Faint).SprintFunc()
var Italic = color.New(color.Italic).SprintFunc()
var Underline = color.New(color.Underline).SprintFunc()
var Strike = color.New(color.CrossedOut).SprintFunc()
var Blink = color.New(color.BlinkSlow).SprintFunc()

var Red = color.New(color.FgRed).SprintFunc()
var Green = color.New(color.FgGreen).SprintFunc()
var Yellow = color.New(color.FgYellow).SprintFunc()
var Blue = color.New(color.FgBlue).SprintFunc()
var Magenta = color.New(color.FgMagenta).SprintFunc()
var Cyan = color.New(color.FgCyan).SprintFunc()
var White = color.New(color.FgWhite).SprintFunc()

var HiRed = color.New(color.FgHiRed).SprintFunc()
var HiGreen = color.New(color.FgHiGreen).SprintFunc()
var HiYellow = color.New(color.FgHiYellow).SprintFunc()
var HiBlue = color.New(color.FgHiBlue).SprintFunc()
var HiMagenta = color.New(color.FgMagenta).SprintFunc()
var HiCyan = color.New(color.FgHiCyan).SprintFunc()
var HiWhite = color.New(color.FgHiWhite).SprintFunc()

//
// Little Printer: ------------------------------------------------
//

func hline(lchar string, n int) string {
	if n == 0 {
		n, _, _ = GetSize()
	}

	return strings.Repeat(lchar, n)
}

// Printer is a little shortcut for verbose and debug output
type Printer struct {
	flagVerbose, flagDebug, flagSilent bool
}

func NewPrinter() *Printer {
	var p Printer

	// Detect locale for printing:
	if runtime.GOOS == "windows" {
		AsciiChars()
	}

	return &p
}

// Management functions for Printer
func (l *Printer) SetDebug(b bool) {
	l.flagDebug = b
}

func (l *Printer) SetVerbose(b bool) {
	l.flagVerbose = b
}

func (l *Printer) SetSilent(b bool) {
	l.flagSilent = b
}

// Helper functions for Printer
func (l Printer) Frame(str string) string {
	sl := len(str)
	rh := FrameOpenL + strings.Repeat(FrameHLine, sl+2) + FrameOpenR
	rt := FrameCloseL + strings.Repeat(FrameHLine, sl+2) + FrameCloseR
	return fmt.Sprintf("%s\n%s %s %s\n%s\n", rh, FrameVLine, str, FrameVLine, rt)
}

func (l Printer) OFrame(str string) string {
	sl := len(str)
	rh := FrameOOpenL + strings.Repeat(FrameOHLine, sl+2) + FrameOOpenR
	rt := FrameOCloseL + strings.Repeat(FrameOHLine, sl+2) + FrameOCloseR
	return fmt.Sprintf("\n%s\n%s %s %s\n%s\n", rh, FrameOVLine, str, FrameOVLine, rt)
}

func (l Printer) Underlines(row []string) []string {
	anonRow := make([]string, len(row))
	for i, v := range row {
		anonRow[i] = strings.Repeat(FrameHLine, len(v))
	}
	return anonRow
}

func (l Printer) WriteOut(fmtString string, args ...interface{}) {
	if !l.flagSilent {
		fmt.Printf(fmtString, args...)
	}
}

func (l Printer) WriteAny(fmtString string, args ...interface{}) {
	fmt.Printf(fmtString, args...)
}

// Print functions for logger
func (l Printer) Banner(fmtString string, args ...interface{}) {
	rStr := fmt.Sprintf(fmtString, args...)
	str := l.Frame(rStr)
	l.WriteOut(Bold(Green(str)))
}

// Print functions for logger
func (l Printer) OBanner(fmtString string, args ...interface{}) {
	rStr := fmt.Sprintf(fmtString, args...)
	str := l.OFrame(rStr)
	l.WriteOut(Bold(Green(str)))
}

func (p Printer) ModuleHeading(subPage bool, modName, fmtString string, args ...interface{}) {
	var eStr string
	if subPage {
		fmt.Println("\n" + hline(ThinHLine, 80))
		eStr = fmt.Sprintf("\nModule %-10s   %s\n", modName+":", fmt.Sprintf(fmtString, args...))
	} else {
		eStr = fmt.Sprintf("\nUsage:   %s\n", fmt.Sprintf(fmtString, args...))
	}
	p.WriteAny(Yellow(eStr))
}

func (l Printer) Verbose(fmtString string, args ...interface{}) {
	if l.flagVerbose {
		l.WriteOut(fmtString, args...)
	}
}

func (l Printer) Verboseln(fmtString string, args ...interface{}) {
	l.Verbose(fmtString+"\n", args...)
}

func (l Printer) VerboseInfo(fmtString string, args ...interface{}) {
	if l.flagVerbose {
		l.WriteOut(Green(fmtString), args...)
	}
}

func (l Printer) VerboseInfoln(fmtString string, args ...interface{}) {
	l.VerboseInfo(fmtString+"\n", args...)
}

func (l Printer) VerboseBold(fmtString string, args ...interface{}) {
	if l.flagVerbose {
		l.WriteOut(Bold(fmtString), args...)
	}
}

func (l Printer) VerboseBoldln(fmtString string, args ...interface{}) {
	l.VerboseBold(fmtString+"\n", args...)
}

func (l Printer) Debug(fmtString string, args ...interface{}) {
	if l.flagDebug {
		fs := "*** DEB: " + fmtString
		l.WriteOut(Red(fs), args...)
	}
}

func (l Printer) Debugln(fmtString string, args ...interface{}) {
	l.Debug(fmtString+"\n", args...)
}

func (l Printer) Warning(fmtString string, args ...interface{}) {
	fs := "*** WARN: " + fmtString
	l.WriteAny(Yellow(fs), args...)
}

func (l Printer) Warningln(fmtString string, args ...interface{}) {
	l.Warning(fmtString+"\n", args...)
}

func (l Printer) Error(fmtString string, args ...interface{}) {
	fs := "*** ERR: " + fmt.Sprintf(fmtString, args...)
	l.WriteAny(Red(fs))
}

func (l Printer) Errorln(fmtString string, args ...interface{}) {
	l.Error(fmtString+"\n", args...)
}

package AnsiTerm

import (
	"bufio"
	"fmt"
	"os"

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
	Harrow       = "⮕"
	HRarrow      = "⮕"
	HLarrow      = "⬅"
	Rarrow       = "⋙"
	Larrow       = "⋘"
	BulletChar   = "•"
	MarkChar     = "★"
	ContStr      = "…"
	PromptChar   = "»"
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
	PromptChar = ">"

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

// GetSize returns the current terminal size (width int, height int, err error)
func GetSize() (int, int, error) {
	return term.GetSize(int(os.Stdout.Fd()))
}

func IsColorTerm() bool {
	return HasColor() && IsTTY()
}

func NoColor() bool {
	return !IsColorTerm()
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

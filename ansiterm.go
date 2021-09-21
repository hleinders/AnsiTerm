package AnsiTerm

import (
	"bufio"
	"fmt"
	"os"
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

var Out *bufio.Writer = bufio.NewWriter(os.Stdout)

// ClearScr clears screen
func ClearScr() {
	Out.WriteString(HOME)
	Out.Flush()
}

// Reset clears screen
func Reset() {
	Out.WriteString(RESET)
	Out.Flush()
}

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

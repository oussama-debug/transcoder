package color

import (
	"fmt"
	"strconv"
)

const _colorFormat = "\x1b[%dm%s\x1b[0m"

// AINSI terminal
// color code
type Color int

// basic AINSI colors
const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	LightGray
	DarkGray = 90

	Bold = 1
)

/**
* {s} string - message string to colorize
* {c} Color or int - message color terminal AINSI code
* returns string - ASCII colored string based on given color.
 */
func colorizeMessage(s string, c Color) string {
	if c == -1 {
		return s
	}
	return fmt.Sprintf(_colorFormat, c, s)
}

/**
* {s} string - message string to colorize
* {c} int - message color terminal AINSI code
* returns string - ASCII colored string based on given color.
 */
func colorfulMessage(text string, color int) string {
	return "\033[38;5;" + strconv.Itoa(color) + "m" + text + "\033[0m"
}

/**
* {bb} byte - the bytes array
* {ii} int - the given indices to highlight
* {c}  int - the color for that specific array of indices
* returns [] byte - the byte array back with highlights
 */
func hightlightMessage(bb []byte, ii []int, c int) []byte {
	b := make([]byte, 0, len(bb))
	for i, j := 0, 0; i < len(bb); i++ {
		if j < len(ii) && ii[j] == i {
			b = append(b, colorfulByte(bb[i], c)...)
			j++
		} else {
			b = append(b, bb[i])
		}
	}
	return b
}

func colorfulByte(b byte, color int) []byte {
	return []byte(colorfulMessage(string(b), color))
}

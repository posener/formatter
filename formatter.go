// formatter is a library for formatting text.
//
// The `Formatter` wraps an `io.Writer` and formats text written to it. It enables text indenting
// and line width wrapping.
package formatter

import (
	"io"
)

// Formatter formats text that is written to it and write the formatted text to a given `io.Writer`.
type Formatter struct {
	// Writer to write the formatted text.
	io.Writer
	// Indent will be prefixed to every new line. An empty slice will be ignored.
	Indent []byte
	// Width defines the line size for wrapping the text. A non positive value will be ignored.
	Width int

	// curWidth stores the width of the current line.
	curWidth int
}

func (f *Formatter) Write(b []byte) (int, error) {
	return f.Writer.Write(f.format(b))
}

func (f *Formatter) format(b []byte) []byte {
	// lastSpace stores the last location of a spotted space charcter in order to insert a line
	// break. A negative value is set to identify that a space was not spotted since the last
	// inserted line break.
	lastSpace := -1

	for i := 0; i < len(b); i++ {
		// Insert indentation if a new line.
		if len(f.Indent) > 0 && f.curWidth == 0 {
			i, b = insert(b, i, f.Indent)
			f.curWidth = len(f.Indent) + 1
		} else {
			f.curWidth++
		}

		switch b[i] {
		case '\n':
			f.curWidth = 0
		case ' ', '\t':
			lastSpace = i
		default:
			if f.Width > 0 && f.curWidth > f.Width {
				if lastSpace >= 0 {
					b[lastSpace] = '\n'
					i = lastSpace - 1 // start next loop from the new line.
					lastSpace = -1
				}
			}
		}
	}
	return b
}

// insert inserts `in` into `buf` at location `i`. It returns the index of the byte after the
// inserted bytes and the buffer containing the buffer with the inserted bytes.
func insert(buf []byte, i int, in []byte) (int, []byte) {
	before, after := buf[:i], buf[i:]
	buf = append(before, append(in, after...)...)
	return i + len(in), buf
}

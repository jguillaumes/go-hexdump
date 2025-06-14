// hexdump provides a function to generate hexadecimal dumps in mainframe
// style, applying a codepage conversion to show the plain text.
package hexdump

import (
	"fmt"
	"strings"
	"unicode"

	e "github.com/jguillaumes/go-encoding/encodings"
)

// HexDump generates an hexadecimal dump in "mainframe style", applying a given
// emcoding using [github.com/jguillaumes/go-encoding/encodings].
func HexDump(data []byte, codepage string) string {
	const colheader = "....|....1....|....2....|....3....|....4....|....5....|....6...."
	const fiveblanks = "     "

	enc := e.NewEncoding()

	// Encode the byte slice into an uppercase hex string
	hexString := make([]byte, len(data)*2)
	for i, b := range data {
		hexString[i*2] = "0123456789ABCDEF"[b>>4]
		hexString[i*2+1] = "0123456789ABCDEF"[b&0x0F]
	}

	// Split the hex string into lines of 64 bytes each
	var lines []string
	for i := 0; i < len(hexString); i += 128 {
		end := i + 128
		if end > len(hexString) {
			end = len(hexString)
		}
		lines = append(lines, string(hexString[i:end]))
	}

	// Split the data into blocks of 64 bytes
	// For each block, build a string of printable characters
	// using EBCDIC encoding
	var printableLines []string
	for i := 0; i < len(data); i += 64 {
		end := i + 64
		if end > len(data) {
			end = len(data)
		}
		block := data[i:end]
		converted, _ := enc.DecodeBytes(block, codepage)
		printable := strings.Map(func(r rune) rune {
			if unicode.IsPrint(r) {
				return r
			} else {
				return '.'
			}
		}, converted)
		// Check every rune to see if it is printable

		printableLines = append(printableLines, printable)
	}

	// For each line, get the high and low nibbles (odd and even indices)
	var highNibbles, lowNibbles []string
	for _, line := range lines {
		highNibble := make([]byte, 0, 64)
		lowNibble := make([]byte, 0, 64)
		for j := 0; j < len(line); j += 2 {
			if j < len(line) {
				highNibble = append(highNibble, line[j])
			}
			if j+1 < len(line) {
				lowNibble = append(lowNibble, line[j+1])
			}
		}
		highNibbles = append(highNibbles, string(highNibble))
		lowNibbles = append(lowNibbles, string(lowNibble))
	}

	// Build the final output string
	var output string
	output = fiveblanks + colheader + "\n"
	offset := 0
	for i := 0; i < len(highNibbles); i++ {
		output = fmt.Sprintf("%s\n%s%s\n%04x %s\n%s%s\n", output, fiveblanks,
			printableLines[i], offset, highNibbles[i], fiveblanks, lowNibbles[i])
		offset += 64
	}
	return output

}

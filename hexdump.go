package hexdump

import (
	"fmt"

	ebcdic "github.com/jguillaumes/go-ebcdic"
)

func HexDump(data []byte) string {
	const colheader = "....|....1....|....2....|....3....|....4....|....5....|....6...."

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
		printable, _ := ebcdic.Decode(block, ebcdic.EBCDIC037)
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
	output = colheader + "\n"
	for i := 0; i < len(highNibbles); i++ {
		output = fmt.Sprintf("%s\n%s\n%s\n%s\n", output, printableLines[i], highNibbles[i], lowNibbles[i])
	}
	return output

}

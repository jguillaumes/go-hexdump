package hexdump

import (
	"testing"

	e "github.com/jguillaumes/go-encoding/encodings"
)

var enc = e.NewEncoding()

func Test_hexdump(t *testing.T) {

	exp1 := `     ....|....1....|....2....|....3....|....4....|....5....|....6....

     Hello, World! This is a test of the hexdump function. 1234567890
0000 C899964E999854E88A48A484A8AA4984A88488A8A9948A98A89944FFFFFFFFFF
     85336B066934A038920920103523066038508574447064533965B01234567890

      ABCDEFGHIJKLMNOPQRSTUVWXYZ abcdefghijklmnopqrstuvwxyz
0040 4CCCCCCCCCDDDDDDDDDEEEEEEEE4888888888999999999AAAAAAAA
     012345678912345678923456789012345678912345678923456789
`

	exp2 := `     ....|....1....|....2....|....3....|....4....|....5....|....6....

     1234567890123456789012345678901234567890123456789012345678901234
0000 FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF
     1234567890123456789012345678901234567890123456789012345678901234

     5678901234567890123456789012345678901234567890
0040 FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF
     5678901234567890123456789012345678901234567890
`

	var data = "Hello, World! This is a test of the hexdump function. 1234567890 ABCDEFGHIJKLMNOPQRSTUVWXYZ abcdefghijklmnopqrstuvwxyz"
	var data_ebcdic, _ = enc.EncodeString(data, "IBM-037")

	result := HexDump(data_ebcdic, "IBM-037")
	println(result)
	if result != exp1 {
		t.Errorf("HexDump() = %v, want %v", result, exp1)
	}

	data = "12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"
	data_ebcdic, _ = enc.EncodeString(data, "IBM-037")

	result = HexDump(data_ebcdic, "IBM-037")
	println(result)
	if result != exp2 {
		t.Errorf("HexDump() = %v, want %v", result, exp2)
	}

}

func Test_cp1047(t *testing.T) {
	var data = "ABCDEFGHIJKLMNOPQRSTUVWXYZ abcdefghijklmnopqrstuvwxyz 1234567890"
	var data_ebcdic, _ = enc.EncodeString(data, "IBM-1047")

	result := HexDump(data_ebcdic, "IBM-037")
	println(result)
}

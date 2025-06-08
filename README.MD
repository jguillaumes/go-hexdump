# go-hexdump: Mainframe style hexadecimal dumps

This package adds a function to create mainframe style hexadecimal dumps for byte slices, applying a given encoding translation.

This is an example of such a dump:

```
     ....|....1....|....2....|....3....|....4....|....5....|....6....

     Hello, World! This is a test of the hexdump function. 1234567890
0000 C899964E999854E88A48A484A8AA4984A88488A8A9948A98A89944FFFFFFFFFF
     85336B066934A038920920103523066038508574447064533965B01234567890

      ABCDEFGHIJKLMNOPQRSTUVWXYZ abcdefghijklmnopqrstuvwxyz
0040 4CCCCCCCCCDDDDDDDDDEEEEEEEE4888888888999999999AAAAAAAA
     012345678912345678923456789012345678912345678923456789
```

In this case, the IBM-037 codepage has been used. Actually, this dump comes from the go unit test program.

## Using this package

Just import it into your golang project as

`github.com/jguillaumes/go-hexdump/hexdump`

The package exports a single function:

```
func HexDump(data []byte, codepage string) string
```

Ths codepage name must correspond to one supported by my [go-encoding](https://github.com/jguillaumes/go-encoding) package. For instance, `IBM-037` or `IBM-1047`. Se the go-encoding docs for details.

The hexdump text is returned as a string, with embedded newline markers, so it can be output by simply passing it to fmt.Print() or log.Print().

## Example

Coming straight from the test case:

```golang
func Test_cp1047(t *testing.T) {
	var data = "ABCDEFGHIJKLMNOPQRSTUVWXYZ abcdefghijklmnopqrstuvwxyz 1234567890"
	var data_ebcdic, _ = enc.EncodeString(data, "IBM-1047")

	result := HexDump(data_ebcdic, "IBM-037")
	println(result)
}
```
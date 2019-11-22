package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(byte(0x01), byte(0x02), byte(0x03), byte(0x2), byte(30), byte(0x05<<4|byte(2)))
	buf := &bytes.Buffer{}

	buf.WriteByte(0x05<<4 | byte(2))
	buf.WriteByte(0x80 | (0x01 << 4) | 0x00)
	buf.WriteByte(0x00)

	fmt.Println(buf.Bytes())
}

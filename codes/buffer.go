package codes

import (
	"bytes"
	"fmt"
)

func WriteBuffer() {
	var buf bytes.Buffer

	buf.Write([]byte("an old"))
	buf.WriteByte(32)
	buf.WriteString("cactus")
	buf.WriteByte(32)
	buf.WriteByte(32)
	buf.WriteRune('🌵')

	fmt.Println(buf)
	fmt.Println(buf.String())
}

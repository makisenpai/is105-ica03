package hex

import (
	"encoding/hex"
	"fmt"
)

func HexReturnBytes(b []byte) []byte {

	t := b
	d := make([]byte, hex.EncodedLen(len(t)))
	hex.Encode(d, t)
	return d

}

func HexReturn(s string) string {

	t := []byte(s)
	d := make([]byte, hex.EncodedLen(len(t)))
	hex.Encode(d, t)
	return fmt.Sprintf("%s", d)

}

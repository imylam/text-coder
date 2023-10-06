package textcoder

import (
	"encoding/hex"
)

type HexCoder struct{}

func (e *HexCoder) Decode(input string) (output []byte, err error) {
	return hex.DecodeString(input)
}

func (e *HexCoder) Encode(input []byte) string {
	return hex.EncodeToString(input)
}

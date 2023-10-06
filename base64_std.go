package textcoder

import "encoding/base64"

type Base64StdCoder struct{}

func (e *Base64StdCoder) Decode(input string) (output []byte, err error) {
	return base64.StdEncoding.DecodeString(input)
}

func (e *Base64StdCoder) Encode(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

package textcoder

import "encoding/base64"

type Base64RawStdCoder struct{}

func (e *Base64RawStdCoder) Decode(input string) (output []byte, err error) {
	return base64.RawStdEncoding.DecodeString(input)
}

func (e *Base64RawStdCoder) Encode(input []byte) string {
	return base64.RawStdEncoding.EncodeToString(input)
}

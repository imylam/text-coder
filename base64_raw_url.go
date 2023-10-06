package textcoder

import "encoding/base64"

type Base64RawUrlCoder struct{}

func (e *Base64RawUrlCoder) Decode(input string) (output []byte, err error) {
	return base64.RawURLEncoding.DecodeString(input)
}

func (e *Base64RawUrlCoder) Encode(input []byte) string {
	return base64.RawURLEncoding.EncodeToString(input)
}

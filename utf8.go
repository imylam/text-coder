package textcoder

type Utf8Coder struct{}

func (e *Utf8Coder) Encode(input []byte) string {
	return string(input)
}

func (e *Utf8Coder) Decode(input string) (output []byte, err error) {
	return []byte(input), nil
}

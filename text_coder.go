package textcoder

type Coder interface {
	Encode([]byte) string
	Decode(string) ([]byte, error)
}

package textcoder

import (
	"bytes"
	cryptoRand "crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"testing"
)

func Test_Decoding(t *testing.T) {

	b64Coder := &Base64StdCoder{}
	b64RawStdCoder := &Base64RawStdCoder{}
	b64RawUrlCoder := &Base64RawUrlCoder{}
	hexCoder := &HexCoder{}
	utf8Coder := &Utf8Coder{}

	randomInput := randomBytes()

	b64Input := base64.StdEncoding.EncodeToString(randomInput)
	b64RawStdInput := base64.RawStdEncoding.EncodeToString(randomInput)
	b64RawUrlInput := base64.RawURLEncoding.EncodeToString(randomInput)
	hexInput := hex.EncodeToString(randomInput)
	utf8Input := string(randomInput)

	type args struct {
		s string
	}

	testCases := []struct {
		name  string
		coder Coder
		args  args
		want  int
	}{
		{name: "test-b64std", coder: b64Coder, args: args{s: b64Input}, want: 0},
		{name: "test-b64RawStd", coder: b64RawStdCoder, args: args{s: b64RawStdInput}, want: 0},
		{name: "test-b64RawUrl", coder: b64RawUrlCoder, args: args{s: b64RawUrlInput}, want: 0},
		{name: "test-hex", coder: hexCoder, args: args{s: hexInput}, want: 0},
		{name: "test-utf8", coder: utf8Coder, args: args{s: utf8Input}, want: 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, _ := tc.coder.Decode(tc.args.s)

			isSame := bytes.Compare(result, randomInput)

			if isSame != tc.want {
				t.Fail()
			}
		})
	}
}

func Test_Encoding(t *testing.T) {
	b64Coder := &Base64StdCoder{}
	b64RawStdCoder := &Base64RawStdCoder{}
	b64RawUrlCoder := &Base64RawUrlCoder{}
	hexCoder := &HexCoder{}
	utf8Coder := &Utf8Coder{}

	randomInput := randomBytes()

	b64Input := base64.StdEncoding.EncodeToString(randomInput)
	b64RawStdInput := base64.RawStdEncoding.EncodeToString(randomInput)
	b64RawUrlInput := base64.RawURLEncoding.EncodeToString(randomInput)
	hexInput := hex.EncodeToString(randomInput)
	utf8Input := string(randomInput)

	testCases := []struct {
		name  string
		coder Coder
		want  string
	}{
		{name: "test-b64std", coder: b64Coder, want: b64Input},
		{name: "test-b64RawStd", coder: b64RawStdCoder, want: b64RawStdInput},
		{name: "test-b64RawUrl", coder: b64RawUrlCoder, want: b64RawUrlInput},
		{name: "test-hex", coder: hexCoder, want: hexInput},
		{name: "test-utf8", coder: utf8Coder, want: utf8Input},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.coder.Encode(randomInput)

			if result != tc.want {
				t.Errorf("Failed, expected: %s, got: %s", tc.want, result)
			}
		})
	}
}

func randomBytes() []byte {
	num := rand.Intn(100-10) + 10
	randomInput := make([]byte, num)
	cryptoRand.Read(randomInput)

	return randomInput
}

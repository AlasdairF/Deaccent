package deaccent

import (
    "unicode"
	"unicode/utf8"
    "golang.org/x/text/transform"
    "golang.org/x/text/unicode/norm"
	"io"
)

func isMn(r rune) bool {
    return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func NewReader(r io.Reader) io.Reader {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	return transform.NewReader(r, t)
}

func Bytes(b []byte) ([]byte, error) {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	res, _, err := transform.Bytes(t, b)
	return res, err
}

func Rune(r rune) rune {
	b := make([]byte, 3)
	n := utf8.EncodeRune(b, r)
	b, _ = Bytes(b[0:n])
	newrune, _ := utf8.DecodeRune(b)
	return newrune
}

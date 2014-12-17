package deaccent

import (
    "unicode"
    "golang.org/x/text/transform"
    "golang.org/x/text/unicode/norm"
	"io"
)

isMn := func(r rune) bool {
    return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func NewReader(r io.Reader) io.Reader {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	return transform.NewReader(r, t)
}

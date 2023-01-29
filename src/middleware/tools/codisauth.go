package tools

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

func NewXAuth(segs ...string) string {
	t := &bytes.Buffer{}
	fmt.Fprintf(t, "Codis-XAuth")
	for _, s := range segs {
		fmt.Fprintf(t, "-[%s]", s)
	}
	b := sha256.Sum256(t.Bytes())
	return fmt.Sprintf("%x", b[:16])
}

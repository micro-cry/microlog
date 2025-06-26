package microlog

import (
	"encoding/hex"
	"golang.org/x/crypto/blake2b"
	"testing"
)

func TestName(t *testing.T) {
	h, _ := blake2b.New(4, hashKey)
	h.Write([]byte("hello world"))

	b := h.Sum(nil)
	s := hex.EncodeToString(b)

	t.Log(b, len(b), "\t", s, len(s))
}

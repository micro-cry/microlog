package microlog

import (
	"encoding/hex"
	"github.com/micro-cry/microlog/target"
	"golang.org/x/crypto/blake2b"
)

// // // // // // // // // // // // // // // // // //

func hash(text string) []byte {
	h, _ := blake2b.New(16, hashKey)
	h.Write([]byte(text))
	return h.Sum(nil)
}

// // // //

type HashType []byte

func (hs HashType) String() string {
	return hex.EncodeToString(hs)
}

var hashKey []byte

func init() {
	hashKey = hash(target.GlobalName)
}

// // // //

func HashPointer(data []byte) HashType {
	h, _ := blake2b.New(16, hashKey)
	h.Write(data)
	return h.Sum(nil)
}

func HashLog(data []byte) HashType {
	h, _ := blake2b.New(32, hashKey)
	h.Write(data)
	return h.Sum(nil)
}

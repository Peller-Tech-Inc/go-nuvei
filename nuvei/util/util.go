package util

import (
	"crypto"
	"fmt"
)

func SHA256(toConvert string) string {
	hash := crypto.SHA256.New()
	hash.Write([]byte(toConvert))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

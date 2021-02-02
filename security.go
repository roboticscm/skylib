package skylib

import (
	"crypto/sha1"
	"fmt"
	"io"
)

//EncodeSHA1Password encode password with key
func EncodeSHA1Password(password string, key string) string {
	if key == "" {
		key = "Skyhub@010116"
	}
	h := sha1.New()
	io.WriteString(h, key+password)

	return fmt.Sprintf("%x", h.Sum(nil))
}

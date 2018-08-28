package data

import (
	"crypto/sha256"
	"fmt"
	"io"
)

// User is sql users table
type User struct {
	ID    int
	Name  string
	PW    string
	Token string
}

// MakeHash creats sha256
func MakeHash(str string) (hs string) {
	sha := sha256.New()
	io.WriteString(sha, str+"abcdefghijklmn")
	hs = fmt.Sprintf("%x", sha.Sum(nil))
	return
}

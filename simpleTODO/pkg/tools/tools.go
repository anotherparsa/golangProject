package tools

import (
	"crypto"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func GenerateUUID() string {
	bytes := make([]byte, 33)
	_, err := rand.Read(bytes)
	if err != nil {
		fmt.Println(err)
	}
	return base64.RawStdEncoding.EncodeToString(bytes)

}

func HashThis(originalText string) string {
	hash := crypto.SHA256.New()
	hash.Write([]byte(originalText))
	hashed_byte := hash.Sum(nil)
	hashedPassword := hex.EncodeToString(hashed_byte)
	return hashedPassword
}

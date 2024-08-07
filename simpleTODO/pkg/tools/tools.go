package tools

import (
	"crypto/rand"
	"encoding/base64"
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

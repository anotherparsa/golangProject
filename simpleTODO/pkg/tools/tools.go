package tools

import (
	"crypto"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
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

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "session_id", MaxAge: -1})
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

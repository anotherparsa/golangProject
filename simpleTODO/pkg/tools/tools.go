package tools

import (
	"crypto"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"regexp"
)

//other tools
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

func ValidateSignupFormInputs(username string, password string, firstname string, lastname string, email string, phoneNumber string) bool {
	validationFlag := true
	if username == "" {
		fmt.Println("Empty username")
		validationFlag = false
		return validationFlag
	} else {
		if len(username) < 5 || len(username) > 30 {
			fmt.Println("invalid length for username")
			validationFlag = false
			return validationFlag
		} else {
			if !(regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(username)) {
				fmt.Println("username must only contain alpha numeric characters")
				validationFlag = false
				return validationFlag
			}
		}
	}

	if password == "" {
		fmt.Println("Empty Password")
		validationFlag = false
		return validationFlag
	} else {
		if len(password) < 5 {
			fmt.Println("invalid length for username")
			validationFlag = false
			return validationFlag
		} else {
			if !(regexp.MustCompile(`[a-zA-Z]`).MatchString(password) && regexp.MustCompile(`\d`).MatchString(password) && regexp.MustCompile(`[\W_]`).MatchString(password)) {
				fmt.Println("Password must be alpha numeric and at least a symbol character")
				validationFlag = false
				return validationFlag
			}
		}
	}

	if firstname == "" {
		fmt.Println("Empty First name")
		validationFlag = false
		return validationFlag
	} else {
		if len(firstname) < 3 || len(firstname) > 20 {
			fmt.Println("invalid length for First name")
			validationFlag = false
			return validationFlag
		} else {
			if !(regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(firstname)) {
				fmt.Println("firstname must only contain alpha numeric characters")
				validationFlag = false
				return validationFlag
			}
		}
	}

	if lastname == "" {
		fmt.Println("Empty Last name")
		validationFlag = false
		return validationFlag
	} else {
		if len(lastname) < 3 || len(lastname) > 20 {
			fmt.Println("invalid length for First name")
			validationFlag = false
			return validationFlag
		} else {
			if !(regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(lastname)) {
				fmt.Println("firstname must only contain alpha numeric characters")
				validationFlag = false
				return validationFlag
			}
		}
	}

	if email == "" {
		fmt.Println("Empty email")
		validationFlag = false
		return validationFlag
	} else {
		if len(email) > 40 {
			fmt.Println("invalid length for email")
			validationFlag = false
			return validationFlag
		} else {
			if !(regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(email)) {
				fmt.Println("invalid email")
				validationFlag = false
				return validationFlag
			}
		}
	}

	if phoneNumber == "" {
		fmt.Println("Empty phone number")
		validationFlag = false
		return validationFlag
	} else {
		if len(phoneNumber) != 10 {
			fmt.Println("Invalid length for phone number")
			validationFlag = false
			return validationFlag
		} else {
			if !(regexp.MustCompile(`(^\d+$)`).MatchString(phoneNumber)) {
				fmt.Println("Phone number must be only numeric")
				validationFlag = false
				return validationFlag
			}
		}
	}

	return validationFlag
}

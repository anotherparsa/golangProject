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

func ValidateUserInfoFormInputs(tobevalidated string, valuetobevalidated string) bool {
	validationFlag := true

	if tobevalidated == "username" {
		if valuetobevalidated == "" {
			validationFlag = false
			return validationFlag
		} else {
			if len(valuetobevalidated) < 5 || len(valuetobevalidated) > 30 {
				validationFlag = false
				return validationFlag
			} else {
				if !(regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(valuetobevalidated)) {
					validationFlag = false
					return validationFlag
				}
			}
		}
	} else if tobevalidated == "password" {
		if tobevalidated == "" {
			validationFlag = false
			return validationFlag
		} else {
			if len(valuetobevalidated) < 5 {
				validationFlag = false
				return validationFlag
			} else {
				if !(regexp.MustCompile(`[a-zA-Z]`).MatchString(valuetobevalidated) && regexp.MustCompile(`\d`).MatchString(valuetobevalidated) && regexp.MustCompile(`[\W_]`).MatchString(valuetobevalidated)) {
					validationFlag = false
					return validationFlag
				}
			}
		}
	} else if tobevalidated == "firstname" {
		if valuetobevalidated == "" {
			validationFlag = false
			return validationFlag
		} else {
			if len(valuetobevalidated) < 3 || len(valuetobevalidated) > 20 {
				validationFlag = false
				return validationFlag
			} else {
				if !(regexp.MustCompile((`^[A-Za-z]+$`)).MatchString(valuetobevalidated)) {
					validationFlag = false
					return validationFlag
				}
			}
		}
	} else if tobevalidated == "lastname" {
		if valuetobevalidated == "" {
			validationFlag = false
			return validationFlag
		} else {
			if len(valuetobevalidated) < 3 || len(valuetobevalidated) > 20 {
				validationFlag = false
				return validationFlag
			} else {
				if !(regexp.MustCompile(`^[A-Za-z]+$`).MatchString(valuetobevalidated)) {
					validationFlag = false
					return validationFlag
				}
			}
		}
	} else if tobevalidated == "email" {
		if valuetobevalidated == "" {
			fmt.Println("Empty email")
			validationFlag = false
			return validationFlag
		} else {
			if len(valuetobevalidated) > 40 {
				fmt.Println("invalid length for email")
				validationFlag = false
				return validationFlag
			} else {
				if !(regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(valuetobevalidated)) {
					fmt.Println("invalid email")
					validationFlag = false
					return validationFlag
				}
			}
		}
	} else if tobevalidated == "phonenumber" {
		if valuetobevalidated == "" {
			validationFlag = false
			return validationFlag
		} else {
			if len(valuetobevalidated) != 10 {
				validationFlag = false
				return validationFlag
			} else {
				if !(regexp.MustCompile(`(^\d+$)`).MatchString(valuetobevalidated)) {
					validationFlag = false
					return validationFlag
				}
			}
		}
	} else if tobevalidated == "id" {
		if valuetobevalidated == "" {
			validationFlag = false
			return validationFlag
		} else {
			if !(regexp.MustCompile(`(^\d+$)`).MatchString(valuetobevalidated)) {
				validationFlag = false
				return validationFlag
			}
		}
	}

	return validationFlag
}

func ValidateTaskOrMessageInfoFormInputs(tobevalidated string, valuetobevalidated string) bool {
	validationFlag := true
	if tobevalidated == "priority" {
		if valuetobevalidated == "" {
			validationFlag = false
			return validationFlag
		} else {
			if len(valuetobevalidated) < 3 || len(valuetobevalidated) > 6 {
				validationFlag = false
				return validationFlag
			} else {
				if !(regexp.MustCompile((`^[A-Za-z]+$`)).MatchString(valuetobevalidated)) {
					validationFlag = false
					return validationFlag
				}
			}
		}
	} else if tobevalidated == "category" {
		if tobevalidated == "" {
			validationFlag = false
			return validationFlag
		} else {
			if len(valuetobevalidated) > 20 {
				validationFlag = false
				return validationFlag
			} else {
				if !(regexp.MustCompile((`^[A-Za-z]+$`)).MatchString(valuetobevalidated)) {
					validationFlag = false
					return validationFlag
				}
			}
		}
	} else if tobevalidated == "title" {
		if valuetobevalidated == "" {
			validationFlag = false
			return validationFlag
		} else {
			if len(valuetobevalidated) < 3 || len(valuetobevalidated) > 20 {
				validationFlag = false
				return validationFlag
			}
		}
	} else if tobevalidated == "description" {
		if valuetobevalidated == "" {
			validationFlag = false
			return validationFlag
		} else {
			if len(valuetobevalidated) < 3 || len(valuetobevalidated) > 20 {
				validationFlag = false
				return validationFlag
			}
		}
	}
	return validationFlag
}

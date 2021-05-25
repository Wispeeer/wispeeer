package utils

import (
	"log"
	"regexp"
)

// IsValid ...
func IsValid(str string) bool {
	reg := regexp.MustCompile(`[\\\\/:*?\"<>|]`)
	if reg == nil {
		log.Println("Title Incorrect")
		return false
	}
	result := reg.FindAllString(str, -1)
	return len(result) <= 0
}

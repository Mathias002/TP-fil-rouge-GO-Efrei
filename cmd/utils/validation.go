package utils

import (
	"bufio"
	"regexp"
	"strings"
)

func IsValidEmail(email string) bool {
	// Pattern regex simple pour email
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)
}


func ReaderLine(reader *bufio.Reader) (string, error) {
	readerValue, _ := reader.ReadString('\n')
	readerValue = strings.TrimSpace(readerValue)
	return readerValue, nil
}

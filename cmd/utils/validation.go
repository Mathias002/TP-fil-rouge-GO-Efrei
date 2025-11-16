package utils

import (
	"bufio"
	"regexp"
	"strings"
)

// Permet de vérifier si le format de l'email est correct 
func IsValidEmail(email string) bool {
	// Pattern regex simple pour email
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)
}

// Permet de lire et nettoyer la saisie d'un utilisateur depuis un reader
func ReaderLine(reader *bufio.Reader) (string, error) {
	readerValue, _ := reader.ReadString('\n')
	readerValue = strings.TrimSpace(readerValue)
	return readerValue, nil
}

package main_menu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type optionSelected uint

func Main_menu() optionSelected {

	// Texte menu principal
	fmt.Println("--- Mini CRM ---")
	fmt.Println("1. Ajouter un contact")
	fmt.Println("2. Lister les contacts")
	fmt.Println("3. Modifier un contact")
	fmt.Println("4. Supprimer un contact")
	fmt.Println("5. Quitter")

	// reader pour retourner le choix de l'utilisateur.rice
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Votre choix : ")
	choiceStr, _ := reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr)

	choiseParsed, err := strconv.ParseUint(choiceStr, 10, 64)
	if err != nil {
		fmt.Println("Entrée invalide:", err)
		return optionSelected(0)
	}

	return optionSelected(choiseParsed)
}

package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Mathias002/TP-fil-rouge-GO-efrei/cmd/utils"
	storage "github.com/Mathias002/TP-fil-rouge-GO-efrei/internal/store"
	"github.com/spf13/cobra"
)

var (
	nameStr  string
	emailStr string
	err      error
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Permet d'ajouter un contact",
	Long:  `La commande 'add' permet l'ajout d'un contact via la saisie de deux champs ; 'name' et 'email'`,
	Run: func(cmd *cobra.Command, args []string) {

		// faire la section interactive d'abord le name puis l'email

		if nameStr == "" {
			fmt.Print("➡️  Renseigner le nom du contact: ")
			reader := bufio.NewReader(os.Stdin)
			nameStr, err = utils.ReaderLine(reader)
			if err != nil {
				log.Fatalf("‼️Erreur de saisi: %v \n", err)
			}
		}

		if emailStr == "" {
			fmt.Print("➡️  Renseigner l'email du contact: ")
			reader := bufio.NewReader(os.Stdin)
			emailStr, err = utils.ReaderLine(reader)
			if err != nil {
				log.Fatalf("‼️Erreur de saisi: %v \n", err)
			}
		}

		name := storage.NameContact(nameStr)
		email := storage.EmailContact(emailStr)

		if !utils.IsValidEmail(string(email)) {
			fmt.Println()
			fmt.Println("‼️Erreur: L'email renseigné n'est pas valide veuillez respecter le format suivant : email.example@gmail.com")
			fmt.Println()
			return
		}

		newContact := &storage.Contact{
			Name:  name,
			Email: email,
		}

		err := store.AddContact(newContact)
		if err != nil {
			fmt.Println()
			log.Fatalf("‼️Erreur lors de l'ajout du contact: %v", err)
			fmt.Println()
		}
		fmt.Println()
		fmt.Printf("✅ Nouveau contact '%s' ajouté.e avec l'ID %d \n", newContact.Name, newContact.ID)
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Définition des drapeaux pour la commande 'add'
	addCmd.Flags().StringVarP(&nameStr, "name", "n", "", "Nom du nouveau contact")
	addCmd.Flags().StringVarP(&(emailStr), "email", "o", "", "Email du nouveau contact")

	// Marquer tous les drapeaux comme obligatoires
	// addCmd.MarkFlagRequired("name")
	// addCmd.MarkFlagRequired("email")
}

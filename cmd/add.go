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
	name  string
	email string
	err   error
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Permet d'ajouter un contact",
	Long:  `La commande 'add' permet l'ajout d'un contact via la saisie de deux champs, 'name' et 'email'`,
	Run: func(cmd *cobra.Command, args []string) {

		// On vérifie si le nom à été fourni via les flags de la commande, si non on le demande directement via un reader
		if name == "" {
			fmt.Print("➡️  Renseigner le nom du contact: ")
			reader := bufio.NewReader(os.Stdin)
			name, err = utils.ReaderLine(reader)
			if err != nil {
				log.Fatalf("‼️  Erreur de saisi: %v \n", err)
			}
		}

		// On vérifie si l'emial à été fourni via les flags de la commande, si non on le demande directement via un reader
		if email == "" {
			fmt.Print("➡️  Renseigner l'email du contact: ")
			reader := bufio.NewReader(os.Stdin)
			email, err = utils.ReaderLine(reader)
			if err != nil {
				log.Fatalf("‼️  Erreur de saisi: %v \n", err)
			}
		}

		// On vérifie si l'email est valide 
		if !utils.IsValidEmail(email) {
			fmt.Println()
			fmt.Println("‼️  Erreur: L'email renseigné n'est pas valide veuillez respecter le format suivant : email.example@gmail.com")
			fmt.Println()
			return
		}

		// On init un nouveau contact avec les informations fourni par l'utilisateur
		newContact := &storage.Contact{
			Name:  name,
			Email: email,
		}

		// On appel la func d'ajout d'un contact en lui passant le nouveau
		err := store.AddContact(newContact)

		// Gestion des erreurs
		if err != nil {
			fmt.Println()
			log.Fatalf("‼️  Erreur lors de l'ajout du contact: %v", err)
			fmt.Println()
		}

		// Message de confirmation
		fmt.Println()
		fmt.Printf("✅ Nouveau contact '%s' ajouté.e avec l'ID %d \n", newContact.Name, newContact.ID)
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Définition des drapeaux pour la commande 'add'
	addCmd.Flags().StringVarP(&name, "name", "n", "", "Nom du nouveau contact")
	addCmd.Flags().StringVarP(&email, "email", "o", "", "Email du nouveau contact")

	// Marquer tous les drapeaux comme obligatoires
	// addCmd.MarkFlagRequired("name")
	// addCmd.MarkFlagRequired("email")
}

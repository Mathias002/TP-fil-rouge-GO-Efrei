package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var getAllCmd = &cobra.Command{
	Use:   "list",
	Short: "Permet de récupérer la liste de tous les contacts",
	Long:  `La commande 'list' permet la récupération de tous les contacts avec leurs informations tel que : 'id', 'name', 'email'`,
	Run: func(cmd *cobra.Command, args []string) {

		// On appel la func d'affichage de tous les contacts
		contacts, err := store.DisplayContacts()

		// Gestion des erreurs
		if err != nil {
			fmt.Println()
			log.Fatalf("‼️  Erreur lors de l'affichage des contacts: %v", err)
			fmt.Println()
		}

		fmt.Println()
		fmt.Println("--- Liste des contacts ---")
		fmt.Println()

		// Message sympa en cas d'absence de contact
		if len(contacts) == 0 {
			fmt.Printf("Humm, il semblerait que vous n'ayez pas d'amis...")
		}

		// boucle sur les contacts
		for _, contact := range contacts {
			fmt.Printf("ID: %d | Name: %s | Email: %s\n", contact.ID, contact.Name, contact.Email)
		}
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(getAllCmd)
}

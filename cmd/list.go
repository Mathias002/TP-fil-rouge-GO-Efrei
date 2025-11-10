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

		contacts, err := store.DisplayContacts()
		if err != nil {
			fmt.Println()
			log.Fatalf("‼️Erreur lors de l'affichage des contacts: %v", err)
			fmt.Println()
		}

		fmt.Println()
		fmt.Println("--- Liste des contacts ---")
		fmt.Println()

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

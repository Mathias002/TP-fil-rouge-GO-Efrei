package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var idContactGet int

var getOneByIdCmd = &cobra.Command{
	Use:   "get",
	Short: "Permet de récupérer un contact via son id",
	Long:  `La commande 'get' permet la récupération d'un contact via son id avec ses informations tel que : 'id', 'name', 'email'`,
	Run: func(cmd *cobra.Command, args []string) {

		// On appel la func d'affichage d'un contact en lui passant l'id du contact en paramètre
		contact, err := store.DisplayContact(idContactGet)

		// Gestion des erreurs 
		if err != nil {
			fmt.Println()
			log.Fatalf("‼️  Erreur lors de l'affichage des informations du contact: %v", err)
			fmt.Println()
		}

		// Affichage des informations du contact
		fmt.Println()
		fmt.Printf("--- Information du contact avec l'ID: %d ---", idContactGet)
		fmt.Println()

		fmt.Println()
		fmt.Printf("ID: %d | Name: %s | Email: %s\n", contact.ID, contact.Name, contact.Email)
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(getOneByIdCmd)

	// Définition du drapeau pour la commande 'get_by_id'
	getOneByIdCmd.Flags().IntVarP(&idContactGet, "id", "i", 0, "ID du contact")

	// Marquer le drapeau id comme obligatoires
	addCmd.MarkFlagRequired("id")
}

package cmd

import (
	"fmt"
	"log"

	storage "github.com/Mathias002/TP-fil-rouge-GO-efrei/internal/store"
	"github.com/spf13/cobra"
)

var idContactGet int

var getOneByIdCmd = &cobra.Command{
	Use:   "get",
	Short: "Permet de récupérer un contact via son id",
	Long:  `La commande 'get' permet la récupération d'un contact via son id avec ses informations tel que : 'id', 'name', 'email'`,
	Run: func(cmd *cobra.Command, args []string) {

		contact, err := store.DisplayContact(storage.IDContact(idContactGet))
		if err != nil {
			fmt.Println()
			log.Fatalf("‼️Erreur lors de l'affichage des des informations du contact: %v", err)
			fmt.Println()
		}

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

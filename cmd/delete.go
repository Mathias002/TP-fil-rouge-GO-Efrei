package cmd

import (
	"fmt"
	"log"

	storage "github.com/Mathias002/TP-fil-rouge-GO-efrei/internal/store"
	"github.com/spf13/cobra"
)

var idContactDelete int

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Permet de supprimer un contact",
	Long:  `La commande 'delete' permet la suppression d'un contact via son id`,
	Run: func(cmd *cobra.Command, args []string) {

		err := store.DeleteContact(storage.IDContact(idContactDelete))
		if err != nil {
			fmt.Println()
			log.Fatalf("‼️Erreur lors de la suppression du contact : %v", err)
			fmt.Println()
		}

		fmt.Println()
		fmt.Printf("✅ Le contact avec l'ID '%d' a bien été supprimé \n", idContactDelete)
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Définition du drapeau pour la commande 'get_by_id'
	deleteCmd.Flags().IntVarP(&idContactDelete, "id", "i", 0, "ID du contact")

	// Marquer le drapeau id comme obligatoires
	deleteCmd.MarkFlagRequired("id")
}

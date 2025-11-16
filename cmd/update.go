package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Mathias002/TP-fil-rouge-GO-efrei/cmd/utils"
	"github.com/spf13/cobra"
)

var (
	idContactUpdate int
	newName         string
	newEmail        string
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Permet de modifier un contact",
	Long:  `La commande 'update' permet la modification d'un contact via la saisie de deux champs ; 'name' et 'email'`,
	Run: func(cmd *cobra.Command, args []string) {

		// On vérifie si le nom à été fourni via les flags de la commande, si non on le demande directement via un reader
		if newName == "" {
			fmt.Print("➡️  Renseigner le nouveau nom du contact: ")
			reader := bufio.NewReader(os.Stdin)
			newName, err = utils.ReaderLine(reader)
			if err != nil {
				log.Fatalf("‼️  Erreur de saisi: %v \n", err)
			}
		}

		// On vérifie si l'emial à été fourni via les flags de la commande, si non on le demande directement via un reader
		if newEmail == "" {
			fmt.Print("➡️  Renseigner le nouveau email du contact: ")
			reader := bufio.NewReader(os.Stdin)
			newEmail, err = utils.ReaderLine(reader)
			if err != nil {
				log.Fatalf("‼️  Erreur de saisi: %v \n", err)
			}
		}

		// On vérifie si l'email est valide
		if newEmail != "" && !utils.IsValidEmail(string(newEmail)) {
			fmt.Println()
			fmt.Println("‼️  Erreur: L'email renseigné n'est pas valide veuillez respecter le format suivant : email.example@gmail.com")
			fmt.Println()
			return
		}

		// On appel la func de mise à jour d'un contact en lui passant en paramètre l'id du contact à modifier et les nouvelles informations du contact
		err := store.UpdateContact(idContactUpdate, newName, newEmail)
		
		// Gestion des erreurs
		if err != nil {
			fmt.Println()
			log.Fatalf("‼️  Erreur lors de la modification du contact : %v", err)
			fmt.Println()
		}

		// Message de confirmation
		fmt.Println()
		fmt.Printf("✅ Mise à jour des informations du contact avec l'ID '%d' effectué avec succés \n", idContactUpdate)
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Définition des drapeaux pour la commande 'add'
	updateCmd.Flags().IntVarP(&idContactUpdate, "id", "i", 0, "Id du contact à mettre à jour")
	updateCmd.Flags().StringVarP(&newName, "name", "n", "", "Nouveau nom du contact")
	updateCmd.Flags().StringVarP(&newEmail, "email", "e", "", "Nouveau email du contact")

	// Marquer tous les drapeaux comme obligatoires
	updateCmd.MarkFlagRequired("id")
	// updateCmd.MarkFlagRequired("name")
	// updateCmd.MarkFlagRequired("email")
}

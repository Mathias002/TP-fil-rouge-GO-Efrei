package main

import (
	"flag"
	"fmt"

	"github.com/Mathias002/TP-fil-rouge-GO-efrei/contact"
	"github.com/Mathias002/TP-fil-rouge-GO-efrei/main_menu"
)

func main() {
	// On init les contacts par défaut
	contact.Init()

	// On Init les flags
	addName := flag.String("name", "", "Nom du contact")
	addEmail := flag.String("email", "", "Email du contact")

	// On Parse les flags
	flag.Parse()

	// On vérifie que des flags sont fournis <> si oui on ajoute le contact sinon on lance le menu interactif
	if *addName != "" && *addEmail != "" {
		contact.AddContactFlag(contact.NameContact(*addName), contact.EmailContact(*addEmail))
		return // Quitter après l'ajout
	}

	// boucle infini pour l'affichage du menu en continu
	Loop_Main_Menu:
		for { // boucle infinie

			// Récupération du choix saisi par l'utilisateur.rice
			choice := main_menu.Main_menu()
			// if err != nil {
			// 	fmt.Println("Erreur : " + err)
			// 	break
			// }

			// switch sur les options du menu
			// en fonction du choix on appele les fonctions correspondantes
			switch choice {
			// Ajout d'un contact 
			case 1:
				contact.AddContact()
			// Affichage des contacts 
			case 2:
				contact.Contacts.DisplayContacts()
			// Mise à jour d'un contact
			case 3:
				contact.UpdateContact()
			// Suppression d'un contact
			case 4:
				contact.DeleteContact()
			// Fermeture du programme 
			case 5:
				fmt.Println("Merci de votre visite à bientôt ! 🫡")
				break Loop_Main_Menu
			}
		}
}

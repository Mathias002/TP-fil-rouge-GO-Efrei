package app

import (
	"bufio"
	"strconv"
	"fmt"
	"os"
	"regexp"
	"strings"
	"github.com/Mathias002/TP-fil-rouge-GO-efrei/cmd/api/main_menu"
	storage "github.com/Mathias002/TP-fil-rouge-GO-efrei/internal/store"
)

func Run(storage storage.Storer) {
	reader := bufio.NewReader(os.Stdin)

	// // On Init les flags
	// addName := flag.String("name", "", "Nom du contact")
	// addEmail := flag.String("email", "", "Email du contact")

	// // On Parse les flags
	// flag.Parse()

	// // On vérifie que des flags sont fournis <> si oui on ajoute le contact sinon on lance le menu interactif
	// if *addName != "" && *addEmail != "" {
	// 	contact.AddContactFlag(contact.NameContact(*addName), contact.EmailContact(*addEmail))
	// 	return // Quitter après l'ajout
	// }

	// boucle infini pour l'affichage du menu en continu
Loop_Main_Menu:
	for { // boucle infinie

		// Récupération du choix saisi par l'utilisateur.rice
		choice := main_menu.Main_menu()

		// switch sur les options du menu
		// en fonction du choix on appele les fonctions correspondantes
		switch choice {
		// Ajout d'un contact
		case 1:
			handleAddContact(reader, storage)

		// Affichage des contacts
		case 2:
			handleDisplaContacts(storage)

		case 3:
			handleDisplayContact(reader, storage)

		// Mise à jour d'un contact
		case 4:
			handleUpdateContact(reader, storage)
		// Suppression d'un contact
		case 5:
			handleDeleteContact(reader, storage)
		// Fermeture du programme
		case 6:
			fmt.Println("Merci de votre visite à bientôt ! 🫡")
			break Loop_Main_Menu
		default:
			fmt.Println("Option inconu :/, veuillez réessayer")
			fmt.Println()
		}
	}
}

func handleAddContact(reader *bufio.Reader, store storage.Storer) {
	var name string
	var email string
	var err error
	fmt.Println()
	fmt.Println("--- Ajouter un contact ---")
	for {
		fmt.Println()
		fmt.Println("Entrez le nom du contact :")

		name, err = readerLine(reader)
		if err != nil {
			// Erreur
			fmt.Printf("Erreur de saisi: %v. veuillez réessayer. \n", err)
			continue
		}
		if name != "" {
			// La saisi est valide on continu
			break
		}
		fmt.Println("Le nom ne peut pas être vide. Veuillez renseigner une valeur")
	}

	for {
		fmt.Println()
		fmt.Println("Entrez l'email du contact :")

		email, err = readerLine(reader)
		if err != nil {
			// Erreur
			fmt.Printf("Erreur de saisi: %v. veuillez réessayer. \n", err)
			continue
		}
		if isValidEmail(email) {
			break
		}
		fmt.Println("L'email renseigné n'est pas valide veuillez respecter le format suivant : email.example@gmail.com")
		fmt.Println()
	}

	// La suite de la logique d'ajout
	contact := &storage.Contact{Name: storage.NameContact(name), Email: storage.EmailContact(email)}
	err = store.AddContact(contact)
	if err != nil {
		fmt.Printf("Erreur: %v\n", err)
		return
	}
	fmt.Println()
	fmt.Printf("Nouveau contact '%s' ajouté.e avec l'ID %d \n", contact.Name, contact.ID)
	fmt.Println()
}

func handleDisplaContacts(store storage.Storer) {
	var contacts []*storage.Contact
	var err error

	fmt.Println()
	fmt.Println("--- Affichage des contacts ---")
	fmt.Println()

	contacts, err = store.DisplayContacts()
	if err != nil {
		fmt.Println()
		fmt.Printf("Erreur: %v\n", err)
		fmt.Println()
		return
	}
	// boucle sur les contacts
	for _, contact := range contacts {
		fmt.Printf("ID: %d | Name: %s | Email: %s\n", contact.ID, contact.Name, contact.Email)
	}
	fmt.Println()
}

func handleDisplayContact(reader *bufio.Reader, store storage.Storer) {
	var idContact string
	var idContactParsed int
	var err error

	fmt.Println()
	fmt.Print("--- Affichage d'un utilisateur ---")
	fmt.Println()

	fmt.Println("Entrez l'ID du contact qui vous souhaitez consulter")

	idContact, err = readerLine(reader)
	if err != nil {
		// Erreur
		fmt.Println()
		fmt.Printf("Erreur : %v \n", err)
		fmt.Println()
	}

	idContactParsed, err = strconv.Atoi(idContact)
	if err != nil {
		// Erreur
		fmt.Println()
		fmt.Println("Entrée invalide:", err)
		fmt.Println()
		return
	}

	contact, err := store.DisplayContact(storage.IDContact(idContactParsed))
	if err != nil {
		// Erreur
		fmt.Println()
		fmt.Printf("Erreur : %v \n", err)
		fmt.Println()
		return
	}

	fmt.Println()
	fmt.Printf("Information de l'utilisateur avec l'ID : %d", contact.ID)
	fmt.Println()
	fmt.Printf("Name: %s | Email: %s\n", contact.Name, contact.Email)
	fmt.Println()
}

func handleUpdateContact(reader *bufio.Reader, store storage.Storer) {
	var idContact string
	var idContactParsed int
	var err error
	var newName string
	var newEmail string
	var changes = []string{}

	fmt.Println()
	fmt.Print("--- Début de la phase de modification d'un contact ---")
	fmt.Println()

	for {
		fmt.Println()
		fmt.Println("Entrez l'ID du contact à modifier:")

		idContact, err = readerLine(reader)
		if err != nil {
			// Erreur
			fmt.Println()
			fmt.Printf("Erreur de saisi: %v. veuillez réessayer. \n", err)
			continue
		}
		idContactParsed, err = strconv.Atoi(idContact)
		if err != nil {
			// Erreur
			fmt.Println()
			fmt.Println("Entrée invalide:", err)
			fmt.Println()
			continue
		}
		fmt.Println()
		fmt.Printf("Modification de l'utilisateur avec l'ID: %d", idContactParsed)
		break
	}

	for {
		fmt.Println()
		fmt.Printf("Entrez le nouveau nom de l'utilisateur (vide si pas de changement):")
		fmt.Println()

		newName, err = readerLine(reader)
		if err != nil {
			// Erreur
			fmt.Println()
			fmt.Printf("Erreur de saisi: %v. veuillez réessayer. \n", err)
			fmt.Println()
			continue
		}
		if newName != "" {
			changes = append(changes, fmt.Sprintf("Nom: → %s", newName))
		}
		break
	}

	for {

		fmt.Println()
		fmt.Printf("Entrez le nouveau email de l'utilisateur (vide si pas de changement):")
		fmt.Println()

		newEmail, err = readerLine(reader)
		if err != nil {
			// Erreur
			fmt.Println()
			fmt.Printf("Erreur de saisi: %v. veuillez réessayer. \n", err)
			fmt.Println()
		}
		if newEmail != "" {
			if !isValidEmail(newEmail) {
				fmt.Println()
				fmt.Println("L'email renseigné n'est pas valide veuillez respecter le format suivant : email.example@gmail.com")
				fmt.Println()
				continue
			}
			changes = append(changes, fmt.Sprintf("Email: → %s", newEmail))
		}
		break
	}

	err = store.UpdateContact(storage.IDContact(idContactParsed), storage.NameContact(newName), storage.EmailContact(newEmail))
	if err != nil {
		fmt.Printf("Erreur: %v\n", err)
		fmt.Println()
		return
	}

	fmt.Println()
	if len(changes) > 0 {
		fmt.Printf("✅ Contact avec l'ID %d mis à jour\n", idContactParsed)
		fmt.Println("\nModifications effectuées :")
		for _, change := range changes {
			fmt.Printf("  • %s\n", change)
		}
	} else {
		fmt.Println("ℹ️  Aucune modification effectuée")
	}
	fmt.Println()
}

func handleDeleteContact(reader *bufio.Reader, store storage.Storer) {
	var idContact string
	var idContactParsed int
	var err error

	fmt.Println()
	fmt.Printf("--- Supprimer un contact ---")
	fmt.Println()

	for {
		fmt.Println()
		fmt.Println("Entrez l'ID du contact à supprimer:")

		idContact, err = readerLine(reader)
		if err != nil {
			// Erreur
			fmt.Println()
			fmt.Printf("Erreur de saisi: %v. veuillez réessayer. \n", err)
			continue
		}
		idContactParsed, err = strconv.Atoi(idContact)
		if err != nil {
			// Erreur
			fmt.Println()
			fmt.Println("Entrée invalide:", err)
			fmt.Println()
			continue
		}
		break
	}

	err = store.DeleteContact(storage.IDContact(idContactParsed))
	if err != nil {
		fmt.Println()
		fmt.Printf("Erreur: %v\n", err)
		fmt.Println()
		return
	}

	fmt.Println()
	fmt.Printf("Utilisateur avec l'ID : %d supprimé.e", idContactParsed)
	fmt.Println()
}

func readerLine(reader *bufio.Reader) (string, error) {
	readerValue, _ := reader.ReadString('\n')
	readerValue = strings.TrimSpace(readerValue)
	return readerValue, nil
}

func isValidEmail(email string) bool {
	// Pattern regex simple pour email
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)
}

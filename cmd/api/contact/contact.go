package contact

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type IDContact int
type NameContact string
type EmailContact string
type ContactsType map[IDContact]Contact

// Définition d'une structure "Contact"
type Contact struct {
	ID    IDContact
	Name  NameContact
	Email EmailContact
}

// Initialisation d'un map Contacts global
var Contacts ContactsType

// function d'initialisation des utilisateurs par défaut
func Init() {
	// Initialisation du générateur aléatoire
	rand.Seed(time.Now().UnixNano())

	// Ajout de mes utilisateurs par défaut dans mon map Contacts
	Contacts = make(ContactsType)

	id1 := IDContact(randomInteger())
	Contacts[id1] = Contact{
		ID:    id1,
		Name:  "Poipoi",
		Email: "poipoi@gmail.com",
	}

	id2 := IDContact(randomInteger())
	Contacts[id2] = Contact{
		ID:    id2,
		Name:  "Lala",
		Email: "lala@gmail.com",
	}
}

// Function qui retourne un nombre aléatoire entre 0 et 9999
func randomInteger() int {
	return rand.Intn(10000)
}

// Function qui affiche tout les utilisateurs
func (c ContactsType) DisplayContacts() {
	fmt.Printf("\n")
	fmt.Println("--- Liste des contacts ---")
	for id, contact := range c {
		fmt.Printf("ID: %d | Name: %s | Email: %s\n", id, contact.Name, contact.Email)
	}
	fmt.Printf("\n")
}

// Function qui affiche un utilisateur via son ID
func (c Contact) DisplayContact(id IDContact) {
	fmt.Printf("\n")
	fmt.Printf("--- Contact avec l'ID: %d ---", id)
	// contactToDisplay, exists := Contacts[IDContact(id)]
	// c = Contacts[IDContact(id)]
	// if !exists {
	// 	fmt.Printf("Utilisateur.rice avec l'ID %d n'existe pas veuillez réessayer", id)
	// }
	fmt.Printf("\n")
	fmt.Printf("ID: %d | Nom: %s | Email: %s\n",
		c.ID,
		c.Name,
		c.Email,
	)
	fmt.Printf("\n")
}

// Crée et valide un nouveau contact
func NewContact(name string, email string) (*Contact, error) {
	// Validation du nom
	if name == "" {
		return nil, errors.New("le nom ne peut pas être vide")
	}

	// Validation de l'email
	if email == "" {
		return nil, errors.New("l'email ne peut pas être vide")
	}

	if !isValidEmail(email) {
		return nil, errors.New("l'email n'est pas valide")
	}

	// Création du contact avec ID unique
	newContact := &Contact{
		ID:    IDContact(randomInteger()),
		Name:  NameContact(name),
		Email: EmailContact(email),
	}

	// Ajout dans le map
	Contacts[newContact.ID] = *newContact

	return newContact, nil
}

// isValidEmail valide le format d'un email
func isValidEmail(email string) bool {
	// Pattern regex simple pour email
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)
}

// Function qui permet d'ajouter un utilisateur
func AddContact() {

	for {

		// utiliser buffer pour récupérer les infos
		fmt.Println("\n")
		fmt.Println("--- Ajout d'un contact ---")
		fmt.Println("Entrez le nom du contact :")
		// reader pour retourner le nom du contact
		readerContactName := bufio.NewReader(os.Stdin)
		nameContact, _ := readerContactName.ReadString('\n')
		nameContact = strings.TrimSpace(nameContact)

		fmt.Println("Entrez l'adresse email du contact :")

		// reader pour retourner l'email du contact
		readerEmailEmail := bufio.NewReader(os.Stdin)
		emailContact, _ := readerEmailEmail.ReadString('\n')
		emailContact = strings.TrimSpace(emailContact)

		// On créer le nouveau contact et on l'ajoute dans le map global Contacts
		// Création du contact avec validation
		newContact, err := NewContact(nameContact, emailContact)
		if err != nil {
			fmt.Println(err.Error())
			AddContact() // Réessayer
			return
		}

		// formatage message de retour
		fmt.Println("\n")
		fmt.Println("Contact ajouté : ")
		fmt.Printf("ID: %d | Nom: %s | Email: %s\n",
			newContact.ID,
			newContact.Name,
			newContact.Email,
		)
		fmt.Println("\n")
		break
	}
}

// Function pour supprimer un utilisateur
func DeleteContact() {

	Contacts.DisplayContacts()

	fmt.Println("\n")
	fmt.Println("--- Supprimer un contact ---")
	fmt.Println("Entrez l'ID du contact à supprimer:")
	// reader pour retourner l'ID du contact à supprimer
	readerContactID := bufio.NewReader(os.Stdin)
	idContact, _ := readerContactID.ReadString('\n')
	idContact = strings.TrimSpace(idContact)

	idContactParsed, err := strconv.ParseInt(idContact, 10, 64)
	if err != nil {
		fmt.Println("Entrée invalide:", err)
		DeleteContact()
	}

	delete(Contacts, IDContact(idContactParsed))

	fmt.Println("\n")
	fmt.Printf(" ✅✅✅ Utilisateur.rice avec l'ID %d supprimé.e avec succés ✅✅✅ \n", idContactParsed)
	Contacts.DisplayContacts()

}

// Function pour modifier un utilisateur
func UpdateContact() {

	Contacts.DisplayContacts()

	fmt.Println("\n")
	fmt.Println("--- Mettre à jour un contact ---")
	fmt.Println("Entrez l'ID du contact à mettre à jour:")
	// reader pour retourner l'ID du contact à supprimer
	readerContactID := bufio.NewReader(os.Stdin)
	idContact, _ := readerContactID.ReadString('\n')
	idContact = strings.TrimSpace(idContact)

	idContactParsed, err := strconv.Atoi(idContact)
	if err != nil {
		fmt.Println("Entrée invalide:", err)
		UpdateContact()
	}

	// comma ok idiom
	contactToUpdate, exists := Contacts[IDContact(idContactParsed)]
	if !exists {
		fmt.Printf("Utilisateur.rice avec l'ID %d n'existe pas veuillez réessayer", idContactParsed)
	}

	fmt.Printf("--- Début de la modification des informations de l'utilisateur.rice avec l'ID %d ---", idContactParsed)
	fmt.Println("\n")
	fmt.Println("Entrez le nouveau nom de l'utilisateur.rice (vide si pas de changement)")
	// reader pour retourner le nom du contact à modifier
	readerNewContactName := bufio.NewReader(os.Stdin)
	newContactName, _ := readerNewContactName.ReadString('\n')
	newContactName = strings.TrimSpace(newContactName)

	fmt.Println("Entrez le nouveau email de l'utilisateur.rice (vide si pas de changement)")
	// reader pour modifier l'email du contact à modifier
	readerNewContactEmail := bufio.NewReader(os.Stdin)
	newContactEmail, _ := readerNewContactEmail.ReadString('\n')
	newContactEmail = strings.TrimSpace(newContactEmail)

	// Si les valeur renseigné sont différente des valeurs existante on update
	if contactToUpdate.Name != NameContact(newContactName) && NameContact(newContactName) != "" {
		contactToUpdate.Name = NameContact(newContactName)
	}
	if contactToUpdate.Email != EmailContact(newContactEmail) && EmailContact(newContactEmail) != "" {
		contactToUpdate.Email = EmailContact(newContactEmail)
	}

	Contacts[IDContact(idContactParsed)] = contactToUpdate

	fmt.Println("\n")
	fmt.Printf(" ✅✅✅ Utilisateur.rice avec l'ID %d modifié.e avec succés ✅✅✅ \n", idContactParsed)
	contactToUpdate.DisplayContact(IDContact(idContactParsed))

}

// Function pour ajouter un utilisateur via des flags
func AddContactFlag(nameContact NameContact, emailContact EmailContact) {
	newIdContact := IDContact(randomInteger())
	Contacts[newIdContact] = Contact{
		ID:    newIdContact,
		Name:  NameContact(nameContact),
		Email: EmailContact(emailContact),
	}

	// formatage message de retour
	fmt.Println("\n")
	fmt.Println("Contact ajouté : ")
	fmt.Printf("ID: %d | Nom: %s | Email: %s\n",
		newIdContact,
		nameContact,
		emailContact,
	)
	fmt.Println("\n")
}

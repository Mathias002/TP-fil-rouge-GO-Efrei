package storage

import (
	// "errors"
	// "fmt"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

type JSONStore struct {
	contacts map[int]*Contact
	filename string
}

// Permet de créer le store JSON utilisé pour les différentes opération CRUD et à la persistence des données
func NewJSONStore(filename string) *JSONStore {
	store := &JSONStore{
		contacts: make(map[int]*Contact),
		filename: filename,
	}

	// Charger les contacts existants si le fichier existe
	if data, err := os.ReadFile(filename); err == nil {
		json.Unmarshal(data, &store.contacts)
	}

	return store
}

// Permet de générer un nombre aléatoire entre 0 et 9999
func (js *JSONStore) randomInteger() int {
	return rand.Intn(10000)
}

// Permet d'ajouter un contact
func (js *JSONStore) AddContact(contact *Contact) error {

	// On complète les informations fourni par l'utilisateur (Name | Email)
	// en générant un id aléatoire et en l'attribuant au nouveau contact
	contact.ID = js.randomInteger()
	js.contacts[contact.ID] = contact


	// Sauvegarde de l'ensemble des contacts au format JSON indenté
	data, err := json.MarshalIndent(js.contacts, "", "  ")
	if err != nil {
		return fmt.Errorf("impossible de marshaller les contacts %s: %w", js.filename, err)
	}

	// Réécriture du fichier de stockage avec les données JSON
	if err := os.WriteFile(js.filename, data, 0644); err != nil {
		return fmt.Errorf("impossible d'écrire dans le fichier %s: %w", js.filename, err)
	}
	return nil
}

// Permet d'afficher tout les contacts
func (js *JSONStore) DisplayContacts() ([]*Contact, error) {

	// création d'un slice de la taille du map des contacts
	contacts := make([]*Contact, 0, len(js.contacts))

	// bouclage sur le map en ignorant la clé IDContact pour ajouter la valeur du contact dans le slice contacts
	for _, contact := range js.contacts {
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

// Permet d'afficher un contact via son ID
func (js *JSONStore) DisplayContact(idContact int) (*Contact, error) {
	contact, exists := js.contacts[idContact]
	if !exists {
		return nil, fmt.Errorf("le contact avec l'ID %d n'existe pas", idContact)
	}

	return contact, nil
}

// Permet de mettre à jour un contact via son ID
func (js *JSONStore) UpdateContact(idContact int, newName string, newEmail string) error {

	// On vérifie que le contact existe
	contactToUpdate, exists := js.contacts[idContact]
	if !exists {
		return fmt.Errorf("le contact avec l'ID %d n'existe pas", idContact)
	}

	// On met à jour le contact dans le map mémoire
	if newName != "" && contactToUpdate.Name != newName {
		contactToUpdate.Name = newName
	}

	if newEmail != "" && contactToUpdate.Email != newEmail {
		contactToUpdate.Email = newEmail
	}

	// Sauvgarde de toute la map dans le fichier
	data, err := json.MarshalIndent(js.contacts, "", "  ")
	if err != nil {
		return fmt.Errorf("impossible de marshaller les contacts %s: %w", js.filename, err)
	}

	// Réecriture du fichier
	if err := os.WriteFile(js.filename, data, 0644); err != nil {
		return fmt.Errorf("impossible d'écrire dans le fichier %s: %w", js.filename, err)
	}
	return nil
}

// Permet de supprimer un contact via son ID
func (js *JSONStore) DeleteContact(idContact int) error {
	// On vérifie que le contact existe
	contactToDelete, exists := js.contacts[idContact]
	if !exists {
		return fmt.Errorf("le contact avec l'ID %d n'existe pas", idContact)
	}

	// suppression du contact dans la map mémoire
	delete(js.contacts, contactToDelete.ID)

	// Sauvgarde de toute la map dans le fichier
	data, err := json.MarshalIndent(js.contacts, "", "  ")
	if err != nil {
		return fmt.Errorf("impossible de marshaller les contacts %s: %w", js.filename, err)
	}

	// Réecriture du fichier
	if err := os.WriteFile(js.filename, data, 0644); err != nil {
		return fmt.Errorf("impossible d'écrire dans le fichier %s: %w", js.filename, err)
	}

	return nil
}

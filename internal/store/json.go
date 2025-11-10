package storage

import (
	// "errors"
	// "fmt"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type JSONStore struct {
	contacts map[IDContact]*Contact
	filename string
}

func NewJSONStore(filename string) *JSONStore {
	rand.Seed(time.Now().UnixNano())
	store := &JSONStore{
		contacts: make(map[IDContact]*Contact),
		filename: filename,
	}

	// Charger les contacts existants si le fichier existe
	if data, err := os.ReadFile(filename); err == nil {
		json.Unmarshal(data, &store.contacts)
	}

	return store
}

func (js *JSONStore) randomInteger() IDContact {
	return IDContact(rand.Intn(10000))
}

func (js *JSONStore) AddContact(contact *Contact) error {

	contact.ID = js.randomInteger()

	js.contacts[contact.ID] = contact

	data, err := json.MarshalIndent(js.contacts, "", "  ")
	if err != nil {
		return fmt.Errorf("impossible de marshaller les contacts %s: %w", js.filename, err)
	}

	if err := os.WriteFile(js.filename, data, 0644); err != nil {
		return fmt.Errorf("impossible d'écrire dans le fichier %s: %w", js.filename, err)
	}
	return nil
}

func (js *JSONStore) DisplayContacts() ([]*Contact, error) {

	// création d'un slice de la taille du map des contacts
	contacts := make([]*Contact, 0, len(js.contacts))

	// bouclage sur le map en ignorant la clé IDContact pour ajouter la valeur du contact dans le slice contacts
	for _, contact := range js.contacts {
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func (js *JSONStore) DisplayContact(idContact IDContact) (*Contact, error) {
	contact, exists := js.contacts[IDContact(idContact)]
	if !exists {
		return nil, fmt.Errorf("le contact avec l'ID %d n'existe pas", idContact)
	}

	return contact, nil
}

func (js *JSONStore) UpdateContact(idContact IDContact, newName NameContact, newEmail EmailContact) error {

	// On vérifie que le contact existe
	contactToUpdate, exists := js.contacts[IDContact(idContact)]
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

func (js *JSONStore) DeleteContact(idContact IDContact) error {
	// On vérifie que le contact existe
	contactToDelete, exists := js.contacts[IDContact(idContact)]
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

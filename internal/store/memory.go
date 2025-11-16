package storage

import (
	"errors"
	"fmt"
	"math/rand"
)

type MemoryStore struct {
	contacts map[int]*Contact
}

// Permet de créer le store en mémoire utilisé pour les différentes opération CRUD et à la persistence des données
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		contacts: make(map[int]*Contact),
	}
}

// Permet de générer un nombre aléatoire entre 0 et 9999
func (ms *MemoryStore) randomInteger() int {
	return rand.Intn(10000)
}

// Permet d'ajouter un contact
func (ms *MemoryStore) AddContact(contact *Contact) error {

	// On complète les informations fourni par l'utilisateur (Name | Email)
	// en générant un id aléatoire et en l'attribuant au nouveau contact
	contact.ID = ms.randomInteger()
	ms.contacts[contact.ID] = contact
	return nil
}

// Permet d'afficher tout les contacts
func (ms *MemoryStore) DisplayContacts() ([]*Contact, error) {

	// On vérifie si il existe au moins un contact
	if len(ms.contacts) == 0 {
		return nil, errors.New("aucun contact... vous n'avez pas d'amis")
	}

	// On initialise la var contacts en tant que slice de *Contact
	// de la longueur du nombre de contacts existant
	contacts := make([]*Contact, 0, len(ms.contacts))

	// On boucle sur les contacts en mémoire et on les ajoutes au slice contacts
	for _, contact := range ms.contacts {
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

// Permet d'afficher un contact via son ID
func (ms *MemoryStore) DisplayContact(idContact int) (*Contact, error) {
	contact, exists := ms.contacts[idContact]
	if !exists {
		return nil, fmt.Errorf("le contact avec l'ID %d n'existe pas", idContact)
	}

	return contact, nil
}

// Permet de mettre à jour un contact via son ID
func (ms *MemoryStore) UpdateContact(idContact int, newName string, newEmail string) error {

	// On vérifie si le contact existe
	contactToUpdate, exists := ms.contacts[idContact]
	if !exists {
		return fmt.Errorf("le contact avec l'ID %d n'existe pas", idContact)
	}

	// On met à jour les informations nécessaires
	if newName != "" && contactToUpdate.Name != newName {
		contactToUpdate.Name = newName
	}

	if newEmail != "" && contactToUpdate.Email != newEmail {
		contactToUpdate.Email = newEmail
	}

	return nil
}

// Permet de supprimer un contact via son ID
func (ms *MemoryStore) DeleteContact(idContact int) error {
	
	// On vérifie si le contact existe 
	contactToDelete, exists := ms.contacts[idContact]
	if !exists {
		return fmt.Errorf("le contact avec l'ID %d n'existe pas", idContact)
	}

	// On supprime le contact en mémoire
	delete(ms.contacts, contactToDelete.ID)

	return nil
}

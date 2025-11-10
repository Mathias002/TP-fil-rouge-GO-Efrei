package storage

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type MemoryStore struct {
	contacts map[IDContact]*Contact
}

func NewMemoryStore() *MemoryStore {
	rand.Seed(time.Now().UnixNano())
	return &MemoryStore{
		contacts: make(map[IDContact]*Contact),
	}
}

func (ms *MemoryStore) randomInteger() IDContact {
	return IDContact(rand.Intn(10000))
}

func (ms *MemoryStore) AddContact(contact *Contact) error {
	contact.ID = ms.randomInteger()
	ms.contacts[contact.ID] = contact
	return nil
}

func (ms *MemoryStore) DisplayContacts() ([]*Contact, error) {
	if len(ms.contacts) == 0 {
		return nil, errors.New("aucun contact... vous n'avez pas d'amis")
	}

	contacts := make([]*Contact, 0, len(ms.contacts))
	for _, contact := range ms.contacts {
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func (ms *MemoryStore) DisplayContact(idContact IDContact) (*Contact, error) {
	contact, exists := ms.contacts[IDContact(idContact)]
	if !exists {
		return nil, fmt.Errorf("le contact avec l'ID %d n'existe pas", idContact)
	}

	return contact, nil
}

func (ms *MemoryStore) UpdateContact(idContact IDContact, newName NameContact, newEmail EmailContact) error {
	contactToUpdate, exists := ms.contacts[IDContact(idContact)]
	if !exists {
		return fmt.Errorf("le contact avec l'ID %d n'existe pas", idContact)
	}

	if newName != "" && contactToUpdate.Name != newName {
		contactToUpdate.Name = newName
	}

	if newEmail != "" && contactToUpdate.Email != newEmail {
		contactToUpdate.Email = newEmail
	}
	return nil
}

func (ms *MemoryStore) DeleteContact(idContact IDContact) error {
	contactToDelete, exists := ms.contacts[IDContact(idContact)]
	if !exists {
		return fmt.Errorf("le contact avec l'ID %d n'existe pas", idContact)
	}

	delete(ms.contacts, contactToDelete.ID)

	return nil
}

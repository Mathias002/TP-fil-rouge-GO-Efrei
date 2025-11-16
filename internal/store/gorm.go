package storage

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/Mathias002/TP-fil-rouge-GO-efrei/internal/database"
	"gorm.io/gorm"
)

type GormStore struct {
	db *gorm.DB
}

// Permet de créer le store gorm utilisé pour les différentes opération CRUD et à la persistence des données
func NewGormStore() (*GormStore, error) {
	db := database.DB
	if db == nil {
		return nil, errors.New("la base de données n'est pas initialisée")
	}

	if err := db.AutoMigrate(&Contact{}); err != nil {
		return nil, fmt.Errorf("erreur de migration Contact: %w", err)
	}

	return &GormStore{db: db}, nil
}

// Permet de générer un nombre aléatoire entre 0 et 9999
func (gs *GormStore) randomInteger() int {
	return rand.Intn(10000)
}

// Permet d'ajouter un contact
func (gs *GormStore) AddContact(contact *Contact) error {

	// On complète les informations fourni par l'utilisateur (Name | Email)
	// en générant un id aléatoire et en l'attribuant au nouveau contact
	contact.ID = gs.randomInteger()
	return gs.db.Create(contact).Error
}

// Permet d'afficher tout les contacts
func (gs *GormStore) DisplayContacts() ([]*Contact, error) {
	var contacts []*Contact
	err := gs.db.Find(&contacts).Error
	return contacts, err
}

// Permet d'afficher un contact via son ID
func (gs *GormStore) DisplayContact(idContact int) (*Contact, error) {
	var contact *Contact
	err := gs.db.First(&contact, idContact).Error
	return contact, err
}

// Permet de mettre à jour un contact via son ID
func (gs *GormStore) UpdateContact(idContact int, newName string, newEmail string) error {
	var contactToUpdate *Contact

	// On vérifie si le contact existe
	err := gs.db.First(&contactToUpdate, idContact).Error
	if err != nil {
		return err
	}

	// On met à jour le contact
	if newName != "" && contactToUpdate.Name != newName {
		contactToUpdate.Name = newName
	}

	if newEmail != "" && contactToUpdate.Email != newEmail {
		contactToUpdate.Email = newEmail
	}

	// On met à jour le contact
	contactUpdated := gs.db.Save(&contactToUpdate)

	// Vérifier si il y a une erreur SQL
	if contactUpdated.Error != nil {
		return contactUpdated.Error
	}

	// Vérifier si une ligne à été affectée
	if contactUpdated.RowsAffected == 0 {
		return fmt.Errorf("contact avec l'ID %d non trouvé", idContact)
	}

	return nil
}

// Permet de supprimer un contact via son ID
func (gs *GormStore) DeleteContact(idContact int) error {
	var contact *Contact

	// On supprime le contact en ignorant le soft-delete avec Unscoped()
	result := gs.db.Unscoped().Delete(&contact, idContact)

	// Vérifier si il y a une erreur SQL
	if result.Error != nil {
		return result.Error
	}

	// Vérifier si une ligne à été affectée
	if result.RowsAffected == 0 {
		return fmt.Errorf("contact avec l'ID %d non trouvé", idContact)
	}

	return nil
}

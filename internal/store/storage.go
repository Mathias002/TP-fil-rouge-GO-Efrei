package storage

type IDContact int
type NameContact string
type EmailContact string
type ContactsType map[IDContact]Contact

// Définition d'une structure "Contact"
type Contact struct {
	ID    IDContact    `json:"id" binding:"required"`
	Name  NameContact  `json:"name" binding:"required"`
	Email EmailContact `json:"email" binding:"required,email"`
}

type Storer interface {
	AddContact(contact *Contact) error
	DisplayContacts() ([]*Contact, error)
	DisplayContact(id IDContact) (*Contact, error)
	UpdateContact(id IDContact, newName NameContact, newEmail EmailContact) error
	DeleteContact(id IDContact) error
}

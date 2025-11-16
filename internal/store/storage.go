package storage

// type IDContact int
type NameContact string
type EmailContact string
type ContactsType map[int]Contact

// MAJ de la struct contact avec GORM

// Définition d'une structure "Contact"

type Contact struct {
	ID    int    `gorm:"primaryKey"`
	Name  string `gorm:"type:varchar(100);not null" json:"name" binding:"required"`
	Email string `gorm:"type:varchar(255);not null;unique" json:"email" binding:"required,email"`
}

type Storer interface {
	AddContact(contact *Contact) error
	DisplayContacts() ([]*Contact, error)
	DisplayContact(id int) (*Contact, error)
	UpdateContact(id int, newName string, newEmail string) error
	DeleteContact(id int) error
}

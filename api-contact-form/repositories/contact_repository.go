package repositories

import (
	"time"

	"github.com/dedenurr/contactflow/api-contact-form/models"

	"gorm.io/gorm"
)

type ContactRepository interface{
	// create add a new contact to the database
	Create(contact *models.Contact) error
	// find all contacts from the database
	FindAll() ([]models.Contact, error)
	// find contact by ID
	FindByID(id uint) (*models.Contact, error)
	// Update contact by ID
	Update(contact *models.Contact) error
	// Delete marks a contact as deleted in the database
	Delete(contact *models.Contact) error
}

// contactRepository is the GORM-based implementation of ContactRepository.
type contactRepository struct {
	db *gorm.DB
}

// NewContactRepository creates a new instance of contactRepository with the provided GORM DB.
func NewContactRepository(db *gorm.DB) ContactRepository {
	return &contactRepository{db}
}

// Create adds a new contact to the database.
func (r *contactRepository) Create(contact *models.Contact) error {
	return r.db.Create(contact).Error
}

// FIndAll retrieves all contacts from the database.
func(r *contactRepository) FindAll() ([]models.Contact, error) {
	var contacts []models.Contact
	err := r.db.Where("deleted_at =?", "0000-00-00 00:00:00").Find(&contacts).Error
	return contacts, err
}

// FindByID retrieves a contact by its ID.
func (r *contactRepository) FindByID(id uint) (*models.Contact, error){
	var contact models.Contact
	err := r.db.Where("id = ? AND deleted_at = ?", id, "0000-00-00 00:00:00").First(&contact).Error
	return &contact, err
}

// update modifies an existing contact in the database.
func (r *contactRepository) Update(contact *models.Contact) error{
	return r.db.Save(contact).Error
}

// delete marks a contact as deleted in the database.
func (r *contactRepository) Delete(contact *models.Contact) error{
	contact.DeletedAt = time.Now()	
	return r.db.Save(contact).Error
}
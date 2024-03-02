package repository

import (
	"architecture_go/pkg/store/postgres"
	"architecture_go/services/contact/internal/domain"
	"database/sql"
	"errors"
)

type ContactRepositoryImpl struct {
	DB *postgres.Database
}

func (r *ContactRepositoryImpl) CreateContact(contact *domain.Contact) error {
	_, err := r.DB.Exec("INSERT INTO contacts (id, first_name, last_name, middle_name, phone_number) VALUES ($1, $2, $3, $4, $5)", contact.ID, contact.FullName.FirstName, contact.FullName.LastName, contact.FullName.MiddleName, contact.PhoneNumber)
	if err != nil {
		return err
	}
	return nil
}

func (r *ContactRepositoryImpl) ReadContact(contactID int) (*domain.Contact, error) {
	var contact domain.Contact
	err := r.DB.QueryRow("SELECT id, first_name, last_name, middle_name, phone_number FROM contacts WHERE id = $1", contactID).Scan(&contact.ID, &contact.FullName.FirstName, &contact.FullName.LastName, &contact.FullName.MiddleName, &contact.PhoneNumber)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("contact not found")
		}
		return nil, err
	}
	return &contact, nil
}

func (r *ContactRepositoryImpl) UpdateContact(contactID int, newContact *domain.Contact) error {
	_, err := r.DB.Exec("UPDATE contacts SET first_name = $1, last_name = $2, middle_name = $3, phone_number = $4 WHERE id = $5", newContact.FullName.FirstName, newContact.FullName.LastName, newContact.FullName.MiddleName, newContact.PhoneNumber, contactID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ContactRepositoryImpl) DeleteContact(contactID int) error {
	_, err := r.DB.Exec("DELETE FROM contacts WHERE id = $1", contactID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ContactRepositoryImpl) CreateGroup(group *domain.Group) error {
	_, err := r.DB.Exec("INSERT INTO groups (id ,name) VALUES ($1, $2)", group.GetID(), group.GetName())
	if err != nil {
		return err
	}
	return nil
}

func (r *ContactRepositoryImpl) ReadGroup(groupID int) (*domain.Group, error) {
	var id int
	var name string

	err := r.DB.QueryRow("SELECT id, name FROM groups WHERE id = $1", groupID).Scan(id, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("group not found")
		}
		return nil, err
	}
	return domain.NewGroup(id, name), nil
}

func (r *ContactRepositoryImpl) AddContactToGroup(contactID, groupID int) error {
	_, err := r.DB.Exec("INSERT INTO contact_groups (contact_id, group_id) VALUES ($1, $2)", contactID, groupID)
	if err != nil {
		return err
	}
	return nil
}

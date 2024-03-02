package use_case

import (
	"architecture_go/services/contact/internal/domain"
	"architecture_go/services/contact/internal/repository"
)

type ContactUseCase interface {
	CreateContact(contact *domain.Contact) error
	ReadContact(contactID int) (*domain.Contact, error)
	UpdateContact(contactID int, newContact *domain.Contact) error
	DeleteContact(contactID int) error

	CreateGroup(group *domain.Group) error
	ReadGroup(groupID int) (*domain.Group, error)

	AddContactToGroup(contactID, groupID int) error
}

type ContactUseCaseImpl struct {
	repository.ContactRepository
}



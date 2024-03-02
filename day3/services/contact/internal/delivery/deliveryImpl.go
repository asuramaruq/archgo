package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"architecture_go/services/contact/internal/domain"
	"architecture_go/services/contact/internal/repository"
)

type ContactDeliveryImpl struct {
	ContactRepo repository.ContactRepository
}

func NewContactDelivery(contactRepo repository.ContactRepository) *ContactDeliveryImpl {
	return &ContactDeliveryImpl{
		ContactRepo: contactRepo,
	}
}

func (d *ContactDeliveryImpl) CreateContactHandler(w http.ResponseWriter, r *http.Request) {
	var contact domain.Contact
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	err = d.ContactRepo.CreateContact(&contact)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create contact: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(contact)
}

func (d *ContactDeliveryImpl) ReadContactHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid contact ID: %v", err), http.StatusBadRequest)
		return
	}

	contact, err := d.ContactRepo.ReadContact(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read contact: %v", err), http.StatusInternalServerError)
		return
	}

	if &contact == nil {
		http.Error(w, "contact not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(contact)
}

func (d *ContactDeliveryImpl) UpdateContactHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid contact ID: %v", err), http.StatusBadRequest)
		return
	}

	var updatedContact domain.Contact
	err = json.NewDecoder(r.Body).Decode(&updatedContact)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	updatedContact.ID = id
	err = d.ContactRepo.UpdateContact(id, &updatedContact)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to update contact: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Contact with ID %d updated successfully", id)
}

func (d *ContactDeliveryImpl) DeleteContactHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid contact ID: %v", err), http.StatusBadRequest)
		return
	}

	err = d.ContactRepo.DeleteContact(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to delete contact: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Contact with ID %d deleted successfully", id)
}

func (d *ContactDeliveryImpl) CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	var group domain.Group
	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	err = d.ContactRepo.CreateGroup(&group)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create group: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(group)
}

func (d *ContactDeliveryImpl) ReadGroupHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid group ID: %v", err), http.StatusBadRequest)
		return
	}

	group, err := d.ContactRepo.ReadGroup(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read group: %v", err), http.StatusInternalServerError)
		return
	}

	if &group == nil {
		http.Error(w, "group not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(group)
}

func (d *ContactDeliveryImpl) AddContactToGroupHandler(w http.ResponseWriter, r *http.Request) {
	contactIDStr := r.URL.Query().Get("contact_id")
	groupIDStr := r.URL.Query().Get("group_id")
	contactID, err := strconv.Atoi(contactIDStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid contact ID: %v", err), http.StatusBadRequest)
		return
	}
	groupID, err := strconv.Atoi(groupIDStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid group ID: %v", err), http.StatusBadRequest)
		return
	}

	err = d.ContactRepo.AddContactToGroup(contactID, groupID)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to add contact to group: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Contact with ID %d added to group with ID %d successfully", contactID, groupID)
}

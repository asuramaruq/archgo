package delivery

import (
	"net/http"
)

type ContactDelivery interface {
	CreateContactHandler(w http.ResponseWriter, r *http.Request)
	ReadContactHandler(w http.ResponseWriter, r *http.Request)
	UpdateContactHandler(w http.ResponseWriter, r *http.Request)
	DeleteContactHandler(w http.ResponseWriter, r *http.Request)

	CreateGroupHandler(w http.ResponseWriter, r *http.Request)
	ReadGroupHandler(w http.ResponseWriter, r *http.Request)

	AddContactToGroupHandler(w http.ResponseWriter, r *http.Request)
}

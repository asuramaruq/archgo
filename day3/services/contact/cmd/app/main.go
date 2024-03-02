package main

import (
	"architecture_go/pkg/store/postgres"
	"architecture_go/services/contact/internal"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello World!")
	db := postgres.DBconnection("postgres", "12345678", "localhost", "5432", "archgo")

	fmt.Println("DB connection: ", db)

	defer db.Close()
	contactRepository := internal.NewContactRepository(db)
	contactUseCase := internal.NewContactUseCase(contactRepository)
	contactDelivery := internal.NewContactDelivery(contactUseCase)
	router := mux.NewRouter()

	router.HandleFunc("/contacts", contactDelivery.CreateContactHandler).Methods("POST")
	router.HandleFunc("/contacts/{id}", contactDelivery.ReadContactHandler).Methods("GET")
	router.HandleFunc("/contacts/{id}", contactDelivery.UpdateContactHandler).Methods("PUT")
	router.HandleFunc("/contacts/{id}", contactDelivery.DeleteContactHandler).Methods("DELETE")
	router.HandleFunc("/groups", contactDelivery.CreateGroupHandler).Methods("POST")
	router.HandleFunc("/groups/{id}", contactDelivery.ReadGroupHandler).Methods("GET")
	router.HandleFunc("/contacts/{contactID}/groups/{groupID}", contactDelivery.AddContactToGroupHandler).Methods("POST")

	http.Handle("/", router)
	http.ListenAndServe(":3000", nil)

}

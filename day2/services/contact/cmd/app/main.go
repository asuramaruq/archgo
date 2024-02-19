package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"architecture_go/pkg/store/postgres"
	"architecture_go/services/contact/internal"
)

func main() {
	fmt.Println("Hello World!")

	db, err := postgres.DBconnection("postgres", "12345678", "localhost", "5432", "archgo")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	defer db.Close()

	contactRepository := internal.NewContactRepository()
	contactUseCase := internal.NewContactUseCase(contactRepository)
	contactDelivery := internal.NewContactDelivery(contactUseCase)

	router := mux.NewRouter()

	router.HandleFunc("/contacts", contactDelivery.CreateContactHandler).Methods("POST")
	router.HandleFunc("/contacts/{id}", contactDelivery.ReadContactHandler).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":3000", nil)

}

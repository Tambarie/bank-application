package app

import (
	"bank-application/domain"
	"bank-application/service"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	//wiring
	c := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	// define routes
	router.HandleFunc("/customers", c.getAllCustomers).Methods(http.MethodGet)

	fmt.Println("server started successfully")
	//starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}

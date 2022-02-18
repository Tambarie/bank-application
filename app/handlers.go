package app

import (
	"bank-application/service"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (c *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := c.service.GetAllCustomers()
	if err != nil {
		panic(err)
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}

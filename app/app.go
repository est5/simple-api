package app

import (
	"github.com/est5/simple-api/domain"
	"github.com/est5/simple-api/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	router := mux.NewRouter()

	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}

}

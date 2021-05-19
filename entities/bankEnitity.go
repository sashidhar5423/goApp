package entities

import (
	"banks/services"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// type RBI interface {
// 	bankRegstration() Bank
// }
func RouterFunc() {
	fmt.Println("from router function")
	services.DefalutFunc()
	router := mux.NewRouter()

	router.HandleFunc("/", services.GetAll).Methods("GET")
	router.HandleFunc("/getBank/{name}", services.GetBank).Methods("GET")
	router.HandleFunc("/create", services.CreateBank).Methods("POST")
	router.HandleFunc("/update/{name}", services.UpdateBank).Methods("PUT")
	router.HandleFunc("/delete/{name}", services.DeleteBank).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":5252", router))

}

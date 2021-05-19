package services

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Bank struct {
	Name         string `json:"name"`
	Sector       string `json:"sector"`
	Capital      string `json:"capital"`
	NoOfBranches int    `json:"noOfBranches"`
}

var bankArry []Bank

// bankArry = append(bankArry, Bank{Name: "SBI", Sector: "Public", Capital: "Mumbai", NoOfBranches: 38})
// bankArry = append(bankArry, Bank{Name: "ICICI", Sector: "Private", Capital: "Mumbai", NoOfBranches: 26})

//bankArry[] = &en.Banks
func DefalutFunc() {
	bankArry = append(bankArry, Bank{Name: "SBI", Sector: "Public", Capital: "Mumbai", NoOfBranches: 38})
	bankArry = append(bankArry, Bank{Name: "ICICI", Sector: "Private", Capital: "Mumbai", NoOfBranches: 26})
}
func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bankArry)

}
func GetBank(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, data := range bankArry {
		if data.Name == params["name"] {
			json.NewEncoder(w).Encode(data)
			return
		}
	}
	json.NewEncoder(w).Encode(&Bank{})
}
func CreateBank(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var bank Bank
	_ = json.NewDecoder(r.Body).Decode(&bank)
	bankArry = append(bankArry, bank)
	json.NewEncoder(w).Encode(bankArry)

}
func UpdateBank(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, data := range bankArry {
		if data.Name == params["name"] {
			bankArry = append(bankArry[:index], bankArry[index+1:]...)
			var bank Bank
			_ = json.NewDecoder(r.Body).Decode(&bank)
			bankArry = append(bankArry, bank)
			json.NewEncoder(w).Encode(bankArry)
			return
		}

	}

	json.NewEncoder(w).Encode(bankArry)

}
func DeleteBank(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, data := range bankArry {
		if data.Name == params["name"] {
			bankArry = append(bankArry[:index], bankArry[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(bankArry)
}

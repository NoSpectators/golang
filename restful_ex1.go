/*
example of restful-api from
https://www.codementor.io/codehakase/
building-a-restful-api-with-golang-a6yivzqdo
*/

/*NOTE: this requires POSTMAN 
step1: click on "capture API requests with POSTMAN"
step2: run the go file from CLI
step3: pass this as url: "localhost:8080/people" without quotes
RESULT: [{"ID":"1","Firstname":"John","lastname":"Doe",
"address":{"city":"City X","state":"State X"}}, {"ID":"2","Firstname":"Koko","lastname":"Doe",
"address":{"city":"City Z","state":"State Y"}}]
step4: ctr-c to end go file, and check postman status: it should say 
"could not get any response"
*/

package main

import (
	"encoding/json" 			//https://golang.org/doc/articles/json_and_go.html
	"github.com/gorilla/mux" 	//https://github.com/gorilla/mux/blob/master/README.md
	"log" 						//https://golang.org/pkg/log/
	"net/http"					//https://golang.org/pkg/net/http/
)

type Person struct {
	ID        string   `json:"id,omitempty`
	Firstname string   `json:"firstname,omitempty`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}
func GetPeopleEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}
func CreatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}
func DeletePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}

func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
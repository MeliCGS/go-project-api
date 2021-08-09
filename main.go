package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Persona struct {
	ID        string     `json:"id"`
	Nombre    string     `json:"nombre"`
	Apellido  string     `json:"apellido"`
	Edad      int        `json:"edad"`
	Direccion *Direccion `json:"direccion"`
}

type Direccion struct {
	Ciudad    string `json:"ciudad"`
	Municipio string `json:"municipio"`
}

var personas []Persona

func GetPersonasEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(personas)
}

func GetIndividuoEndpoint(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	for _, item := range personas {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Persona{})
}

func CreateIndividuoEndpoint(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	var persona Persona
	_ = json.NewDecoder(r.Body).Decode(&persona)
	persona.ID = params["id"]
	personas = append(personas, persona)
	json.NewEncoder(w).Encode(personas)
}

func DeleteIndividuoEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, item := range personas {
		if item.ID == params["id"] {
			personas = append(personas[:index], personas[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(personas)
}

func main() {
	router := mux.NewRouter()

	personas = append(personas, Persona{
		ID:       "1",
		Nombre:   "Mauricio",
		Apellido: "Garces",
		Edad:     33,
		Direccion: &Direccion{
			Ciudad:    "Bogota",
			Municipio: "Cundinamarca",
		},
	})

	personas = append(personas, Persona{
		ID:       "2",
		Nombre:   "Paul",
		Apellido: "Mackarnie",
		Edad:     99,
	})

	router.HandleFunc("/personas", GetPersonasEndpoint).Methods("GET")
	router.HandleFunc("/personas/{id}", GetIndividuoEndpoint).Methods("GET")
	router.HandleFunc("/personas/{id}", CreateIndividuoEndpoint).Methods("POST")
	router.HandleFunc("/personas/{id}", DeleteIndividuoEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

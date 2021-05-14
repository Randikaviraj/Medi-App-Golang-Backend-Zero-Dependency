package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/randika/Medi-App-Golang-Backend/database"
)

type doctorName struct {
	name string
}

type speciality struct {
	specializedField string
}

func FindByNameController(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		var spe speciality

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&spe)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var data []database.Speciality
		if err := database.QueryBySpecialityName(spe.specializedField, &data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		jsonData, err := json.Marshal(data)

		if err != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Write(jsonData)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
	}

}

func FindBySpecilityController(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		var doc doctorName

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&doc)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var data []database.HospitalAndDoctor
		if err := database.QueryByDoctorName(doc.name, &data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		jsonData, err := json.Marshal(data)

		if err != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Write(jsonData)

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
	}

}

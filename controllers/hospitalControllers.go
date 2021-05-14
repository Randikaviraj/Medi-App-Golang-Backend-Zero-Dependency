package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/randika/Medi-App-Golang-Backend/database"
)

type hospital struct {
	hospitalName string
}

type date struct {
	hospital
	date string
}

func FindByHospitalController(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		var hos hospital

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&hos)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var data []database.Hospital

		if err := database.QueryByHospitalName(hos.hospitalName, &data); err != nil {
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

func FindByDateController(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		var date date

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&date)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var data []database.Doctor

		if err := database.QueryDoctorByDate(date.hospitalName, date.date, &data); err != nil {
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
